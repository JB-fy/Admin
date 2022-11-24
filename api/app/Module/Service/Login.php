<?php

declare(strict_types=1);

namespace App\Module\Service;

use App\Module\Cache\Login as CacheLogin;
use App\Module\Db\Dao\Platform\Admin;
use App\Module\Logic\Login as LogicLogin;

class Login extends AbstractService
{
    /**
     * 获取加密字符串
     *
     * @param string $account
     * @param string $type 类型。值唯一，否则容易出错。例如：数据库两个不同表，用户表和管理员表，可能存在同样的账号名，同时登录时可能会登录失败
     * @return void
     */
    public function encryptStr(string $account, string $type)
    {
        $encryptStr = $this->container->get(LogicLogin::class)->createEncryptStr($account, $type);
        throwSuccessJson(['encryptStr' => $encryptStr]);
    }

    /**
     * 登录
     *
     * @param string $account
     * @param string $password
     * @param string $type 类型。值唯一，否则容易出错。例如：数据库两个不同表，用户表和管理员表，可能存在同样的账号名，同时登录时可能会登录失败
     * @return void
     */
    public function login(string $account, string $password, string $type)
    {
        switch ($type) {
            case 'platformAdmin':
                /**--------验证账号密码 开始--------**/
                $info = getDao(Admin::class)->where(['loginStr' => $account])->getInfo();
                if (empty($info)) {
                    throwFailJson('001010');
                }
                if ($info->isStop) {
                    throwFailJson('001011');
                }
                if (!$this->container->get(LogicLogin::class)->checkPassword($info->password, $password, $account, $type)) {
                    throwFailJson('001010');
                }
                /**--------验证账号密码 结束--------**/

                //生成token
                $payload = [
                    'id' => $info->adminId
                ];
                $platformAdminJwt = $this->container->get('platformAdminJwt', true);   //如果
                $token = $platformAdminJwt->createToken($payload);

                //缓存token（选做。限制多地登录，多设备登录等情况下可用）
                $cacheLogin = getCache(CacheLogin::class);
                $cacheLogin->setTokenKey($payload['id'], $type);
                $cacheLogin->setToken($token, $platformAdminJwt->getConfig()['expireTime']);

                throwSuccessJson(['token' => $token]);
                break;
            default:
                throwFailJson('001004');
                break;
        }
    }

    /**
     * 验证Token
     *
     * @param string $type  类型。值唯一，否则容易出错。例如：数据库两个不同表，用户表和管理员表，可能存在同样的账号名，同时登录时可能会登录失败
     * @return void
     */
    public function verifyToken(string $type)
    {
        switch ($type) {
            case 'platformAdmin':
                $request = request();
                /**--------验证token 开始--------**/
                $token = $request->header('AdminToken');
                if (empty($token)) {
                    throwFailJson('001400');
                }
                $platformAdminJwt = container('platformAdminJwt');
                $payload = $platformAdminJwt->verifyToken($token);
                /**--------验证token 结束--------**/

                /**--------选做。限制多地登录，多设备登录等情况下可用（前提必须在登录时做过token缓存） 开始--------**/
                $cacheLogin = container(CacheLogin::class, true);
                $cacheLogin->setTokenKey($payload['id'], $type);
                $checkToken = $cacheLogin->getToken();
                if ($checkToken != $token) {
                    throwFailJson('001402');
                }
                /**--------选做。限制多地登录，多设备登录等情况下可用（前提必须在登录时做过token缓存） 结束--------**/

                /**--------获取登录用户信息并验证 开始--------**/
                //用户信息保存在请求对象内
                $request->platformAdminInfo = container(Admin::class, true)
                    //->field(['adminId', 'nickname', 'avatar', 'isStop'])
                    ->where(['adminId' => $payload['id']])
                    ->getInfo();
                if (empty($request->platformAdminInfo)) {
                    throwFailJson('001403');
                }
                if ($request->platformAdminInfo->isStop) {
                    throwFailJson('001404');
                }
                unset($request->platformAdminInfo->password);
                unset($request->platformAdminInfo->isStop);
                /**--------获取用户信息并验证 结束--------**/

                /**--------选做。如果token即将过期，刷新token 开始--------**/
                /* if ($payload['expireTime'] - time() < 5 * 60) {
                    $refreshToken = $platformAdminJwt->getToken($payload);
                    //缓存token（选做。限制多地登录，多设备登录等情况下可用）
                    $cacheLogin->setToken($refreshToken, $type);

                    //refreshToken保存在请求对象内（在exception\handler\Handler内返回给前端，用于刷新token）
                    $request->platformAdminToken = $refreshToken;
                } */
                /**--------选做。如果token即将过期，刷新token 结束--------**/
                break;
            default:
                throwFailJson('001004');
                break;
        }
    }
}
