<?php

declare(strict_types=1);

namespace app\module\validate;

use think\Validate;

abstract class AbstractValidate extends Validate
{
    protected $failException = true;    //默认抛出错误

    /**
     * 使用filter_var方式验证
     * @access public
     * @param  mixed $value  字段值
     * @param  mixed $rule  验证规则
     * @return bool
     */
    public function filter($value, $rule): bool
    {
        if (is_string($rule) && strpos($rule, ',')) {
            list($rule, $param) = explode(',', $rule);
        } elseif (is_array($rule)) {
            $param = $rule[1] ?? null;
            $rule  = $rule[0];
        } else {
            $param = null;
        }

        //框架这里有bug，需要增加这个。防止bug
        if ($param === null) {
            return false !== filter_var($value, is_int($rule) ? $rule : filter_id($rule));
        }
        return false !== filter_var($value, is_int($rule) ? $rule : filter_id($rule), $param);
    }
}
