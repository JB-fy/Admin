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
     * 获取apiCode
     *
     * @return string
     */
    final public function getApiCode(): string
    {
        return $this->apiCode;
    }

    /**
     * 获取apiData
     *
     * @return string
     */
    final public function getApiMsg(): string
    {
        return $this->apiMsg;
    }

    /**
     * 获取apiCode
     *
     * @return array
     */
    final public function getApiData(): array
    {
        return $this->apiData;
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
