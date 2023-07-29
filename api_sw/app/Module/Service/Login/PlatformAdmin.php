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
     * 获取加密盐
     *
     * @param string $account
     * @param string $sceneCode
     * @return void
     */
    public function salt(string $account)
    {
        $info = getDao(Admin::class)->parseFilter(['accountOrPhone' => $account])->info();
        if (empty($info)) {
            throwFailJson(39990000);
        }
        $saltStatic = $info->salt;
        $sceneCode = 'platform';
        $saltDynamic = $this->logic->createSalt($account, $sceneCode);
        throwSuccessJson(['saltStatic' => $saltStatic, 'saltDynamic' => $saltDynamic]);
    }

    /**
     * 登录
     *
     * @param string $account
     * @param string $password
     * @return void
     */
    public function login(string $account, string $password)
    {
        $sceneCode = 'platform';
        /**--------验证账号密码 开始--------**/
        $info = getDao(Admin::class)->parseFilter(['accountOrPhone' => $account])->info();
        if (empty($info)) {
            throwFailJson(39990000);
        }
        if ($info->isStop) {
            throwFailJson(39990002);
        }
        if (!$this->logic->checkPassword($info->password, $password, $account, $sceneCode)) {
            throwFailJson(39990001);
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
    }
}
