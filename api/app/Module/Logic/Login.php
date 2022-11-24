<?php

declare(strict_types=1);

namespace App\Module\Logic;

use App\Module\Cache\Login as CacheLogin;
use Hyperf\Context\Context;
use Hyperf\HttpServer\Contract\RequestInterface;

class Login extends AbstractLogic
{
    /**
     * 生成加密字符串
     *
     * @param string $account
     * @param string $type
     * @return string
     */
    public function createEncryptStr(string $account, string $type): string
    {
        $cacheLogin = getCache(CacheLogin::class);
        $cacheLogin->setEncryptStrKey($account, $type);
        $encryptStr = randStr(8);
        $cacheLogin->setEncryptStr($encryptStr);
        return $encryptStr;
    }

    /**
     * 验证密码是否正确
     *
     * @param string $rawPassword
     * @param string $password
     * @param string $type
     * @return boolean
     */
    public function checkPassword(string $rawPassword, string $password, string $account, string $type): bool
    {
        $cacheLogin = getCache(CacheLogin::class);
        $cacheLogin->setEncryptStrKey($account, $type);
        $encryptStr = $cacheLogin->getEncryptStr();
        return md5($rawPassword . $encryptStr) == $password;
    }

    /**
     * 获取类型对应的jwt
     *  注意：
     * @param string $type
     * @return \App\Plugin\Jwt
     */
    public function getJwt(string $type): \App\Plugin\Jwt
    {
        //return make($type . 'Jwt');   //数据库更改配置可以马上生效。
        return $this->container->get($type . 'Jwt');    //需要重启服务才能生效。但不用每次使用jwt都要从数据库取配置再实例化
    }

    /**
     * 获取类型对应的请求Token
     *
     * @param string $type
     * @return string|null
     */
    public function getRequestToken(string $type): ?string
    {
        switch ($type) {
            case 'platformAdmin':
                return $this->container->get(RequestInterface::class)->header('PlatformAdminToken');
        }
    }

    /**
     * 在当前请求中，设置登录用户信息
     * 
     * @param object $info
     * @param string $type
     * @return void
     */
    public function setInfo(object $info, string $type)
    {
        switch ($type) {
            case 'platformAdmin':
                $request = Context::get(\Psr\Http\Message\ServerRequestInterface::class);
                $request = $request->withAttribute('platformAdminInfo', $info);
                Context::set(\Psr\Http\Message\ServerRequestInterface::class, $request);
                break;
        }
    }

    /**
     * 获取当前请求中的登录用户信息
     * 
     * @param string $type
     * @return object
     */
    public function getInfo(string $type): object
    {
        switch ($type) {
            case 'platformAdmin':
                return $this->container->get(RequestInterface::class)->getAttribute('platformAdminInfo');
        }
    }
}
