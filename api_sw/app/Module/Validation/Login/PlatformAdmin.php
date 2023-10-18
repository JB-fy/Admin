<?php

declare(strict_types=1);

namespace App\Module\Validation\Login;

use App\Module\Validation\AbstractValidation;

class PlatformAdmin extends AbstractValidation
{
    protected array $rule = [
        'loginName' => 'required|alpha_dash|between:4,30',
        'password' => 'required|alpha_num|size:32',
    ];

    protected array $scene = [
        'salt' => [
            'only' => [
                'loginName',
            ],
        ],
        'login' => [
            'only' => [
                'loginName',
                'password'
            ],
        ],
    ];
}
