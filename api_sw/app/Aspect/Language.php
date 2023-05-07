<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Contract\TranslatorInterface;
use Hyperf\Di\Annotation\Aspect;
use Hyperf\Di\Aop\ProceedingJoinPoint;
use Hyperf\HttpServer\Contract\RequestInterface;

#[Aspect]
class Language extends \Hyperf\Di\Aop\AbstractAspect
{
    #[\Hyperf\Di\Annotation\Inject]
    protected \Psr\Container\ContainerInterface $container;

    //执行优先级（大值优先）
    public ?int $priority = 40;

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
        \App\Controller\Log\Request::class,
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
        $language = $this->container->get(RequestInterface::class)->header('Language', 'zh-cn');
        $this->container->get(TranslatorInterface::class)->setLocale($language);

        try {
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
