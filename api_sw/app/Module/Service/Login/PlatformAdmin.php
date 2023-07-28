<?php

declare(strict_types=1);

namespace App\Module\Service\Login;

use App\Module\Cache\Login as CacheLogin;
use App\Module\Db\Dao\Platform\Admin;
use App\Module\Service\AbstractService;
use Hyperf\Di\Annotation\Inject;

class PlatformAdmin extends AbstractService
{
    #[Inject]
    protected \App\Module\Logic\Login $logic;

    /**
     * 获取加密字符串
     *
     * @param string $account
     * @param string $sceneCode
     * @return void
     */
    public function encryptStr(string $account, string $sceneCode)
    {
        $encryptStr = $this->logic->createEncryptStr($account, $sceneCode);
        throwSuccessJson(['encryptStr' => $encryptStr]);
    }

    /**
     * 登录
     *
     * @param string $account
     * @param string $password
     * @param string $sceneCode
     * @return void
     */
    public function login(string $account, string $password, string $sceneCode)
    {
        switch ($sceneCode) {
            case 'platform':
                /**--------验证账号密码 开始--------**/
                $info = getDao(Admin::class)->parseFilter(['accountOrPhone' => $account])->info();
                if (empty($info)) {
                    throwFailJson(39990000);
                }
                if ($info->isStop) {
                    throwFailJson(39990001);
                }
                if (!$this->logic->checkPassword($info->password, $password, $account, $sceneCode)) {
                    throwFailJson(39990000);
                }
                /**--------验证账号密码 结束--------**/

                //生成token
                $payload = [
                    'id' => $info->adminId
                ];
                $jwt = make($sceneCode . 'Jwt');
                $token = $jwt->createToken($payload);

                //缓存token（选做。限制多地登录，多设备登录等情况下可用）
                $cacheLogin = getCache(CacheLogin::class);
                $cacheLogin->setTokenKey($payload['id'], $sceneCode);
                $cacheLogin->setToken($token, $jwt->getConfig()['expireTime']);

                throwSuccessJson(['token' => $token]);
                break;
        }
    }

    /**
     * 验证Token
     *
     * @param string $sceneCode
     * @return void
     */
    public function verifyToken(string $sceneCode)
    {
        switch ($sceneCode) {
            case 'platform':
                /**--------验证token 开始--------**/
                $token = $this->logic->getCurrentToken($sceneCode);
                if (empty($token)) {
                    throwFailJson(39994000);
                }
                $jwt = make($sceneCode . 'Jwt');
                $payload = $jwt->verifyToken($token);
                /**--------验证token 结束--------**/

                /**--------选做。限制多地登录，多设备登录等情况下可用（前提必须在登录时做过token缓存） 开始--------**/
                $cacheLogin = getCache(CacheLogin::class);
                $cacheLogin->setTokenKey($payload['id'], $sceneCode);
                $checkToken = $cacheLogin->getToken();
                if ($checkToken != $token) {
                    throwFailJson(39994002);
                }
                /**--------选做。限制多地登录，多设备登录等情况下可用（前提必须在登录时做过token缓存） 结束--------**/

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
                unset($info->password);
                unset($info->isStop);

                $this->logic->setCurrentInfo($info, $sceneCode);    //用户信息保存在协程上下文
                /**--------获取用户信息并验证 结束--------**/
                break;
        }
    }
}
