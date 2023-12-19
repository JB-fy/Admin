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

        'uploadType' => 'string|in:uploadOfLocal,uploadOfAliyunOss',
        'uploadOfLocalUrl' => 'url',
        'uploadOfLocalSignKey' => 'string',
        'uploadOfLocalFileSaveDir' => 'string',
        'uploadOfLocalFileUrlPrefix' => 'url',
        'uploadOfAliyunOssHost' => 'url',
        'uploadOfAliyunOssBucket' => 'string',
        'uploadOfAliyunOssAccessKeyId' => 'alpha_dash',
        'uploadOfAliyunOssAccessKeySecret' => 'alpha_dash',
        'uploadOfAliyunOssCallbackUrl' => 'string',
        'uploadOfAliyunOssEndpoint' => 'string',
        'uploadOfAliyunOssRoleArn' => 'string',

        'payOfAliAppId' => 'string',
        'payOfAliSignType' => 'in:RSA2,RSA',
        'payOfAliPrivateKey' => 'string',
        'payOfAliPublicKey' => 'string',

        'payOfWxAppId' => 'string',
        'payOfWxMchid' => 'string',
        'payOfWxSerialNo' => 'string',
        'payOfWxApiV3Key' => 'string',
        'payOfWxCertPath' => 'string',

        'smsType' => 'string|in:smsOfAliyun',
        'smsOfAliyunAccessKeyId' => 'alpha_dash',
        'smsOfAliyunAccessKeySecret' => 'alpha_dash',
        'smsOfAliyunEndpoint' => 'string',
        'smsOfAliyunSignName' => 'string',
        'smsOfAliyunTemplateCode' => 'string',

        'idCardType' => 'string|in:idCardOfAliyun',
        'idCardOfAliyunHost' => 'url',
        'idCardOfAliyunPath' => 'string',
        'idCardOfAliyunAppcode' => 'string',

        'pushType' => 'string|in:pushOfTx',
        'pushOfTxHost' => 'url',
        'pushOfTxAndroidAccessID' => 'string',
        'pushOfTxAndroidSecretKey' => 'string',
        'pushOfTxIosAccessID' => 'string',
        'pushOfTxIosSecretKey' => 'string',
        'pushOfTxMacOSAccessID' => 'string',
        'pushOfTxMacOSSecretKey' => 'string',

        'vodType' => 'string|in:vodOfAliyun',
        'vodOfAliyunBucket' => 'string',
        'vodOfAliyunAccessKeyId' => 'alpha_dash',
        'vodOfAliyunAccessKeySecret' => 'alpha_dash',
        'vodOfAliyunEndpoint' => 'string',
        'vodOfAliyunRoleArn' => 'string',
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
                'uploadOfLocalUrl',
                'uploadOfLocalSignKey',
                'uploadOfLocalFileSaveDir',
                'uploadOfLocalFileUrlPrefix',
                'uploadOfAliyunOssHost',
                'uploadOfAliyunOssBucket',
                'uploadOfAliyunOssAccessKeyId',
                'uploadOfAliyunOssAccessKeySecret',
                'uploadOfAliyunOssCallbackUrl',
                'uploadOfAliyunOssRoleArn',
                'uploadOfAliyunOssEndpoint',

                'payOfAliAppId',
                'payOfAliSignType',
                'payOfAliPrivateKey',
                'payOfAliPublicKey',

                'payOfWxAppId',
                'payOfWxMchid',
                'payOfWxSerialNo',
                'payOfWxApiV3Key',
                'payOfWxCertPath',

                'smsType',
                'smsOfAliyunAccessKeyId',
                'smsOfAliyunAccessKeySecret',
                'smsOfAliyunEndpoint',
                'smsOfAliyunSignName',
                'smsOfAliyunTemplateCode',

                'idCardType',
                'idCardOfAliyunHost',
                'idCardOfAliyunPath',
                'idCardOfAliyunAppcode',

                'pushType',
                'pushOfTxHost',
                'pushOfTxAndroidAccessID',
                'pushOfTxAndroidSecretKey',
                'pushOfTxIosAccessID',
                'pushOfTxIosSecretKey',
                'pushOfTxMacOSAccessID',
                'pushOfTxMacOSSecretKey',

                'vodType',
                'vodOfAliyunBucket',
                'vodOfAliyunAccessKeyId',
                'vodOfAliyunAccessKeySecret',
                'vodOfAliyunEndpoint',
                'vodOfAliyunRoleArn',
            ]
        ]
    ];
}
