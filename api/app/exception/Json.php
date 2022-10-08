<?php

declare(strict_types=1);

namespace app\exception;

class Json extends AbstractException
{
    /**
     * 构造函数
     *
     * @param string $code
     * @param array $data
     * @param string $msg
     */
    public function __construct(string $code = '000000', array $data = [], string $msg = '')
    {
        $responseData = [
            'code' => $code,
            'msg' => $msg ? $msg : trans($code, [], 'code'),
            'data' => $data,
        ];
        $responseBody = json_encode($responseData, JSON_UNESCAPED_UNICODE);
        parent::__construct($responseBody);
    }
}
