<?php

declare(strict_types=1);

namespace App\Middleware;

use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\RequestHandlerInterface;

class Language implements \Psr\Http\Server\MiddlewareInterface
{
    #[\Hyperf\Di\Annotation\Inject]
    protected \Psr\Container\ContainerInterface $container;

    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        $language = $this->container->get(\Hyperf\HttpServer\Contract\RequestInterface::class)->header('Language', 'zh-cn');
        $this->container->get(\Hyperf\Contract\TranslatorInterface::class)->setLocale($language);

        try {
            $request = \Hyperf\Context\Context::get(\Psr\Http\Message\ServerRequestInterface::class);
            $response = $handler->handle($request);
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
