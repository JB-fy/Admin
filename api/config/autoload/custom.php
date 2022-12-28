<?php

declare(strict_types=1);

//自定义配置文件
return [
    //缓存的key的格式
    'cache' => [
        'encryptStrFormat' => 'encryptStr_%s_%s', //加密字符串缓存key。参数：场景标识，账号
        'tokenFormat' => 'token_%s_%s', //登录后的token缓存key。参数：场景标识，用户标识
    ],
];
