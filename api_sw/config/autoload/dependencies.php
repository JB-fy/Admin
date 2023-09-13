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
            case 'aliyunOss':
                $config = [
                    'accessKeyId' =>  getConfig('inDb.platformConfig.aliyunOssAccessKeyId'),
                    'accessKeySecret' => getConfig('inDb.platformConfig.aliyunOssAccessKeySecret'),
                    'host' => getConfig('inDb.platformConfig.aliyunOssHost'),
                    'bucket' => getConfig('inDb.platformConfig.aliyunOssBucket'),
                    'callbackUrl' => getConfig('inDb.platformConfig.aliyunOssCallbackUrl'),
                ];
                return make(\App\Plugin\Upload\AliyunOss::class, ['config' => $config]);
            case 'local':
            default:
                $config = [
                    'url' =>  getConfig('inDb.platformConfig.localUploadUrl'),
                    'signKey' => getConfig('inDb.platformConfig.localUploadSignKey'),
                    'fileSaveDir' => getConfig('inDb.platformConfig.localUploadFileSaveDir'),
                    'fileUrlPrefix' => getConfig('inDb.platformConfig.localUploadFileUrlPrefix'),
                ];
                return make(\App\Plugin\Upload\Local::class, ['config' => $config]);
        }
    },
    //短信组件
    'sms' => function (ContainerInterface $container) {
        $smsType = getConfig('inDb.platformConfig.smsType');
        switch ($smsType) {
            case 'aliyunSms':
            default:
                $config = [
                    'accessKeyId' =>  getConfig('inDb.platformConfig.aliyunSmsAccessKeyId'),
                    'accessKeySecret' => getConfig('inDb.platformConfig.aliyunSmsAccessKeySecret'),
                    'endpoint' => getConfig('inDb.platformConfig.aliyunSmsEndpoint'),
                    'signName' => getConfig('inDb.platformConfig.aliyunSmsSignName'),
                    'templateCode' => getConfig('inDb.platformConfig.aliyunSmsTemplateCode'),
                ];
                return make(\App\Plugin\Sms\AliyunSms::class, ['config' => $config]);
        }
    },
    //实名认证组件
    'idCard' => function (ContainerInterface $container) {
        $idCardType = getConfig('inDb.platformConfig.idCardType');
        switch ($idCardType) {
            case 'aliyunIdCard':
            default:
                $config = [
                    'host' =>  getConfig('inDb.platformConfig.aliyunIdCardHost'),
                    'path' => getConfig('inDb.platformConfig.aliyunIdCardPath'),
                    'appcode' => getConfig('inDb.platformConfig.aliyunIdCardAppcode'),
                ];
                return make(\App\Plugin\IdCard\AliyunIdCard::class, ['config' => $config]);
        }
    },
];
