<?php

declare(strict_types=1);

namespace app\exception\handler;

use Throwable;
use Webman\Http\Request;
use Webman\Http\Response;

class Handler extends AbstractHandler
{
    public function render(Request $request, Throwable $e): Response
    {
        /* if ($e instanceof \app\exception\AbstractException) {
            return $e->getResponse();
        } */
        if ($e instanceof \app\exception\Json) {
            return $e->getResponse();
        } elseif ($e instanceof \think\exception\ValidateException) {
            $responseData = [
                'code' => '000999',
                'msg' => $e->getMessage(),
                'data' => [],
            ];
            $responseBody = json_encode($responseData, JSON_UNESCAPED_UNICODE);
            return response($responseBody)->withHeader('Content-Type', 'application/json');
        } elseif ($e instanceof \app\exception\Raw) {
            return $e->getResponse();
        }
        return parent::render($request, $e);
    }
}
