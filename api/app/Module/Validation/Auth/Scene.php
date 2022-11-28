<?php

declare(strict_types=1);

namespace App\Module\Validation\Auth;

use App\Module\Validation\AbstractValidation;

class Scene extends AbstractValidation
{
    protected array $rule = [
        'id' => 'integer|min:1',
        'excId' => 'integer|min:1',
        'sceneId' => 'integer|min:1',
        'sceneName' => 'alpha_dash|between:1,30',
        'sceneCode' => 'alpha_dash|between:1,30',
        //'isStop' => 'in:' . implode(',', array_keys(trans('const.yesOrNo'))), //需在构造函数中创建该规则
        //'isStop' => 'in:0,1',
    ];

    protected array $scene = [
        'list' => [
            'only' => [
                'id',
                'sceneName',
                'sceneCode',
                'isStop',
            ]
        ],
    ];

    public function __construct()
    {
        $this->rule['isStop'] = 'in:' . implode(',', array_keys(trans('const.yesOrNo')));
        parent::__construct();
    }
}
