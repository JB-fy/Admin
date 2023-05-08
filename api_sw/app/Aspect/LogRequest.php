<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Di\Aop\ProceedingJoinPoint;

//#[\Hyperf\Di\Annotation\Aspect]
class LogRequest extends \Hyperf\Di\Aop\AbstractAspect
{
    //执行优先级（大值优先）
    public ?int $priority = 30;

    //切入的类
    public array $classes = [
        \App\Controller\Test::class,
        \App\Controller\Index::class,
        \App\Controller\Upload::class,

        \App\Controller\Login::class,
        \App\Controller\Auth\Action::class,
        \App\Controller\Auth\Menu::class,
        \App\Controller\Auth\Role::class,
        \App\Controller\Auth\Scene::class,
        //\App\Controller\Log\Request::class,
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
        $request = getRequest();

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
