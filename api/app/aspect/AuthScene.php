<?php

declare(strict_types=1);

namespace app\aspect;

use app\module\db\table\auth\AuthScene as TableAuthScene;
use Hyperf\Di\Aop\ProceedingJoinPoint;

class AuthScene extends AbstractAspect
{
    /**
     * 执行优先级（大值优先）
     *
     * @var integer
     */
    public $priority = 20;

    /**
     * 要切入的类，可以多个，亦可通过 :: 标识到具体的某个方法，通过 * 可以模糊匹配
     *
     * @var array
     */
    public $classes = [
        \app\controller\Index::class,
        \app\controller\Login::class
    ];

    /**
     * @param ProceedingJoinPoint $proceedingJoinPoint
     * @return void
     */
    public function process(ProceedingJoinPoint $proceedingJoinPoint)
    {
        $request = request();
        $sceneCode = $request->header('AuthScene');
        if (empty($sceneCode)) {
            throwFailJson('001001');
        }
        $request->authSceneInfo = container(TableAuthScene::class, true)->where(['sceneCode' => $sceneCode])->getInfo();
        if (empty($request->authSceneInfo)) {
            throwFailJson('001001');
        }
        $request->authSceneInfo->sceneConfig = json_decode($request->authSceneInfo->sceneConfig, true);
        try {
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $e) {
            throw $e;
        }
    }
}
