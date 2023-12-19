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
        $uploadType = getConfig('inDb.platformConfig.uploadType');
        switch ($uploadType) {
            case 'uploadOfAliyunOss':
                $config = [
                    'accessKeyId' =>  getConfig('inDb.platformConfig.uploadOfAliyunOssAccessKeyId'),
                    'accessKeySecret' => getConfig('inDb.platformConfig.uploadOfAliyunOssAccessKeySecret'),
                    'host' => getConfig('inDb.platformConfig.uploadOfAliyunOssHost'),
                    'bucket' => getConfig('inDb.platformConfig.uploadOfAliyunOssBucket'),
                    'callbackUrl' => getConfig('inDb.platformConfig.uploadOfAliyunOssCallbackUrl'),
                ];
                return make(\App\Plugin\Upload\UploadOfAliyunOss::class, ['config' => $config]);
            case 'uploadOfLocal':
            default:
                $config = [
                    'url' =>  getConfig('inDb.platformConfig.uploadOfLocalUrl'),
                    'signKey' => getConfig('inDb.platformConfig.uploadOfLocalSignKey'),
                    'fileSaveDir' => getConfig('inDb.platformConfig.uploadOfLocalFileSaveDir'),
                    'fileUrlPrefix' => getConfig('inDb.platformConfig.uploadOfLocalFileUrlPrefix'),
                ];
                return make(\App\Plugin\Upload\UploadOfLocal::class, ['config' => $config]);
        }
    },
    //短信组件
    'sms' => function (ContainerInterface $container) {
        $smsType = getConfig('inDb.platformConfig.smsType');
        switch ($smsType) {
            case 'smsOfAliyun':
            default:
                $config = [
                    'accessKeyId' =>  getConfig('inDb.platformConfig.smsOfAliyunAccessKeyId'),
                    'accessKeySecret' => getConfig('inDb.platformConfig.smsOfAliyunAccessKeySecret'),
                    'endpoint' => getConfig('inDb.platformConfig.smsOfAliyunEndpoint'),
                    'signName' => getConfig('inDb.platformConfig.smsOfAliyunSignName'),
                    'templateCode' => getConfig('inDb.platformConfig.smsOfAliyunTemplateCode'),
                ];
                return make(\App\Plugin\Sms\SmsOfAliyun::class, ['config' => $config]);
        }
    },
    //实名认证组件
    'idCard' => function (ContainerInterface $container) {
        $idCardType = getConfig('inDb.platformConfig.idCardType');
        switch ($idCardType) {
            case 'idCardOfAliyun':
            default:
                $config = [
                    'host' =>  getConfig('inDb.platformConfig.idCardOfAliyunHost'),
                    'path' => getConfig('inDb.platformConfig.idCardOfAliyunPath'),
                    'appcode' => getConfig('inDb.platformConfig.idCardOfAliyunAppcode'),
                ];
                return make(\App\Plugin\IdCard\IdCardOfAliyun::class, ['config' => $config]);
        }
    },
];
