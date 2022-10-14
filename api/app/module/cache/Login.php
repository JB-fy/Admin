<?php

declare(strict_types=1);

namespace app\module\cache;

class Login extends AbstractCache
{
    protected string $encryptStrKey;

    /**
     * 设置加密字符串缓存key
     *
     * @param string $account
     * @param string $type
     * @return void
     */
    public function setEncryptStrKey(string $account, string $type)
    {
        $this->encryptStrKey = sprintf(config('custom.cache.encryptStrFormat'), $type, $account);
    }

    /**
     * 缓存加密字符串
     *
     * @param string $encryptStr
     * @param integer $timeout
     * @return boolean
     */
    public function setEncryptStr(string $encryptStr, int $timeout = 5): bool
    {
        return $this->cache->setEx($this->encryptStrKey, $timeout, $encryptStr);
    }

    /**
     * 获取加密字符串
     *
     * @return string
     */
    public function getEncryptStr(): string
    {
        $encryptStr = $this->cache->get($this->encryptStrKey);
        $this->cache->del($this->encryptStrKey);
        return $encryptStr;
    }

    /*----------------token相关 开始----------------*/
    protected string $tokenKey;

    /**
     * 设置token缓存key
     *
     * @param string|integer $id
     * @param string $type
     * @return void
     */
    public function setTokenKey(string|int $id, string $type)
    {
        $this->tokenKey = sprintf(config('custom.cache.tokenFormat'), $type, $id);
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
        return $this->cache->setEx($this->tokenKey, $timeout, $token);
    }

    /**
     * 获取token
     *
     * @return string
     */
    public function getToken(): string
    {
        return $this->cache->get($this->tokenKey);
    }
    /*----------------token相关 结束----------------*/
}
