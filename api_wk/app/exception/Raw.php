<?php

declare(strict_types=1);

namespace app\exception;

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
     * 获取response
     *
     * @return \support\Response
     */
    final public function getResponse(): \support\Response
    {
        return response($this->raw);
    }
}
