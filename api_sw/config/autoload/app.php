<?php

declare(strict_types=1);

//自定义配置文件
return [
    'superPlatformAdminId' => (int) env('SUPER_PLATFORM_ADMIN_ID', 1),  //平台超级管理员id
    //缓存的key的格式
    'cache' => [
        'saltFormat' => 'salt_%s_%s', //加密盐缓存key。参数：场景标识，账号
        'tokenFormat' => 'token_%s_%s', //登录后的token缓存key。参数：场景标识，用户标识
    ],
];
