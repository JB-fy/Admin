<?php

declare(strict_types=1);

namespace App\Module\Cache;

class Login extends AbstractCache
{
    protected string $saltKey;

    /**
     * 设置加密字符串缓存key
     *
     * @param string $account
     * @param string $sceneCode
     * @return void
     */
    public function setSaltKey(string $account, string $sceneCode)
    {
        $this->saltKey = sprintf(getConfig('app.cache.saltFormat'), $sceneCode, $account);
    }

    /**
     * 缓存加密字符串
     *
     * @param string $salt
     * @param integer $timeout
     * @return boolean
     */
    public function setSalt(string $salt, int $timeout = 5): bool
    {
        return $this->cache->set($this->saltKey, $salt, $timeout);
    }

    /**
     * 获取加密字符串
     *
     * @return string|boolean
     */
    public function getSalt(): string|bool
    {
        $salt = $this->cache->get($this->saltKey);
        $this->cache->del($this->saltKey);
        return $salt;
    }

    /*----------------token相关 开始----------------*/
    protected string $tokenKey;

    /**
     * 设置token缓存key
     *
     * @param string|integer $id
     * @param string $sceneCode
     * @return void
     */
    public function setTokenKey(string|int $id, string $sceneCode)
    {
        $this->tokenKey = sprintf(getConfig('app.cache.tokenFormat'), $sceneCode, $id);
    }

    /**
     * 缓存token
     *
     * @param string $token
     * @param int $timeout
     * @return boolean
     */
    public function setToken(string $token, int $timeout = 7200): bool
    {
        return $this->cache->set($this->tokenKey, $token, $timeout);
    }

    /**
     * 获取token
     *
     * @return string|boolean
     */
    public function getToken(): string|bool
    {
        return $this->cache->get($this->tokenKey);
    }
    /*----------------token相关 结束----------------*/
}
