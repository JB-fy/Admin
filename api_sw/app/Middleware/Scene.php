<?php

declare(strict_types=1);

namespace App\Middleware;

use Hyperf\Di\Annotation\Inject;
use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\RequestHandlerInterface;

class Scene implements \Psr\Http\Server\MiddlewareInterface
{
    #[Inject]
    protected \Psr\Container\ContainerInterface $container;

    #[Inject]
    protected \App\Module\Logic\Auth\Scene $logicAuthScene;

    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        $sceneCode = $this->logicAuthScene->getCurrentSceneCode();
        if (empty($sceneCode)) {
            throwFailJson('39999999');
        }
        $sceneInfo = getConfig('inDb.authScene.' . $sceneCode);
        if (empty($sceneInfo)) {
            throwFailJson('39999999');
        }
        if ($sceneInfo->isStop) {
            throwFailJson('39999998');
        }
        $this->logicAuthScene->setCurrentInfo($sceneInfo);
        try {
            $request = \Hyperf\Context\Context::get(\Psr\Http\Message\ServerRequestInterface::class);
            $response = $handler->handle($request);
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
