<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Di\Annotation\Aspect;
use Hyperf\Di\Aop\ProceedingJoinPoint;

//这个中间件可以不要，控制器没有对应场景也会报错
//#[Aspect]
class Scene extends AbstractAspect
{
    //执行优先级（大值优先）
    public ?int $priority = 20;

    /**
     * @param ProceedingJoinPoint $proceedingJoinPoint
     * @return void
     */
    public function process(ProceedingJoinPoint $proceedingJoinPoint)
    {
        $sceneCode = $this->container->get(\App\Module\Logic\Auth\Scene::class)->getCurrentSceneCode();
        if (empty($sceneCode)) {
            throwFailJson('39999999');
        }

        $sceneInfo = getContainer()->get(\App\Module\Logic\Auth\Scene::class)->getInfo($sceneCode);
        if (empty($sceneInfo)) {
            throwFailJson('39999999');
        }
        if ($sceneInfo->isStop) {
            throwFailJson('39999998');
        }
        $this->container->get(\App\Module\Logic\Auth\Scene::class)->setCurrentInfo($sceneInfo);
        try {
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
