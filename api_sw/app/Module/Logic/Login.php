<?php

declare(strict_types=1);

namespace App\Module\Logic;

use App\Module\Cache\Login as CacheLogin;

class Login extends AbstractLogic
{
    /**
     * 生成加密字符串
     *
     * @param string $account
     * @param string $sceneCode
     * @return string
     */
    public function createEncryptStr(string $account, string $sceneCode): string
    {
        $cacheLogin = getCache(CacheLogin::class);
        $cacheLogin->setEncryptStrKey($account, $sceneCode);
        $encryptStr = randStr(8);
        $cacheLogin->setEncryptStr($encryptStr);
        return $encryptStr;
    }

    /**
     * 验证密码是否正确
     *
     * @param string $rawPassword
     * @param string $password
     * @param string $sceneCode
     * @return boolean
     */
    public function checkPassword(string $rawPassword, string $password, string $account, string $sceneCode): bool
    {
        $cacheLogin = getCache(CacheLogin::class);
        $cacheLogin->setEncryptStrKey($account, $sceneCode);
        $encryptStr = $cacheLogin->getEncryptStr();
        return md5($rawPassword . $encryptStr) == $password;
    }

    /**
     * 在当前请求中，获取对应的Token
     *
     * @param string $sceneCode
     * @return string|null
     */
    public function getCurrentToken(string $sceneCode): ?string
    {
        /* switch ($sceneCode) {
            default:
                return $this->container->get(\Hyperf\HttpServer\Contract\RequestInterface::class)->header(ucfirst($sceneCode) . 'Token');
        } */
        return $this->container->get(\Hyperf\HttpServer\Contract\RequestInterface::class)->header(ucfirst($sceneCode) . 'Token');
    }

    /**
     * 在当前请求中，设置登录用户信息
     * 
     * @param object $info
     * @param string $sceneCode
     * @return void
     */
    public function setCurrentInfo(object $info, string $sceneCode)
    {
        \Hyperf\Context\Context::set($sceneCode . 'Info', $info);
    }

    /**
     * 当前请求中，获取登录用户信息
     * 
     * @param string $sceneCode
     * @return object
     */
    public function getCurrentInfo(string $sceneCode): object
    {
        return \Hyperf\Context\Context::get($sceneCode . 'Info');
    }
}
