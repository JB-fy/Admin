<?php

declare(strict_types=1);

namespace app\module\service;

use app\module\cache\Login as CacheLogin;
use app\module\db\table\system\SystemAdmin;
use app\module\logic\Login as LogicLogin;

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
        $encryptStr = container(LogicLogin::class)->createEncryptStr($account, $type);
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
            case 'systemAdmin':
                /**--------验证账号密码 开始--------**/
                if (is_numeric($account)) {
                    $info = container(SystemAdmin::class, true)->where(['phone' => $account])->getInfo();
                } else {
                    $info = container(SystemAdmin::class, true)->where(['account' => $account])->getInfo();
                }
                if (empty($info)) {
                    throwFailJson('001010');
                }
                if ($info->isStop) {
                    throwFailJson('001011');
                }
                if (!container(LogicLogin::class)->checkPassword($info->password, $password, $account, $type)) {
                    throwFailJson('001010');
                }
                /**--------验证账号密码 结束--------**/

                //生成token
                $payload = [
                    'id' => $info->adminId
                ];
                $systemAdminJwt = container($type . 'Jwt');
                $token = $systemAdminJwt->createToken($payload);

                //缓存token（选做。限制多地登录，多设备登录等情况下可用）
                $cacheLogin = container(CacheLogin::class, true);
                $cacheLogin->setTokenKey($payload['id'], $type);
                $cacheLogin->setToken($token, $systemAdminJwt->getConfig()['expireTime']);

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
            case 'systemAdmin':
                $request = request();
                /**--------验证token 开始--------**/
                $token = $request->header('SystemAdminToken');
                if (empty($token)) {
                    throwFailJson('001400');
                }
                $systemAdminJwt = container($type . 'Jwt');
                $payload = $systemAdminJwt->verifyToken($token);
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
                $request->systemAdminInfo = container(SystemAdmin::class, true)
                    //->field(['adminId', 'nickname', 'avatar', 'isStop'])
                    ->where(['adminId' => $payload['id']])
                    ->getInfo();
                if (empty($request->systemAdminInfo)) {
                    throwFailJson('001403');
                }
                if ($request->systemAdminInfo->isStop) {
                    throwFailJson('001404');
                }
                unset($request->systemAdminInfo->password);
                unset($request->systemAdminInfo->isStop);
                /**--------获取用户信息并验证 结束--------**/

                /**--------选做。如果token即将过期，刷新token 开始--------**/
                /* if ($payload['expireTime'] - time() < 5 * 60) {
                    $refreshToken = $systemAdminJwt->getToken($payload);
                    //缓存token（选做。限制多地登录，多设备登录等情况下可用）
                    $cacheLogin->setToken($refreshToken, $type);

                    //refreshToken保存在请求对象内（在exception\handler\Handler内返回给前端，用于刷新token）
                    $request->systemAdminToken = $refreshToken;
                } */
                /**--------选做。如果token即将过期，刷新token 结束--------**/
                break;
            default:
                throwFailJson('001004');
                break;
        }
    }
}
