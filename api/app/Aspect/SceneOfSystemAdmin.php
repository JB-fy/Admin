<?php

declare(strict_types=1);

namespace app\aspect;

use Hyperf\Di\Annotation\Aspect;
use Hyperf\Di\Aop\ProceedingJoinPoint;
use Hyperf\HttpServer\Contract\RequestInterface;

#[Aspect]
class SceneOfSystemAdmin extends AbstractAspect
{
    //执行优先级（大值优先）
    public ?int $priority = 19;

    //要切入的类，可以多个，亦可通过 :: 标识到具体的某个方法，通过 * 可以模糊匹配
    public array $classes = [
        \App\Controller\Login::class . '::info',
        \App\Controller\Login::class . '::menuTree',
        \App\Controller\Auth\AuthScene::class
    ];

    //要切入的注解，具体切入的还是使用了这些注解的类，仅可切入类注解和类方法注解
    public array $annotations = [];

    /**
     * @param ProceedingJoinPoint $proceedingJoinPoint
     * @return void
     */
    public function process(ProceedingJoinPoint $proceedingJoinPoint)
    {
        $request = $this->container->get(RequestInterface::class);
        try {
            /* if ($request->authSceneInfo->sceneCode == 'systemAdmin') {
                container(\App\Module\Service\Login::class)->verifyToken('systemAdmin');
            } */
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
