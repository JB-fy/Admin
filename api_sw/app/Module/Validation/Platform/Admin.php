<?php

declare(strict_types=1);

namespace App\Module\Validation\Platform;

use App\Module\Validation\AbstractValidation;

class Admin extends AbstractValidation
{
    protected array $rule = [
        'adminId' => 'sometimes|required|integer|min:1',
        'account' => 'sometimes|required|alpha_dash|between:1,30',
        'phone' => 'sometimes|required|string|between:1,30|regex:/^1[3-9]\d{9}$/',
        'password' => 'sometimes|required|alpha_dash|size:32',
        'nickname' => 'alpha_dash|between:1,30',
        'avatar' => 'url|between:1,120',
        'roleIdArr' => 'sometimes|required_if_null|array|min:1',
        'roleIdArr.*' => 'sometimes|required|integer|min:1|distinct',
        'isStop' => 'sometimes|required|integer|in:0,1',

        'passwordToCheck' => 'sometimes|required_with:account,phone,password|size:32',   //当修改账号，手机，密码时必须

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
                'isStop',
            ],
            'remove' => [
                'account' => ['sometimes', 'required'],
                'phone' => ['sometimes', 'required'],
                'password' => ['sometimes'],
                'roleIdArr' => ['sometimes'],
            ],
            'append' => [
                'account' => ['required_without:phone'],
                'phone' => ['required_without:account']
            ]
        ],
        'update' => [
            'only' => [
                'idArr',
                'idArr.*',
                'account',
                'phone',
                'password',
                'nickname',
                'avatar',
                'roleIdArr',
                'roleIdArr.*',
                'isStop',
            ],
            'remove' => [
                'idArr' => ['sometimes'],
                'idArr.*' => ['sometimes'],
                'account' => ['required'],
                'phone' => ['required'],
            ],
            'append' => [
                'account' => ['required_without:phone'],
                'phone' => ['required_without:account']
            ]
        ],
        'updateSelf' => [
            'only' => [
                'account',
                'phone',
                'nickname',
                'avatar',
                'password',
                'passwordToCheck'
            ],
            'remove' => [
                'passwordToCheck' => ['sometimes']
            ],
            'append' => [
                'password' => ['different:passwordToCheck'],
                // 'passwordToCheck' => ['different:password'],
            ]
        ],
    ];
}
