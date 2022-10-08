<?php

declare(strict_types=1);

namespace app\module\validate;

class Login extends AbstractValidate
{
    protected $rule =   [
        'account|账号'  => 'require|chsDash|length:4,20',
        'password|密码'  => 'require|alphaNum|length:32',
    ];

    protected $message  =   [
        'account.require' => '账号必须',
        'account.alphaNum' => '账号只能由汉字、字母、数字和下划线_及破折号-组成',
        'account.length'     => '账号长度在4-20个字符之间',
        'password.require' => '密码必须',
        'password.alphaNum' => '密码只能由字母、数字组成',
        'password.length'     => '密码长度必须是32个字符',
    ];

    protected $scene = [
        'encryptStr' => ['account']
    ];

    /* public function sceneEncryptStr()
    {
        return $this->only(['account']);
    } */
}
