<?php

declare(strict_types=1);

namespace App\Listener;

use Hyperf\Event\Annotation\Listener;
use Hyperf\Event\Contract\ListenerInterface;
use Hyperf\Validation\Event\ValidatorFactoryResolved;

#[Listener]
class ValidatorFactoryResolvedListener implements ListenerInterface
{

    public function listen(): array
    {
        return [
            ValidatorFactoryResolved::class,
        ];
    }

    public function process(object $event): void
    {
        /**
         * @var \Hyperf\Validation\Contract\ValidatorFactoryInterface 
         */
        $validatorFactory = $event->validatorFactory;
        /**
         *  注册了 required_if_null 验证器（必须使用extendImplicit方法注册）
         *  添加原因：不在\Hyperf\Validation\Validator::implicitRules内的规则，在值是null和空字符串时，不会做验证
         *  一般用于array规则可为空数组的情况，其他规则一般使用require就行了
         *  示例：
         *      'attr' => 'array'，当attr传空字符串''时，不会报错
         *      'attr' => 'required|array'，当attr传[]时，会报错必须
         *      'attr' => 'required_if_null|array'，这样就能确保是数组
         */
        $validatorFactory->extendImplicit('required_if_null', function ($attribute, $value, $parameters, $validator) {
            if (is_null($value)) {
                return false;
            }
            if (is_string($value) && trim($value) === '') {
                return false;
            }
            return true;
        });
    }
}
