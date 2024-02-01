<?php

declare(strict_types=1);

namespace App\Middleware;

use App\Module\Cache\Login as CacheLogin;
use App\Module\Db\Dao\Platform\Admin;
use Hyperf\Di\Annotation\Inject;
use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\RequestHandlerInterface;

class SceneLoginOfPlatform implements \Psr\Http\Server\MiddlewareInterface
{
    #[Inject]
    protected \App\Module\Logic\Login $logic;

    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        try {
            $container = getContainer();
            $sceneCode = $container->get(\App\Module\Logic\Auth\Scene::class)->getCurrentSceneCode();
            /**--------验证token 开始--------**/
            $token = $this->logic->getCurrentToken($sceneCode);
            if (empty($token)) {
                throwFailJson(39994000);
            }
            $jwt = make($sceneCode . 'Jwt');
            $payload = $jwt->verifyToken($token);
            /**--------验证token 结束--------**/

            /**--------选做。限制多地登录，多设备登录等情况下可用（前置条件：登录时做过token缓存） 开始--------**/
            $cacheLogin = getCache(CacheLogin::class);
            $cacheLogin->setTokenKey($payload['id'], $sceneCode);
            $checkToken = $cacheLogin->getToken();
            if ($checkToken != $token) {
                throwFailJson(39994002);
            }
            /**--------选做。限制多地登录，多设备登录等情况下可用（前置条件：登录时做过token缓存） 结束--------**/

            /**--------获取登录用户信息并验证 开始--------**/
            $info = getDao(Admin::class)
                //->parseField(['adminId', 'nickname', 'avatar', 'isStop'])
                ->parseFilter(['adminId' => $payload['id']])
                ->info();
            if (empty($info)) {
                throwFailJson(39994003);
            }
            if ($info->isStop) {
                throwFailJson(39994004);
            }

            $info->loginId  = $payload['id']; //所有场景追加这个字段，方便统一调用
            $this->logic->setCurrentInfo($info, $sceneCode);    //用户信息保存在协程上下文
            /**--------获取用户信息并验证 结束--------**/

            $response = $handler->handle($request);
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
