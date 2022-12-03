<?php

declare(strict_types=1);

namespace App\Module\Validation\Auth;

use App\Module\Validation\AbstractValidation;

class Scene extends AbstractValidation
{
    protected array $rule = [
        'sceneId' => 'sometimes|required|integer|min:1',
        'sceneName' => 'sometimes|required|alpha_dash|between:1,30',
        'sceneCode' => 'sometimes|required|alpha_dash|between:1,30',
        'sceneConfig' => 'json',    //可以为空值。空值需要在Dao类中处理
        //'isStop' => 'sometimes|required|in:' . implode(',', array_keys(trans('const.yesOrNo'))), //需在构造函数中创建该规则

        'id' => 'sometimes|required|integer|min:1',
        'idArr' => 'sometimes|required|array|min:1',
        'idArr.*' => 'sometimes|required|integer|min:1',
        'excId' => 'sometimes|required|integer|min:1',
        'excIdArr' => 'sometimes|required|array|min:1',
        'excIdArr.*' => 'sometimes|required|integer|min:1',
    ];

    protected array $scene = [
        'list' => [],   //可为空，则默认全部规则
        'info' => [
            'only' => [
                'id'
            ],
            'remove' => [
                'id' => ['sometimes']
            ]
        ],
        'create' => [
            'only' => [
                'sceneName',
                'sceneCode',
                'sceneConfig',
                'isStop',
            ],
            'remove' => [
                'sceneName' => ['sometimes'],
                'sceneCode' => ['sometimes'],
            ]
        ],
        'update' => [
            'only' => [
                'id',
                'sceneName',
                'sceneCode',
                'sceneConfig',
                'isStop',
            ],
            'remove' => [
                'id' => ['sometimes']
            ]
        ],
        'delete' => [
            'only' => [
                'idArr',
                'idArr.*'
            ],
            'remove' => [
                'idArr' => ['sometimes'],
                'idArr.*' => ['sometimes']
            ]
        ],
    ];

    public function __construct()
    {
        $this->rule['isStop'] = 'sometimes|required|in:' . implode(',', array_keys(trans('const.yesOrNo')));
        parent::__construct();
    }
}
