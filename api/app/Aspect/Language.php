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
