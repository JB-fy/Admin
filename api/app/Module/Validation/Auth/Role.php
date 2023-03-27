<?php

declare(strict_types=1);

namespace App\Module\Validation\Auth;

use App\Module\Validation\AbstractValidation;

class Role extends AbstractValidation
{
    protected array $rule = [
        'roleId' => 'sometimes|required|integer|min:1',
        'sceneId' => 'sometimes|required|integer|min:1',
        'roleName' => 'sometimes|required|alpha_dash|between:1,30',
        'menuIdArr' => 'sometimes|required_if_null|array|min:0',
        'menuIdArr.*' => 'sometimes|required|integer|min:1|distinct',
        'actionIdArr' => 'sometimes|required_if_null|array|min:0',
        'actionIdArr.*' => 'sometimes|required|integer|min:1|distinct',
        'isStop' => 'sometimes|required|integer|in:0,1',
    ];

    protected array $scene = [
        'create' => [
            'only' => [
                'roleName',
                'sceneId',
                'menuIdArr',
                'menuIdArr.*',
                'actionIdArr',
                'actionIdArr.*',
                'isStop',
            ],
            'remove' => [
                'roleName' => ['sometimes'],
                'sceneId' => ['sometimes'],
                'roleCode' => ['sometimes'],
                'menuIdArr' => ['sometimes'],
                'actionIdArr' => ['sometimes'],
            ]
        ],
        'update' => [
            'only' => [
                'idArr',
                'idArr.*',
                'roleName',
                'sceneId',
                'menuIdArr',
                'menuIdArr.*',
                'actionIdArr',
                'actionIdArr.*',
                'isStop',
            ],
            'remove' => [
                'idArr' => ['sometimes'],
                'idArr.*' => ['sometimes'],
            ]
        ],
    ];
}
