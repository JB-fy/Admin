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
        'isStop' => 'sometimes|required|in:0,1',
    ];

    protected array $scene = [
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
    ];

    /* public function __construct()
    {
        $this->rule['isStop'] = 'sometimes|required|in:' . implode(',', array_keys(trans('const.yesOrNo')));
        parent::__construct();
    } */
}
