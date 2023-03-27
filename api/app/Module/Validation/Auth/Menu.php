<?php

declare(strict_types=1);

namespace App\Module\Validation\Auth;

use App\Module\Validation\AbstractValidation;

class Menu extends AbstractValidation
{
    protected array $rule = [
        'menuId' => 'sometimes|required|integer|min:1',
        'sceneId' => 'sometimes|required|integer|min:1',
        'pid' => 'sometimes|required|integer|min:0',
        'menuName' => 'sometimes|required|alpha_dash|between:1,30',
        'extraData' => 'json',    //可以为空值。空值需要在Dao类中处理
        'sort' => 'sometimes|required|integer|min:0|max:100',
        'isStop' => 'sometimes|required|integer|in:0,1',
    ];

    protected array $scene = [
        'create' => [
            'only' => [
                'sceneId',
                'pid',
                'menuName',
                'extraData',
                'sort',
                'isStop',
            ],
            'remove' => [
                'sceneId' => ['sometimes'],
                'menuName' => ['sometimes'],
            ]
        ],
        'update' => [
            'only' => [
                'idArr',
                'idArr.*',
                'sceneId',
                'pid',
                'menuName',
                'extraData',
                'sort',
                'isStop',
            ],
            'remove' => [
                'idArr' => ['sometimes'],
                'idArr.*' => ['sometimes'],
            ]
        ],
    ];
}
