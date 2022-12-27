<?php

declare(strict_types=1);

namespace App\Module\Validation\Auth;

use App\Module\Validation\AbstractValidation;

class Role extends AbstractValidation
{
    protected array $rule = [
        'roleId' => 'sometimes|required|integer|min:1',
        'roleName' => 'sometimes|required|alpha_dash|between:1,30',
        'menuIdArr' => 'sometimes|required_if_null|array|min:1',
        'menuIdArr.*' => 'sometimes|required|integer|min:1',
        'actionIdArr' => 'sometimes|required_if_null|array|min:1',
        'actionIdArr.*' => 'sometimes|required|integer|min:1',
        'isStop' => 'sometimes|required|integer|in:0,1',
    ];

    protected array $scene = [
        'create' => [
            'only' => [
                'roleName',
                'menuIdArr',
                'menuIdArr.*',
                'actionIdArr',
                'actionIdArr.*',
                'isStop',
            ],
            'remove' => [
                'roleName' => ['sometimes'],
                'roleCode' => ['sometimes'],
                'menuIdArr' => ['sometimes'],
                'actionIdArr' => ['sometimes'],
            ]
        ],
        'update' => [
            'only' => [
                'id',
                'roleName',
                'menuIdArr',
                'menuIdArr.*',
                'actionIdArr',
                'actionIdArr.*',
                'isStop',
            ],
            'remove' => [
                'id' => ['sometimes']
            ]
        ],
    ];
}
