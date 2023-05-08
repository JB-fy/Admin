<?php

declare(strict_types=1);

namespace App\Middleware;

use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\RequestHandlerInterface;

class ScenePlatformAdmin implements \Psr\Http\Server\MiddlewareInterface
{
    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        $sceneCode = 'platformAdmin';
        $sceneInfo = getConfig('inDb.authScene.' . $sceneCode);
        if (empty($sceneInfo)) {
            throwFailJson(39999999);
        }
        if ($sceneInfo->isStop) {
            throwFailJson(39999998);
        }
        $logicAuthScene = getContainer()->get(\App\Module\Logic\Auth\Scene::class);
        $logicAuthScene->setCurrentSceneInfo($sceneInfo);
        try {
            $response = $handler->handle($request);
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
