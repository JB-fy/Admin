<?php

declare(strict_types=1);

namespace App\Module\Validation;

use Hyperf\Contract\ValidatorInterface;
use Hyperf\Di\Annotation\Inject;
use Hyperf\Validation\Contract\ValidatorFactoryInterface;

abstract class AbstractValidation
{
    #[Inject]
    protected ValidatorFactoryInterface $validationFactory;

    protected array $rule = [];

    //建议写在app\Language内的语言文件中。特殊情况写这里
    protected array $message = [];

    //建议写在app\Language内的语言文件中。特殊情况写这里
    protected array $customAttribute = [];

    protected array $scene = [];
    /* protected $scene = [
        'sceneName' => [
            'only'=>[   //只验证哪些字段
                'attributeName',
                ...
            ],
            'append' => [ //新增规则
                'attributeName' => [
                    'rule',
                    ...
                ],
                ...
            ],
            'remove' => [ //删除规则
                'attributeName' => [
                    'ruleName',
                    ...
                ],
                ...
            ],
        ]
    ]; */

    /**
     * 创建验证器
     *
     * @param array $data
     * @return ValidatorInterface
     */
    final public function make(array $data, string $sceneName = ''): ValidatorInterface
    {
        $rule = $this->makeRule($sceneName);
        $validator = $this->validationFactory->make($data, $rule, $this->message, $this->customAttribute);
        return $validator;
        //建议下面这些方法在调用处根据情况使用
        /* $validator->validate(); //验证数据，抛出错误
        if ($validator->fails()) {  //验证数据，不抛出错误
            $rules = $validator->failed();  //获取失败的验证规则
            $errorMessage = $validator->errors()->first();  //获取第一个错误
        }
        $data = $validator->validated();    //这里返回已经验证过的数据。未验证字段不返回 */
    }

    /**
     * 创建验证器所需参数
     *
     * @param string $sceneName
     * @return array
     */
    final protected function makeRule(string $sceneName): array
    {
        if (empty($sceneName) || !isset($this->scene[$sceneName])) {
            return $this->rule;
        }
        //处理顺序 only、remove、append
        $rule = $this->rule;
        if (isset($this->scene[$sceneName]['only'])) {
            foreach ($rule as $k => $v) {
                if (in_array($k, $this->scene[$sceneName]['only'])) {
                    continue;
                }
                unset($rule[$k]);
            }
        }
        foreach ($rule as $k => $v) {
            $tmpRule = explode('|', $v);
            if (isset($this->scene[$sceneName]['remove'][$k])) {
                foreach ($tmpRule as $k1 => $v1) {
                    list($tmpRuleName) = explode(':', $v1);
                    if (in_array($tmpRuleName, $this->scene[$sceneName]['remove'][$k])) {
                        unset($tmpRule[$k1]);
                    }
                }
            }
            if (isset($this->scene[$sceneName]['append'][$k])) {
                $tmpRule = array_unique(array_merge($tmpRule, $this->scene[$sceneName]['append'][$k]));
            }
            $rule[$k] = implode('|', $tmpRule);
        }
        //新增字段并添加规则（是否支持自己看）
        if (isset($this->scene[$sceneName]['append'])) {
            foreach ($this->scene[$sceneName]['append'] as $k => $v) {
                if (isset($rule[$k])) {
                    continue;
                }
                $rule[$k] = implode('|', $v);
            }
        }
        return $rule;
    }
}
