<?php

declare(strict_types=1);

namespace App\Module\Validation\Platform;

use App\Module\Validation\AbstractValidation;

class Config extends AbstractValidation
{
    protected array $rule = [
        'configKeyArr' => 'sometimes|required_if_null|array|min:1',
        'configKeyArr.*' => 'sometimes|required|string|between:1,30|distinct',

        'userAgreement' => 'string',
        'privacyAgreement' => 'string',

        'uploadType' => 'string|in:local,aliyunOss',
        'localUploadUrl' => 'url',
        'localUploadSignKey' => 'string',
        'localUploadFileSaveDir' => 'string',
        'localUploadFileUrlPrefix' => 'url',
        'aliyunOssHost' => 'url',
        'aliyunOssBucket' => 'string',
        'aliyunOssAccessKeyId' => 'alpha_dash',
        'aliyunOssAccessKeySecret' => 'alpha_dash',
        'aliyunOssRoleArn' => 'string',
        'aliyunOssCallbackUrl' => 'string',

        'smsType' => 'string|in:aliyunSms',
        'aliyunSmsAccessKeyId' => 'alpha_dash',
        'aliyunSmsAccessKeySecret' => 'alpha_dash',
        'aliyunSmsEndpoint' => 'string',
        'aliyunSmsSignName' => 'string',
        'aliyunSmsTemplateCode' => 'string',
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
                'userAgreement',
                'privacyAgreement',

                'uploadType',
                'localUploadUrl',
                'localUploadSignKey',
                'localUploadFileSaveDir',
                'localUploadFileUrlPrefix',
                'aliyunOssHost',
                'aliyunOssBucket',
                'aliyunOssAccessKeyId',
                'aliyunOssAccessKeySecret',
                'aliyunOssRoleArn',
                'aliyunOssCallbackUrl',

                'smsType',
                'aliyunSmsAccessKeyId',
                'aliyunSmsAccessKeySecret',
                'aliyunSmsEndpoint',
                'aliyunSmsSignName',
                'aliyunSmsTemplateCode',
            ]
        ]
    ];
}
