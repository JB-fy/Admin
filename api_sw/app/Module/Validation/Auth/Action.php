<?php

declare(strict_types=1);

namespace App\Module\Validation\Auth;

use App\Module\Validation\AbstractValidation;

class Action extends AbstractValidation
{
    protected array $rule = [
        'actionId' => 'sometimes|required|integer|min:1',
        'actionName' => 'sometimes|required|alpha_dash|between:1,30',
        'actionCode' => 'sometimes|required|alpha_dash|between:1,30',
        'sceneIdArr' => 'sometimes|required_if_null|array|min:1',
        'sceneIdArr.*' => 'sometimes|required|integer|min:1|distinct',
        'remark' => 'string|between:1,120',
        'isStop' => 'sometimes|required|integer|in:0,1',

        'sceneId' => 'sometimes|required|integer|min:1',
    ];

    protected array $scene = [
        'create' => [
            'only' => [
                'actionName',
                'actionCode',
                'sceneIdArr',
                'sceneIdArr.*',
                'remark',
                'isStop',
            ],
            'remove' => [
                'actionName' => ['sometimes'],
                'actionCode' => ['sometimes'],
                'sceneIdArr' => ['sometimes'],
            ]
        ],
        'update' => [
            'only' => [
                'idArr',
                'idArr.*',
                'actionName',
                'actionCode',
                'sceneIdArr',
                'sceneIdArr.*',
                'remark',
                'isStop',
            ],
            'remove' => [
                'idArr' => ['sometimes'],
                'idArr.*' => ['sometimes'],
            ]
        ],
    ];
}
