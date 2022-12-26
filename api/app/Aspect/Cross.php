<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Context\Context;
use Hyperf\Di\Annotation\Aspect;
use Hyperf\Di\Aop\ProceedingJoinPoint;
use Hyperf\HttpServer\Contract\RequestInterface;
use Psr\Http\Message\ResponseInterface;

#[Aspect]
class Cross extends AbstractAspect
{
    //执行优先级（大值优先）
    public ?int $priority = 50;

    //要切入的类，可以多个，亦可通过 :: 标识到具体的某个方法，通过 * 可以模糊匹配
    public array $classes = [
        \App\Controller\Test::class,
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
        //也可以在nginx中直接设置全站跨域
        /* location / {
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
        $request = $this->container->get(RequestInterface::class);

        /*--------设置协程上限文响应体可跨域  开始--------*/
        $response = Context::get(ResponseInterface::class);
        $response = $response->withHeader('Access-Control-Allow-Credentials', 'true')
            //->withHeader('Access-Control-Allow-Origin', $request->header('Origin', '*'))
            //->withHeader('Access-Control-Allow-Origin', 'http://www.xxxx.com')
            ->withHeader('Access-Control-Allow-Origin', '*')
            //->withHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, PATCH, OPTIONS')
            ->withHeader('Access-Control-Allow-Methods', '*')
            //->withHeader('Access-Control-Allow-Headers', 'X-Requested-With, Content-Type, Accept, Origin, Authorization')   //如果有自定义头部此处需要加上,为方便直接使用*不限制
            ->withHeader('Access-Control-Allow-Headers', '*');
        Context::set(ResponseInterface::class, $response);
        /*--------设置协程上限文响应体可跨域  结束--------*/

        if ($request->getMethod() == 'OPTIONS') {
            return $response;
        }

        try {
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
