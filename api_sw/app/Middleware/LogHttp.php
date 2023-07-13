<?php

declare(strict_types=1);

namespace App\Middleware;

use Hyperf\Logger\LoggerFactory;
use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\RequestHandlerInterface;

class LogHttp implements \Psr\Http\Server\MiddlewareInterface
{
    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        try {
            $startTime = microtime(true);
            $response = $handler->handle($request);
            $responseBody = json_encode($response, JSON_UNESCAPED_UNICODE);
            /*if ($response instanceof Response) {
                $responseBody = json_encode($response->getData(), JSON_UNESCAPED_UNICODE);
            } else {
                $responseBody = json_encode($response, JSON_UNESCAPED_UNICODE);
            }*/
            return $response;
        } catch (\Throwable $th) {
            if ($th instanceof \App\Exception\Json) {
                /* $responseData = $th->getResponseData();
                //$responseData['data'] = [];   //不记录data。有时数据大，记录会给数据库太大压力
                unset($responseData['data']['list']);   //list数据大
                unset($responseData['data']['tree']);   //tree数据大
                $responseBody = json_encode($responseData, JSON_UNESCAPED_UNICODE); */
                $responseBody = $th->getResponseBody();
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
            $this->logHttp($startTime, $endTime, $responseBody);
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
    public function logHttp(float $startTime, float $endTime, string $responseBody)
    {
        $request = getRequest();

        $LogData = [
            'url' => getRequestUrl(1),
            // 'header' => json_encode($request->getHeaders(), JSON_UNESCAPED_UNICODE),
            // 'reqData' => json_encode($request->all(), JSON_UNESCAPED_UNICODE),
            'header' => $request->getHeaders(),
            'reqData' => $request->all(),
            'resData' => $responseBody,
            'runTime' => round(($endTime - $startTime) * 1000, 3),
        ];
        getContainer()->get(LoggerFactory::class)->get('http')->info(json_encode($LogData, JSON_UNESCAPED_UNICODE));
    }
}
