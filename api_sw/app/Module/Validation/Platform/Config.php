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
        'aliyunOssRoleArn' => 'string',
        'aliyunOssEndpoint' => 'string',

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
            ]
        ]
    ];
}
