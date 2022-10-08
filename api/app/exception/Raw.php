<?php

declare(strict_types=1);

namespace app\exception;

class Raw extends AbstractException
{
    /**
     * 构造函数
     *
     * @param string $raw
     */
    public function __construct(string $raw)
    {
        parent::__construct($raw);
    }
}
