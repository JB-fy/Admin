<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Di\Annotation\Aspect;
use Hyperf\Di\Aop\ProceedingJoinPoint;

//这个中间件可以不要，控制器没有对应场景也会报错
#[Aspect]
class Scene extends AbstractAspect
{
    //执行优先级（大值优先）
    public ?int $priority = 20;

    //要切入的类，可以多个，亦可通过 :: 标识到具体的某个方法，通过 * 可以模糊匹配
    public array $classes = [
        \App\Controller\Index::class,
        \App\Controller\Login::class,
        \App\Controller\Auth\Scene::class
    ];

    //要切入的注解，具体切入的还是使用了这些注解的类，仅可切入类注解和类方法注解
    public array $annotations = [];

    /**
     * @param ProceedingJoinPoint $proceedingJoinPoint
     * @return void
     */
    public function process(ProceedingJoinPoint $proceedingJoinPoint)
    {
        $sceneCode = getRequestScene();
        if (empty($sceneCode)) {
            throwFailJson('001001');
        }
        $sceneInfo = make(\App\Module\Db\Dao\Auth\Scene::class)->where(['sceneCode' => $sceneCode])->getInfo();
        if (empty($sceneInfo)) {
            throwFailJson('001001');
        }
        $sceneInfo->sceneConfig = json_decode($sceneInfo->sceneConfig, true);
        $this->container->get(\App\Module\Logic\Auth\Scene::class)->setRequestSceneInfo($sceneInfo);
        try {
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
