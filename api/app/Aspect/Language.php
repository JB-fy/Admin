<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Contract\TranslatorInterface;
use Hyperf\Di\Annotation\Aspect;
use Hyperf\Di\Aop\ProceedingJoinPoint;
use Hyperf\HttpServer\Contract\RequestInterface;

#[Aspect]
class Language extends AbstractAspect
{
    //执行优先级（大值优先）
    public ?int $priority = 40;

    //要切入的类，可以多个，亦可通过 :: 标识到具体的某个方法，通过 * 可以模糊匹配
    public array $classes = [
        \App\Controller\Index::class,
        \App\Controller\Login::class,
        \App\Controller\Auth\Menu::class,
        \App\Controller\Auth\Scene::class,
    ];

    //要切入的注解，具体切入的还是使用了这些注解的类，仅可切入类注解和类方法注解
    public array $annotations = [];

    /**
     * @param ProceedingJoinPoint $proceedingJoinPoint
     * @return void
     */
    public function process(ProceedingJoinPoint $proceedingJoinPoint)
    {
        $language = $this->container->get(RequestInterface::class)->header('Language', 'zh-cn');
        $this->container->get(TranslatorInterface::class)->setLocale($language);

        try {
            $result = $proceedingJoinPoint->process();
            return $result;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
