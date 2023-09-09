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
    'platformJwt' => function (ContainerInterface $container) {
        $sceneInfo = getConfig('inDb.authScene.platform');
        return make(\App\Plugin\Jwt::class, ['config' => $sceneInfo->sceneConfig]);
    },
    //上传组件
    'upload' => function (ContainerInterface $container) {
        /* $uploadType = getConfig('inDb.platformConfig.uploadType');
        switch ($uploadType) {
            case 'local':
            case 'aliyunOss':
            default:
        } */
        $config = [
            'accessKeyId' =>  getConfig('inDb.platformConfig.aliyunOssAccessKeyId'),
            'accessKeySecret' => getConfig('inDb.platformConfig.aliyunOssAccessKeySecret'),
            'host' => getConfig('inDb.platformConfig.aliyunOssHost'),
            'bucket' => getConfig('inDb.platformConfig.aliyunOssBucket'),
            'callbackUrl' => getConfig('inDb.platformConfig.aliyunOssCallbackUrl'),
        ];
        return make(\App\Plugin\Upload\AliyunOss::class, ['config' => $config]);
    },
];
