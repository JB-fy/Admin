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
            'attributeName',
            'attributeName' => [
                'append' => ['ruleName',...], //新增规则
                'remove' => ['ruleName',...]  //删除规则
            ]
        ]
    ]; */

    /* public function sceneEncryptStr()
    {
        return $this->only(['account']);
    } */

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
        $rule = [];
        foreach ($this->scene[$sceneName] as $k => $v) {
            if (is_array($v)) {
                if (isset($this->rule[$k])) {
                    $tmp = explode('|', $this->rule[$k]);
                } else {
                    $tmp = [];
                }
                if (isset($v['append'])) {
                    $tmp = array_unique(array_merge($tmp, $v['append']));
                }
                if (isset($v['remove'])) {
                    $tmp = array_diff($tmp, $v['remove']);
                }
                $rule[$k] = implode('|', $tmp);
            } else {
                $rule[$v] = $this->rule[$v];
            }
        }
        return $rule;
    }
}
