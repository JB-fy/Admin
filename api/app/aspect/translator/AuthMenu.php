<?php

declare(strict_types=1);

namespace app\aspect\translator;

use app\aspect\AbstractAspect;
use Hyperf\Di\Aop\ProceedingJoinPoint;

class AuthMenu extends AbstractAspect
{
    /**
     * 执行优先级（大值优先）
     *
     * @var integer
     */
    public $priority = 10;

    /**
     * 要切入的类，可以多个，亦可通过 :: 标识到具体的某个方法，通过 * 可以模糊匹配
     *
     * @var array
     */
    public $classes = [
        \app\module\service\auth\AuthMenu::class . '::getTree'
    ];

    /**
     * @param ProceedingJoinPoint $proceedingJoinPoint
     * @return void
     */
    public function process(ProceedingJoinPoint $proceedingJoinPoint)
    {
        try {
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $e) {
            if ($e instanceof \app\exception\Json) {
                $responseData = json_decode($e->getMessage(), true);
                throw container(\app\exception\Json::class, true, [
                    'code' => $responseData['code'],
                    'data' => $this->trans($responseData['data']),
                    'msg' => $responseData['msg']
                ]);
            }
            throw $e;
        }
    }

    /**
     * 翻译数据
     *
     * @param array $data
     * @return array
     */
    public function trans(array $data): array
    {
        if (array_key_exists('list', $data) && $data['list']) {
            foreach ($data['list'] as $k => $v) {
                $data['list'][$k] = $this->trans($v);;
            }
        }

        if (array_key_exists('tree', $data) && $data['tree']) {
            foreach ($data['tree'] as $k => $v) {
                $data['tree'][$k] = $this->trans($v);;
            }
        }

        if (array_key_exists('children', $data) && $data['children']) {
            foreach ($data['children'] as $k => $v) {
                $data['children'][$k] = $this->trans($v);;
            }
        }

        if (array_key_exists('menuName', $data)) {
            $data['menuNameTrans'] = '';
            if ($data['menuName']) {
                $data['menuNameTrans'] = trans('authMenu.menuName.' . $data['menuName'], [], 'db');
            }
        }
        return $data;
    }
}
