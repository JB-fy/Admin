<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Di\Annotation\Aspect;
use Hyperf\Di\Aop\ProceedingJoinPoint;

#[Aspect]
class Scene extends AbstractAspect
{
    //执行优先级（大值优先）
    public ?int $priority = 20;

    //切入的类
    public array $classes = [
        \App\Controller\Login::class,
        \App\Controller\Auth\Action::class,
        \App\Controller\Auth\Menu::class,
        \App\Controller\Auth\Role::class,
        \App\Controller\Auth\Scene::class,
        \App\Controller\Log\Request::class,
        \App\Controller\Platform\Admin::class,
        \App\Controller\Platform\Config::class,
    ];

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

        $sceneInfo = $this->container->get(\App\Module\Logic\Auth\Scene::class)->getInfo($sceneCode);
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
