<?php

declare(strict_types=1);

namespace App\Middleware;

use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\RequestHandlerInterface;

class Language implements \Psr\Http\Server\MiddlewareInterface
{
    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        $language = getRequest()->header('Language', 'zh-cn');
        getContainer()->get(\Hyperf\Contract\TranslatorInterface::class)->setLocale($language);

        try {
            $response = $handler->handle($request);
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
