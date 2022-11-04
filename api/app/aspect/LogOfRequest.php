<?php

declare(strict_types=1);

namespace app\aspect;

use app\module\db\table\system\SystemLogOfRequest;
use Hyperf\Di\Aop\ProceedingJoinPoint;

class LogOfRequest extends AbstractAspect
{
    /**
     * 执行优先级（大值优先）
     *
     * @var integer
     */
    public $priority = 30;

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
        $startTime = microtime(true);
        try {
            $response = $proceedingJoinPoint->process();
            $responseData = json_encode($response, JSON_UNESCAPED_UNICODE);
            /*if ($response instanceof Response) {
                $responseData = json_encode($response->getData(), JSON_UNESCAPED_UNICODE);
            } else {
                $responseData = json_encode($response, JSON_UNESCAPED_UNICODE);
            }*/
            return $response;
        } catch (\Throwable $e) {
            if ($e instanceof \app\exception\Json) {
                $responseData = $e->getResponseData();
                //unset($responseData['data']['list']); //列表数据太大,记录会给数据库太大压力
                $responseData = json_encode($responseData, JSON_UNESCAPED_UNICODE);
            /* } elseif ($e instanceof \app\exception\Raw) {
                $responseData = json_encode($e->getMessage(), JSON_UNESCAPED_UNICODE); */
            } else {
                $responseData = json_encode($e->getMessage(), JSON_UNESCAPED_UNICODE);
            }
            throw $e;
        } finally {
            $endTime = microtime(true);
            $this->logRequest($startTime, $endTime, $responseData);
        }
    }

    /**
     * 日志记录
     *
     * @param float $startTime
     * @param float $endTime
     * @param string $responseData
     * @return void
     */
    public function logRequest(float $startTime, float $endTime, string $responseData)
    {
        /* $request = request();
        $LogData = [
            'requestUrl' => $request->url(),
            'requestData' => json_encode($request->all(), JSON_UNESCAPED_UNICODE),
            'requestHeaders' => json_encode($request->header(), JSON_UNESCAPED_UNICODE),
            'runTime' => round(($endTime - $startTime) * 1000, 3),
            'responseData' => $responseData,
        ];
        container(SystemLogOfRequest::class, true)->add($LogData); */
    }
}
