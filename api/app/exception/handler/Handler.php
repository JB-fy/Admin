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
        if ($e instanceof \app\exception\Json) {
            $responseData = [
                'code' => $e->getApiCode(),
                'msg' => $e->getApiMsg(),
                'data' => $e->getApiData(),
            ];
            $responseBody = json_encode($responseData, JSON_UNESCAPED_UNICODE);
            return response($responseBody)->withHeader('Content-Type', 'application/json');
        }
        if ($e instanceof \think\exception\ValidateException) {
            $responseData = [
                'code' => '000999',
                'msg' => $e->getMessage(),
                'data' => [],
            ];
            /* $msg = $e->getError();
            if (is_array($msg)) {
                $responseData['msg'] = trans($responseData['code'], [], 'code');
                $responseData['data'] = $msg;
            } else {
                $responseData['msg'] = $msg;
            } */
            $responseBody = json_encode($responseData, JSON_UNESCAPED_UNICODE);
            return response($responseBody)->withHeader('Content-Type', 'application/json');
        }
        if ($e instanceof \app\exception\Raw) {
            return response($e->getMessage());
        }
        return parent::render($request, $e);
    }
}
