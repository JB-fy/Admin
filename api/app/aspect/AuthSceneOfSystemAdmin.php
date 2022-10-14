<?php

declare(strict_types=1);

namespace app\aspect;

use app\module\service\Login;
use Hyperf\Di\Aop\ProceedingJoinPoint;

class AuthSceneOfSystemAdmin extends AbstractAspect
{
    /**
     * 执行优先级（大值优先）
     *
     * @var integer
     */
    public $priority = 19;

    /**
     * 要切入的类，可以多个，亦可通过 :: 标识到具体的某个方法，通过 * 可以模糊匹配
     *
     * @var array
     */
    public $classes = [
        \app\controller\Login::class . '::getInfo',
        \app\controller\Login::class . '::getMenuTree',
    ];

    /**
     * @param ProceedingJoinPoint $proceedingJoinPoint
     * @return void
     */
    public function process(ProceedingJoinPoint $proceedingJoinPoint)
    {
        $request = request();
        try {
            if ($request->authSceneInfo['sceneCode'] == 'systemAdmin') {
                container(Login::class)->verifyToken('systemAdmin');
            }
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $e) {
            throw $e;
        }
    }
}
