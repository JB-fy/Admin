<?php

declare(strict_types=1);

namespace App\Module\Validation\Platform;

use App\Module\Validation\AbstractValidation;

class Admin extends AbstractValidation
{
    protected array $rule = [
        'adminId' => 'sometimes|required|integer|min:1',
        'account' => 'sometimes|required_without:phone|alpha_dash|between:6,30',
        'phone' => 'sometimes|required_without:account|telephone_number|between:1,30',
        'password' => 'sometimes|required|alpha_dash|size:32',
        'nickname' => 'sometimes|required|alpha_dash|between:1,30',
        'avatar' => 'sometimes|required|url|between:1,120',
        'roleIdArr' => 'sometimes|required_if_null|array|min:1',
        'roleIdArr.*' => 'sometimes|required|integer|min:1|distinct',
        'isStop' => 'sometimes|required|integer|in:0,1',

        'roleId' => 'sometimes|required|integer|min:1',
    ];

    protected array $scene = [
        'create' => [
            'only' => [
                'account',
                'phone',
                'password',
                'nickname',
                'avatar',
                'roleIdArr',
                'roleIdArr.*',
                'remark',
                'isStop',
            ],
            'remove' => [
                'account' => ['sometimes'],
                'phone' => ['sometimes'],
                'password' => ['sometimes'],
                'roleIdArr' => ['sometimes'],
            ]
        ],
        'update' => [
            'only' => [
                'id',
                'account',
                'phone',
                'password',
                'nickname',
                'avatar',
                'roleIdArr',
                'roleIdArr.*',
                'remark',
                'isStop',
            ],
            'remove' => [
                'id' => ['sometimes']
            ]
        ],
    ];
}
