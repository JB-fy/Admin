<?php

declare(strict_types=1);

namespace App\Plugin\Upload;

abstract class AbstractUpload
{
    protected $config = [];

    public function __construct(array $config)
    {
        $this->config = $config;
    }

    /**
     * 创建签名
     *
     * @param array $option
     * @return void
     */
    abstract public function createSign(array $option = []);

    /**
     * 回调
     *
     * @return void
     */
    abstract public function notify();
}
