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

    protected array $ruleOfCommon = [
        'id' => 'sometimes|required|integer|min:1',
        'idArr' => 'sometimes|required|array|min:1',
        'idArr.*' => 'sometimes|required|integer|min:1',
        'excId' => 'sometimes|required|integer|min:1',
        'excIdArr' => 'sometimes|required|array|min:1',
        'excIdArr.*' => 'sometimes|required|integer|min:1',
        'keyword' => 'sometimes|required|alpha_dash|between:1,30',

        'startTime' => 'sometimes|required|date',
        'endTime' => 'sometimes|required|date|after_or_equal:startTime',
    ];
    protected array $rule = [];

    protected array $sceneOfCommon = [
        'list' => [],   //可为空，则默认全部规则
        'tree' => [],
        'info' => [
            'only' => [
                'id'
            ],
            'remove' => [
                'id' => ['sometimes']
            ],
            'append' => [
                'field' => [
                    'sometimes',
                    'required_if_null',
                    'array',
                ]
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
    protected array $scene = [];    //注意处理顺序：only->remove->append
    /* protected $scene = [
        'sceneName' => [
            'only'=>[   //只验证哪些字段
                'attributeName',
                ...
            ],
            'remove' => [ //删除规则
                'attributeName' => [
                    'ruleName',
                    ...
                ],
                ...
            ],
            'append' => [ //新增规则
                'attributeName' => [
                    'rule',
                    ...
                ],
                ...
            ],
        ]
    ]; */

    //建议写在app\Language内的语言文件中。特殊情况写这里
    protected array $message = [];
    //建议写在app\Language内的语言文件中。特殊情况写这里
    protected array $customAttribute = [];

    /**
     * 创建验证器
     *
     * @param array $data
     * @param string $sceneName
     * @return ValidatorInterface
     */
    final public function make(array $data, string $sceneName = ''): ValidatorInterface
    {
        $rule = $this->makeRule($sceneName);
        $validator = $this->validationFactory->make($data, $rule, $this->message, $this->customAttribute);
        return $validator;
        //建议下面这些方法在调用处根据情况使用
        /* $data = $validator->validate(); //验证数据，抛出错误，且只返回验证的字段。
        $data = $validator->validated();    //验证数据，抛出错误，且只返回验证的字段。不存在的字段不验证，相当于自动增加sometimes规则
        if ($validator->fails()) {  //验证数据，不抛出错误
            $rules = $validator->failed();  //获取失败的验证规则
            $errorMessage = $validator->errors()->first();  //获取第一个错误
        } */
    }

    /**
     * 创建验证器所需参数
     *
     * @param string $sceneName
     * @return array
     */
    final protected function makeRule(string $sceneName): array
    {
        $rule = array_merge($this->ruleOfCommon, $this->rule);
        $scene = array_merge($this->sceneOfCommon, $this->scene);
        if ($sceneName === '') {
            return $rule;
        }
        if (!isset($scene[$sceneName])) {
            //return $rule;
            throwFailJson(89999997);
        }
        if (empty($scene[$sceneName])) {
            return $rule;
        }
        //只验证哪些字段
        if (isset($scene[$sceneName]['only'])) {
            foreach ($rule as $k => $v) {
                if (in_array($k, $scene[$sceneName]['only'])) {
                    continue;
                }
                unset($rule[$k]);
            }
        }
        //删除规则
        foreach ($rule as $k => $v) {
            $tmpRule = explode('|', $v);
            if (isset($scene[$sceneName]['remove'][$k])) {
                foreach ($tmpRule as $k1 => $v1) {
                    list($tmpRuleName) = explode(':', $v1);
                    if (in_array($tmpRuleName, $scene[$sceneName]['remove'][$k])) {
                        unset($tmpRule[$k1]);
                    }
                }
            }
            if (isset($scene[$sceneName]['append'][$k])) {
                $tmpRule = array_unique(array_merge($tmpRule, $scene[$sceneName]['append'][$k]));
            }
            $rule[$k] = implode('|', $tmpRule);
        }
        //新增规则
        if (isset($scene[$sceneName]['append'])) {
            foreach ($scene[$sceneName]['append'] as $k => $v) {
                if (isset($rule[$k])) {
                    continue;
                }
                $rule[$k] = implode('|', $v);
            }
        }
        return $rule;
    }
}
