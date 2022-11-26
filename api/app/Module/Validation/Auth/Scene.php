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
        'isStop' => 'in:' . implode(',', array_keys(trans('const.yesOrNo'))),
    ];
}
