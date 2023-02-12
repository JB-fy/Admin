<?php

declare(strict_types=1);

use Psr\Container\ContainerInterface;

/**
 * This file is part of Hyperf.
 *
 * @link     https://www.hyperf.io
 * @document https://hyperf.wiki
 * @contact  group@hyperf.io
 * @license  https://github.com/hyperf/hyperf/blob/master/LICENSE
 */
return [
    //平台管理员JWT插件
    'platformAdminJwt' => function (ContainerInterface $container) {
        $sceneInfo = getConfig('inDb.authScene.platformAdmin');
        return make(\App\Plugin\Jwt::class, ['config' => $sceneInfo->sceneConfig]);
    },
    //上传组件
    'upload' => function (ContainerInterface $container) {
        $config = [
            'accessId' =>  getConfig('inDb.platformConfig.aliyunOssAccessId'),
            'accessSecret' => getConfig('inDb.platformConfig.aliyunOssAccessSecret'),
            'host' => getConfig('inDb.platformConfig.aliyunOssHost'),
            'bucket' => getConfig('inDb.platformConfig.aliyunOssBucket'),
        ];
        return make(\App\Plugin\Upload\AliyunOss::class, ['config' => $config]);
    },
];
