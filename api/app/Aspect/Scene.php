<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Di\Annotation\Aspect;
use Hyperf\Di\Aop\ProceedingJoinPoint;
use Hyperf\HttpServer\Contract\RequestInterface;

#[Aspect]
class Scene extends AbstractAspect
{
    //执行优先级（大值优先）
    public ?int $priority = 20;

    //要切入的类，可以多个，亦可通过 :: 标识到具体的某个方法，通过 * 可以模糊匹配
    public array $classes = [
        \App\Controller\Index::class,
        \App\Controller\Login::class,
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
        $sceneCode = $request->header('Scene');
        /* if (empty($sceneCode)) {
            throwFailJson('001001');
        }
        $request->authSceneInfo = make(\app\module\db\dao\auth\AuthScene::class)->where(['sceneCode' => $sceneCode])->getInfo();
        if (empty($request->authSceneInfo)) {
            throwFailJson('001001');
        }
        $request->authSceneInfo->sceneConfig = json_decode($request->authSceneInfo->sceneConfig, true); */
        try {
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
