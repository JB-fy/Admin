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
     * @param string $type
     * @return string
     */
    public function createEncryptStr(string $account, string $type): string
    {
        $cacheLogin = make(CacheLogin::class);
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
        $cacheLogin = make(CacheLogin::class);
        $cacheLogin->setEncryptStrKey($account, $type);
        $encryptStr = $cacheLogin->getEncryptStr();
        return md5($rawPassword . $encryptStr) == $password;
    }
}
