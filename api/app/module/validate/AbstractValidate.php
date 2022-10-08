<?php

declare(strict_types=1);

namespace app\module\validate;

use think\Validate;

abstract class AbstractValidate extends Validate
{
    protected $failException = true;    //默认抛出错误
}
