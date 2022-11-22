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

    /* protected $scene = [
        'encryptStr' => ['account']
    ];

    public function sceneEncryptStr()
    {
        return $this->only(['account']);
    } */

    /**
     * 获取验证器
     *
     * @param array $data
     * @return ValidatorInterface
     */
    final public function getValidator(array $data): ValidatorInterface
    {
        $validator = $this->validationFactory->make($data, $this->rule, $this->message, $this->customAttribute);
        return $validator;
        /* $validator->validate(); //验证数据，抛出错误
        if ($validator->fails()) {  //验证数据，不抛出错误
            $errorMessage = $validator->errors()->first();  //获取第一个错误
        }
        $data = $validator->validated();    //这里返回已经验证过的数据。未验证字段不返回
        return $data; */
    }
}
