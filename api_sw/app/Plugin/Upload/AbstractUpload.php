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
     * 上传
     *
     * @return void
     */
    abstract public function upload();

    /**
     * 签名
     *
     * @param array $option
     * @return void
     */
    abstract public function sign($uploadFileType = '');

    /**
     * 回调
     *
     * @return void
     */
    abstract public function notify();
}
