<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Di\Aop\ProceedingJoinPoint;

//#[\Hyperf\Di\Annotation\Aspect]
class Scene extends \Hyperf\Di\Aop\AbstractAspect
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
        \App\Controller\Platform\Server::class,
    ];

    //切入的注解
    public array $annotations = [];

    /**
     * @param ProceedingJoinPoint $proceedingJoinPoint
     * @return void
     */
    public function process(ProceedingJoinPoint $proceedingJoinPoint)
    {
        $logicAuthScene = getContainer()->get(\App\Module\Logic\Auth\Scene::class);
        $sceneCode = $logicAuthScene->getCurrentSceneCode();
        if (empty($sceneCode)) {
            throwFailJson(39999999);
        }

        $sceneInfo = getConfig('inDb.authScene.' . $sceneCode);
        if (empty($sceneInfo)) {
            throwFailJson(39999999);
        }
        if ($sceneInfo->isStop) {
            throwFailJson(39999998);
        }
        $logicAuthScene->setCurrentInfo($sceneInfo);
        try {
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
