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
        $request->authScene = $request->header('AuthScene');
        if (empty($request->authScene)) {
            throwFailJson('001001');
        }
        if (!container(TableAuthScene::class, true)->parseWhere(['sceneCode' => $request->authScene])->getBuilder()->exists()) {
            throwFailJson('001001');
        }
        try {
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $e) {
            throw $e;
        }
    }
}
