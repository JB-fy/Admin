<?php

declare(strict_types=1);

namespace app\exception;

class Json extends AbstractException
{
    protected string $apiCode;
    protected string $apiMsg;
    protected array $apiData;

    /**
     * 构造函数
     *
     * @param string $code
     * @param string $msg
     * @param array $data
     */
    public function __construct(string $code = '000000', string $msg = '', array $data = [])
    {
        $this->apiCode = $code;
        //parent::__construct($msg ? $msg : trans($code, [], 'code'));
        $this->apiMsg = $msg ? $msg : trans($code, [], 'code');
        //$this->apiData = $data;
        $this->setApiData($data);
    }

    /**
     * 获取response
     *
     * @return \support\Response
     */
    final public function getResponse(): \support\Response
    {
        $responseBody = json_encode($this->getResponseData(), JSON_UNESCAPED_UNICODE);
        return response($responseBody)->withHeader('Content-Type', 'application/json');
    }

    /**
     * 获取responseData
     *
     * @return array
     */
    final public function getResponseData(): array
    {
        return [
            'code' => $this->apiCode,
            'msg' => $this->apiMsg,
            'data' => $this->apiData,
        ];
    }

    /**
     * 设置apiData（有时需要拦截数据做特殊处理后，重新设置）
     *
     * @param array $data
     * @return void
     */
    final public function setApiData(array $data)
    {
        $this->apiData = $data;
    }
}
