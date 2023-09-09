<?php

declare(strict_types=1);

namespace App\Module\Validation\Platform;

use App\Module\Validation\AbstractValidation;

class Config extends AbstractValidation
{
    protected array $rule = [
        'configKeyArr' => 'sometimes|required_if_null|array|min:1',
        'configKeyArr.*' => 'sometimes|required|string|between:1,30|distinct',

        'uploadType' => 'string',
        'localUploadUrl' => 'url',
        'localUploadSignKey' => 'string',
        'localUploadFileUrlPrefix' => 'url',
        'aliyunOssHost' => 'url',
        'aliyunOssBucket' => 'string',
        'aliyunOssAccessKeyId' => 'alpha_dash',
        'aliyunOssAccessKeySecret' => 'alpha_dash',
        'aliyunOssRoleArn' => 'string',
        'aliyunOssCallbackUrl' => 'string',
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
                'uploadType',
                'localUploadUrl',
                'localUploadSignKey',
                'localUploadFileUrlPrefix',
                'aliyunOssHost',
                'aliyunOssBucket',
                'aliyunOssAccessKeyId',
                'aliyunOssAccessKeySecret',
                'aliyunOssRoleArn',
                'aliyunOssCallbackUrl',
            ]
        ]
    ];
}
