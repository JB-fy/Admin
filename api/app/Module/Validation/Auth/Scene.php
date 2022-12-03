<?php

declare(strict_types=1);

namespace App\Module\Validation\Auth;

use App\Module\Validation\AbstractValidation;

class Scene extends AbstractValidation
{
    protected array $rule = [
        /*--------场景（list，create，update）共用规则 开始--------*/
        'id' => 'sometimes|required|integer|min:1',
        'sceneName' => 'sometimes|required|alpha_dash|between:1,30',
        'sceneCode' => 'sometimes|required|alpha_dash|between:1,30',
        'sceneConfig' => 'sometimes|required|json',
        //'isStop' => 'sometimes|required|in:' . implode(',', array_keys(trans('const.yesOrNo'))), //需在构造函数中创建该规则
        /*--------场景（list，create，update）共用规则 结束--------*/

        /*--------场景（list）可用规则 开始--------*/
        'excId' => 'sometimes|required|integer|min:1',
        'sceneId' => 'sometimes|required|integer|min:1',
        /*--------场景（list）可用规则 结束--------*/
    ];

    protected array $scene = [
        'list' => [],   //可为空，则默认全部规则
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
                'sceneConfig' => ['sometimes'],
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

    public function __construct()
    {
        $this->rule['isStop'] = 'sometimes|required|in:' . implode(',', array_keys(trans('const.yesOrNo')));
        parent::__construct();
    }
}
