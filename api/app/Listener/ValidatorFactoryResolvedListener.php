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
         *      添加原因：不在\Hyperf\Validation\Validator::implicitRules内的规则，在值是null和空字符串时，不会做验证
         *      以array规则为例，其他规则同理：
         *          'attr' => 'array'，当attr传空字符串''时，不会报错
         *          'attr' => 'required|array'，当attr传[]时，会报错必须
         *          'attr' => 'required_if_null|array'，这样就很好用
         */
        $validatorFactory->extendImplicit('required_if_null', function ($attribute, $value, $parameters, $validator) {
            return !($value === '' || $value === null);
        });
    }
}
