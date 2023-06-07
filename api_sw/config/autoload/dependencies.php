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
            'accessKeyId' =>  getConfig('inDb.platformConfig.aliyunOssAccessKeyId'),  //LTAI5tHx81H64BRJA971DPZF  |   LTAI5tSjYikt3bX33riHezmk
            'accessKeySecret' => getConfig('inDb.platformConfig.aliyunOssAccessKeySecret'),   //nJyNpTtUuIgZqx21FF4G2zi0WHOn51    |   k4uRZU6flv73yz1j4LJu9VY5eNlHas
            'host' => getConfig('inDb.platformConfig.aliyunOssHost'),   //http://oss-cn-hongkong.aliyuncs.com    |   https://oss-cn-hangzhou.aliyuncs.com
            'bucket' => getConfig('inDb.platformConfig.aliyunOssBucket'),   //4724382110    |   gamemt
        ];
        return make(\App\Plugin\Upload\AliyunOss::class, ['config' => $config]);
    },
];
