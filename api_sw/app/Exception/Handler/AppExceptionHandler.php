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
    public function __construct(protected StdoutLoggerInterface $logger, protected \Hyperf\ExceptionHandler\Formatter\FormatterInterface $formatter)
    {
    }

    public function handle(Throwable $throwable, ResponseInterface $response)
    {
        //$response = $response->withHeader('Server', 'Hyperf');    //有默认值：swoole-http-server
        if ($throwable instanceof \App\Exception\Json) {
            $this->stopPropagation();   //阻止异常冒泡
            $responseBody = $throwable->getResponseBody();
            //return getContainer()->get(\Hyperf\HttpServer\Contract\ResponseInterface::class)->json($throwable->getResponseData());
            return $response->withHeader('Content-Type', 'application/json; charset=utf-8')->withBody(new SwooleStream($responseBody));
        } elseif ($throwable instanceof \App\Exception\Raw) {
            $this->stopPropagation();   //阻止异常冒泡
            $responseBody = $throwable->getResponseBody();
            //return getContainer()->get(\Hyperf\HttpServer\Contract\ResponseInterface::class)->raw($responseBody);
            return $response->withHeader('Content-Type', 'text/plain; charset=utf-8')->withBody(new SwooleStream($responseBody));
        } elseif ($throwable instanceof \Hyperf\Validation\ValidationException) {
            $this->stopPropagation();   //阻止异常冒泡
            $responseData = [
                'code' => '89999999',
                'msg' => $throwable->validator->errors()->first(),
                'data' => [],
            ];
            $responseBody = json_encode($responseData, JSON_UNESCAPED_UNICODE);
            /* if (!$response->hasHeader('Content-type')) {
                $response = $response->withAddedHeader('Content-type', 'text/plain; charset=utf-8');
            } */
            return $response->withHeader('Content-Type', 'application/json; charset=utf-8')->withBody(new SwooleStream($responseBody));
        } elseif ($throwable instanceof \Hyperf\Database\Exception\QueryException) {
            //当数据库报1062重复索引时的处理
            if (preg_match('/^SQLSTATE.*1062 Duplicate.*\.([^\']*)\'/', $throwable->getMessage(), $matches) === 1) {
                $this->stopPropagation();   //阻止异常冒泡
                $nameKey = 'validation.attributes.' . $matches[1];
                $responseData = [
                    'code' => '29991062',
                    'msg' => trans('code.29991062', ['errField' => trans($nameKey)]),
                    'data' => [],
                ];
                $responseBody = json_encode($responseData, JSON_UNESCAPED_UNICODE);
                return $response->withHeader('Content-Type', 'application/json; charset=utf-8')->withBody(new SwooleStream($responseBody));
            } else {
                $responseData = [
                    'code' => '29999999',
                    'msg' => $throwable->getMessage(),
                    'data' => [],
                ];
                if (!isDev()) {
                    $responseData['msg']=trans('code.29999999');
                }
                $responseBody = json_encode($responseData, JSON_UNESCAPED_UNICODE);
                return $response->withHeader('Content-Type', 'application/json; charset=utf-8')->withBody(new SwooleStream($responseBody));
			}
        } elseif ($throwable instanceof \Hyperf\HttpMessage\Exception\HttpException) {
            $this->stopPropagation();   //阻止异常冒泡
            $this->logger->debug($this->formatter->format($throwable));
            //return $response->withStatus($throwable->getStatusCode())->withBody(new SwooleStream($throwable->getMessage()));
            return $response->withStatus(404)->withHeader('Content-Type', 'text/plain; charset=utf-8')->withBody(new SwooleStream(trans('code.19990404')));
        }
        $this->logger->error(sprintf('%s[%s] in %s', $throwable->getMessage(), $throwable->getLine(), $throwable->getFile()));
        $this->logger->error($throwable->getTraceAsString());
        //return $response->withStatus(500)->withBody(new SwooleStream('Internal Server Error.'));
        return $response->withStatus(500)->withHeader('Content-Type', 'text/plain; charset=utf-8')->withBody(new SwooleStream(trans('code.19990500')));
    }

    public function isValid(Throwable $throwable): bool
    {
        return true;
    }
}
