<?php

declare(strict_types=1);

namespace App\Middleware;

use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\RequestHandlerInterface;

class Scene implements \Psr\Http\Server\MiddlewareInterface
{
    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        $logicAuthScene = getContainer()->get(\App\Module\Logic\Auth\Scene::class);
        $sceneCode = $logicAuthScene->getCurrentSceneCode();
        if (empty($sceneCode)) {
            throwFailJson(39999999);
        }
        $sceneInfo = getConfig('inDb.authScene.' . $sceneCode);
        if (empty($sceneInfo)) {
            throwFailJson(39999999);
        }
        if ($sceneInfo->isStop) {
            throwFailJson(39999998);
        }
        $logicAuthScene->setCurrentInfo($sceneInfo);
        try {
            $response = $handler->handle($request);
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
