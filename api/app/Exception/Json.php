<?php

declare(strict_types=1);

namespace App\Exception;

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
    public function __construct(string $code = '00000000', string $msg = '', array $data = [])
    {
        $this->apiCode = $code;
        $this->apiMsg = $msg === '' ? trans('code.' . $code) : $msg;
        //parent::__construct($msg === '' ? trans('code.' . $code) : $msg);
        //$this->apiData = $data;
        $this->setApiData($data);
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
     * 获取responseBody
     *
     * @return string
     */
    final public function getResponseBody(): string
    {
        return json_encode($this->getResponseData(), JSON_UNESCAPED_UNICODE);
    }
}
