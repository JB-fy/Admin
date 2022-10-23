<?php

declare(strict_types=1);

namespace app\aspect;

use Hyperf\Di\Aop\ProceedingJoinPoint;

class Access extends AbstractAspect
{
    /**
     * 执行优先级（大值优先）
     *
     * @var integer
     */
    public $priority = 50;

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
        try {
            $request = request();
            $response = $request->method() == 'OPTIONS' ? response() : $proceedingJoinPoint->process();
            $response->withHeaders([
                'Access-Control-Allow-Credentials' => 'true',
                //'Access-Control-Allow-Origin' => $request->header('Origin', '*'),
                //'Access-Control-Allow-Origin' => 'http://www.xxxx.com',
                'Access-Control-Allow-Origin' => '*',
                //'Access-Control-Allow-Methods' => 'GET, POST, PUT, DELETE, PATCH, OPTIONS',
                'Access-Control-Allow-Methods' => '*',
                //'Access-Control-Allow-Headers' => 'X-Requested-With, Content-Type, Accept, Origin, Authorization',    //如果有自定义头部此处需要加上,为方便直接使用*不限制
                'Access-Control-Allow-Headers' => '*',
            ]);
            return $response;
        } catch (\Throwable $e) {
            //因程序中利用错误处理返回结果，所以跨域时，错误处理返回的$response也要设置跨域，在app\exception\handler\ExceptionHandler中设置。
            //但在app\exception\handler\ExceptionHandler中设置跨域，会导致不管任何错误，处理后都被设置跨域，相当于全站都跨域。
            //如果这样，不如在nginx中直接设置，或在support\Response中全局设置
            /*location / {
                add_header Access-Control-Allow-Credentials true;
                #add_header Access-Control-Allow-Origin $http_origin;
                #add_header Access-Control-Allow-Origin 'http://www.xxxx.com';
                add_header Access-Control-Allow-Origin *;
                #add_header Access-Control-Allow-Methods 'GET, POST, PUT, DELETE, PATCH, OPTIONS';
                add_header Access-Control-Allow-Methods *;
                #add_header Access-Control-Allow-Headers 'X-Requested-With, Content-Type, Accept, Origin, Authorization';
                add_header Access-Control-Allow-Headers *;
                if ($request_method = 'OPTIONS') {
                    return 204;
                }
            } */
            throw $e;
        }
    }
}
