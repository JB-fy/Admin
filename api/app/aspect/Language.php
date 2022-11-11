<?php

declare(strict_types=1);

namespace app\aspect;

use Hyperf\Di\Aop\ProceedingJoinPoint;

class Language extends AbstractAspect
{
    /**
     * 执行优先级（大值优先）
     *
     * @var integer
     */
    public $priority = 40;

    /**
     * 要切入的类，可以多个，亦可通过 :: 标识到具体的某个方法，通过 * 可以模糊匹配
     *
     * @var array
     */
    public $classes = [
        \app\controller\Index::class,
        \app\controller\Login::class,
        \app\controller\auth\AuthScene::class
    ];

    /**
     * @param ProceedingJoinPoint $proceedingJoinPoint
     * @return void
     */
    public function process(ProceedingJoinPoint $proceedingJoinPoint)
    {
        $request = request();
        $request->language = $request->header('Language', 'zh-cn');
        locale($request->language);

        try {
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $e) {
            throw $e;
        }
    }
}
