<?php

declare(strict_types=1);

namespace App\Module\Validation\Platform;

use App\Module\Validation\AbstractValidation;

class Config extends AbstractValidation
{
    protected array $rule = [
        'configKeyArr' => 'sometimes|required_if_null|array|min:1',
        'configKeyArr.*' => 'sometimes|required|string|between:1,30|distinct',

        'aliyunOssAccessId' => 'alpha_dash',
        'aliyunOssAccessSecret' => 'alpha_dash',
        'aliyunOssHost' => 'url',
        'aliyunOssBucket' => 'string',
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
                'aliyunOssAccessId',
                'aliyunOssAccessSecret',
                'aliyunOssHost',
                'aliyunOssBucket',
            ]
        ]
    ];
}
