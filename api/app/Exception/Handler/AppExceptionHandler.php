<?php

declare(strict_types=1);
/**
 * This file is part of Hyperf.
 *
 * @link     https://www.hyperf.io
 * @document https://hyperf.wiki
 * @contact  group@hyperf.io
 * @license  https://github.com/hyperf/hyperf/blob/master/LICENSE
 */

namespace App\Exception\Handler;

use Hyperf\Contract\StdoutLoggerInterface;
use Hyperf\ExceptionHandler\ExceptionHandler;
use Hyperf\HttpMessage\Stream\SwooleStream;
use Psr\Http\Message\ResponseInterface;
use Throwable;

class AppExceptionHandler extends ExceptionHandler
{
    public function __construct(protected StdoutLoggerInterface $logger)
    {
    }

    public function handle(Throwable $throwable, ResponseInterface $response)
    {
        //$response = $response->withHeader('Server', 'Hyperf');    //有默认值：swoole-http-server
        if ($throwable instanceof \App\Exception\Json) {
            // 阻止异常冒泡
            $this->stopPropagation();
            $responseBody = $throwable->getResponseBody();
            //return \Hyperf\Utils\ApplicationContext::getContainer()->get(\Hyperf\HttpServer\Contract\ResponseInterface::class)->json($throwable->getResponseData());
            return $response->withHeader('Content-Type', 'application/json; charset=utf-8')->withBody(new SwooleStream($responseBody));
        } elseif ($throwable instanceof \App\Exception\Raw) {
            // 阻止异常冒泡
            $this->stopPropagation();
            $responseBody = $throwable->getResponseBody();
            //return \Hyperf\Utils\ApplicationContext::getContainer()->get(\Hyperf\HttpServer\Contract\ResponseInterface::class)->raw($responseBody);
            return $response->withHeader('Content-Type', 'text/plain; charset=utf-8')->withBody(new SwooleStream($responseBody));
        } elseif ($throwable instanceof \Hyperf\Validation\ValidationException) {
            // 阻止异常冒泡
            $this->stopPropagation();
            $responseData = [
                'code' => '000999',
                'msg' => $throwable->validator->errors()->first(),
                'data' => [],
            ];
            $responseBody = json_encode($responseData, JSON_UNESCAPED_UNICODE);
            return $response->withHeader('Content-Type', 'application/json; charset=utf-8')->withBody(new SwooleStream($responseBody));
        }
        $this->logger->error(sprintf('%s[%s] in %s', $throwable->getMessage(), $throwable->getLine(), $throwable->getFile()));
        $this->logger->error($throwable->getTraceAsString());
        return $response->withStatus(500)->withBody(new SwooleStream('Internal Server Error.'));
    }

    public function isValid(Throwable $throwable): bool
    {
        return true;
    }
}
