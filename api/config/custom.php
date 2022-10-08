<?php

declare(strict_types=1);

/**
 * 自定义配置文件
 */
return [
    //登陆鉴权的配置
    'auth' => [
        'systemAdmin' => [
            'signType' => 'HS256', //算法。共有以下几种选项：HS256，HS384，HS512，RS256，RS384，RS512
            'signKey' => getenv('WEB_HOST') . 'system', //用于签名的秘钥
            'expireTime' => 4 * 60 * 60, //签名有效时间
        ]
    ],
    //缓存的key的格式
    'cache' => [
        'encryptStrFormat' => 'encryptStr_%s_%s', //加密字符串缓存key。参数：类型，账号
        'tokenFormat' => 'token_%s_%s', //登录后的token缓存key。参数：类型，用户标识
    ],
];
