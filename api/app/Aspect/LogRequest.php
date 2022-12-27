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

    //要切入的类，可以多个，亦可通过 :: 标识到具体的某个方法，通过 * 可以模糊匹配
    public array $classes = [
        \App\Controller\Index::class,
        \App\Controller\Login::class,
        \App\Controller\Auth\Action::class,
        \App\Controller\Auth\Menu::class,
        \App\Controller\Auth\Scene::class,
    ];

    //要切入的注解，具体切入的还是使用了这些注解的类，仅可切入类注解和类方法注解
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

        /* $config = $this->container->get(ConfigInterface::class);
        $requestData = [
            'data' => $request->all()
        ];
        $scene = $request->getHeaderLine($config->get('app.auth.sceneName'));
        $loginInfo = $request->getAttribute($config->get('app.auth.' . $scene . '.infoName'));
        $loginInfo ? $requestData['loginInfo'] = $loginInfo : null; */

        /* $LogData = [
            //'requestUrl' => $this->container->get(CommonLogic::class)->getUrl(),
            'requestUrl' => $request->fullUrl(),
            //'requestData' => json_encode($requestData, JSON_UNESCAPED_UNICODE),
            'requestData' => json_encode($request->all(), JSON_UNESCAPED_UNICODE),
            'requestHeaders' => json_encode($request->getHeaders(), JSON_UNESCAPED_UNICODE),
            'runTime' => round(($endTime - $startTime) * 1000, 3),
            'responseBody' => $responseBody,
        ];
        $this->container->get(\App\Module\Db\Dao\System\LogOfRequest::class, true)->add($LogData); */
    }
}
