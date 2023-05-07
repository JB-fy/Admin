<?php

declare(strict_types=1);

namespace App\Middleware;

use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\RequestHandlerInterface;

class SceneOfPlatformAdmin implements \Psr\Http\Server\MiddlewareInterface
{
    #[\Hyperf\Di\Annotation\Inject]
    protected \Psr\Container\ContainerInterface $container;

    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        try {
            $sceneCode = $this->container->get(\App\Module\Logic\Auth\Scene::class)->getCurrentSceneCode();
            if ($sceneCode == 'platformAdmin') {
                $this->container->get(\App\Module\Service\Login::class)->verifyToken($sceneCode);
            }
            $request = \Hyperf\Context\Context::get(\Psr\Http\Message\ServerRequestInterface::class);   //上面步骤改变过协程上下文中的请求体，故$request需要重置下，否则改变内容会丢失
            $response = $handler->handle($request);
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
