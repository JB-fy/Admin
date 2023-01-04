<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Di\Annotation\Aspect;
use Hyperf\Di\Aop\ProceedingJoinPoint;
use Hyperf\HttpServer\Contract\RequestInterface;

#[Aspect]
class LogRequest extends AbstractAspect
{
    //执行优先级（大值优先）
    public ?int $priority = 30;

    /**
     * @param ProceedingJoinPoint $proceedingJoinPoint
     * @return void
     */
    public function process(ProceedingJoinPoint $proceedingJoinPoint)
    {
        try {
            $startTime = microtime(true);
            $response = $proceedingJoinPoint->process();
            $responseBody = json_encode($response, JSON_UNESCAPED_UNICODE);
            /*if ($response instanceof Response) {
                $responseBody = json_encode($response->getData(), JSON_UNESCAPED_UNICODE);
            } else {
                $responseBody = json_encode($response, JSON_UNESCAPED_UNICODE);
            }*/
            return $response;
        } catch (\Throwable $th) {
            if ($th instanceof \App\Exception\Json) {
                $responseBody = $th->getResponseBody();
                /* $responseData = $th->getResponseData();
                //unset($responseData['data']['list']); //列表数据太大,记录会给数据库太大压力
                $responseBody = json_encode($responseData, JSON_UNESCAPED_UNICODE); */
            } elseif ($th instanceof \App\Exception\Raw) {
                $responseBody = $th->getResponseBody();
            } elseif ($th instanceof \Hyperf\Validation\ValidationException) {
                $responseData = [
                    'code' => '89999999',
                    'msg' => $th->validator->errors()->first(),
                    'data' => [],
                ];
                $responseBody = json_encode($responseData, JSON_UNESCAPED_UNICODE);
            } else {
                $responseBody = json_encode($th->getMessage(), JSON_UNESCAPED_UNICODE);
            }
            throw $th;
        } finally {
            $endTime = microtime(true);
            $this->logRequest($startTime, $endTime, $responseBody);
        }
    }

    /**
     * 日志记录
     *
     * @param float $startTime
     * @param float $endTime
     * @param string $responseBody
     * @return void
     */
    public function logRequest(float $startTime, float $endTime, string $responseBody)
    {
        $request = $this->container->get(RequestInterface::class);

        $LogData = [
            'requestUrl' => getRequestUrl(1),
            'requestData' => json_encode($request->all(), JSON_UNESCAPED_UNICODE),
            'requestHeader' => json_encode($request->getHeaders(), JSON_UNESCAPED_UNICODE),
            'runTime' => round(($endTime - $startTime) * 1000, 3),
            'responseBody' => $responseBody,
        ];
        getDao(\App\Module\Db\Dao\Log\Request::class)->insert($LogData)->saveInsert();
    }
}
