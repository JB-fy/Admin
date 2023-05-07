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
            $response = $handler->handle($request);
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
