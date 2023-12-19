<?php

declare(strict_types=1);

namespace App\Module\Validation\Platform;

use App\Module\Validation\AbstractValidation;

class Config extends AbstractValidation
{
    protected array $rule = [
        'configKeyArr' => 'sometimes|required_if_null|array|min:1',
        'configKeyArr.*' => 'sometimes|required|string|between:1,30|distinct',

        'hotSearch' => 'array',
        'hotSearch.*' => 'sometimes|required|string|min:1|distinct',
        'userAgreement' => 'string',
        'privacyAgreement' => 'string',

        'packageUrlOfAndroid' => 'url',
        'packageSizeOfAndroid' => 'integer',
        'packageNameOfAndroid' => 'string',
        'isForceUpdateOfAndroid' => 'integer|in:0,1',
        'versionNumberOfAndroid' => 'integer|min:0',
        'versionNameOfAndroid' => 'string',
        'versionIntroOfAndroid' => 'string',

        'packageUrlOfIos' => 'url',
        'packageSizeOfIos' => 'integer',
        'packageNameOfIos' => 'string',
        'isForceUpdateOfIos' => 'integer|in:0,1',
        'versionNumberOfIos' => 'integer|min:0',
        'versionNameOfIos' => 'string',
        'versionIntroOfIos' => 'string',
        'plistUrlOfIos' => 'url',

        'uploadType' => 'string|in:local,aliyunOss',
        'localUploadUrl' => 'url',
        'localUploadSignKey' => 'string',
        'localUploadFileSaveDir' => 'string',
        'localUploadFileUrlPrefix' => 'url',
        'aliyunOssHost' => 'url',
        'aliyunOssBucket' => 'string',
        'aliyunOssAccessKeyId' => 'alpha_dash',
        'aliyunOssAccessKeySecret' => 'alpha_dash',
        'aliyunOssCallbackUrl' => 'string',
        'aliyunOssEndpoint' => 'string',
        'aliyunOssRoleArn' => 'string',

        'payOfAliAppId' => 'string',
        'payOfAliSignType' => 'in:RSA2,RSA',
        'payOfAliPrivateKey' => 'string',
        'payOfAliPublicKey' => 'string',

        'payOfWxAppId' => 'string',
        'payOfWxMchid' => 'string',
        'payOfWxSerialNo' => 'string',
        'payOfWxApiV3Key' => 'string',
        'payOfWxPrivateKey' => 'string',

        'smsType' => 'string|in:aliyunSms',
        'aliyunSmsAccessKeyId' => 'alpha_dash',
        'aliyunSmsAccessKeySecret' => 'alpha_dash',
        'aliyunSmsEndpoint' => 'string',
        'aliyunSmsSignName' => 'string',
        'aliyunSmsTemplateCode' => 'string',

        'idCardType' => 'string|in:aliyunIdCard',
        'aliyunIdCardHost' => 'url',
        'aliyunIdCardPath' => 'string',
        'aliyunIdCardAppcode' => 'string',

        'pushType' => 'string|in:txTpns',
        'txTpnsHost' => 'url',
        'txTpnsAccessIDOfAndroid' => 'string',
        'txTpnsSecretKeyOfAndroid' => 'string',
        'txTpnsAccessIDOfIos' => 'string',
        'txTpnsSecretKeyOfIos' => 'string',
        'txTpnsAccessIDOfMacOS' => 'string',
        'txTpnsSecretKeyOfMacOS' => 'string',

        'vodType' => 'string|in:aliyunVod',
        'aliyunVodBucket' => 'string',
        'aliyunVodAccessKeyId' => 'alpha_dash',
        'aliyunVodAccessKeySecret' => 'alpha_dash',
        'aliyunVodEndpoint' => 'string',
        'aliyunVodRoleArn' => 'string',
    ];

    protected array $scene = [
        'get' => [
            'only' => [
                'configKeyArr',
                'configKeyArr.*'
            ]
        ],
        'save' => [
            'only' => [
                'hotSearch',
                'userAgreement',
                'privacyAgreement',

                'packageUrlOfAndroid',
                'packageSizeOfAndroid',
                'packageNameOfAndroid',
                'isForceUpdateOfAndroid',
                'versionNumberOfAndroid',
                'versionNameOfAndroid',
                'versionIntroOfAndroid',

                'packageUrlOfIos',
                'packageSizeOfIos',
                'packageNameOfIos',
                'isForceUpdateOfIos',
                'versionNumberOfIos',
                'versionNameOfIos',
                'versionIntroOfIos',
                'plistUrlOfIos',

                'uploadType',
                'localUploadUrl',
                'localUploadSignKey',
                'localUploadFileSaveDir',
                'localUploadFileUrlPrefix',
                'aliyunOssHost',
                'aliyunOssBucket',
                'aliyunOssAccessKeyId',
                'aliyunOssAccessKeySecret',
                'aliyunOssCallbackUrl',
                'aliyunOssRoleArn',
                'aliyunOssEndpoint',

                'payOfAliAppId',
                'payOfAliSignType',
                'payOfAliPrivateKey',
                'payOfAliPublicKey',

                'payOfWxAppId',
                'payOfWxMchid',
                'payOfWxSerialNo',
                'payOfWxApiV3Key',
                'payOfWxPrivateKey',

                'smsType',
                'aliyunSmsAccessKeyId',
                'aliyunSmsAccessKeySecret',
                'aliyunSmsEndpoint',
                'aliyunSmsSignName',
                'aliyunSmsTemplateCode',

                'idCardType',
                'aliyunIdCardHost',
                'aliyunIdCardPath',
                'aliyunIdCardAppcode',

                'pushType',
                'txTpnsHost',
                'txTpnsAccessIDOfAndroid',
                'txTpnsSecretKeyOfAndroid',
                'txTpnsAccessIDOfIos',
                'txTpnsSecretKeyOfIos',
                'txTpnsAccessIDOfMacOS',
                'txTpnsSecretKeyOfMacOS',

                'vodType',
                'aliyunVodBucket',
                'aliyunVodAccessKeyId',
                'aliyunVodAccessKeySecret',
                'aliyunVodEndpoint',
                'aliyunVodRoleArn',
            ]
        ]
    ];
}
