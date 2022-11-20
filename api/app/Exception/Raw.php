<?php

declare(strict_types=1);

namespace App\Exception;

class Raw extends AbstractException
{
    protected string $raw;

    /**
     * 构造函数
     *
     * @param string $raw
     */
    public function __construct(string $raw)
    {
        $this->raw = $raw;
    }

    /**
     * 获取responseData
     *
     * @return string
     */
    final public function getResponseData(): string
    {
        //return trans('msg.' . $this->raw);
        return $this->raw;
    }

    /**
     * 获取responseBody
     *
     * @return string
     */
    final public function getResponseBody(): string
    {
        return $this->getResponseData();
    }
}
