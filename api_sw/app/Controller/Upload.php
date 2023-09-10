<?php

declare(strict_types=1);

namespace App\Controller;

class Upload extends AbstractController
{
    /**
     * 获取签名
     *
     */
    public function sign()
    {
        $uploadFileType = $this->request->input('type');
        /**
         * @var \App\Plugin\Upload\AbstractUpload
         */
        $upload = make('upload');
        $upload->sign($this->request->input('type'));
    }

    /**
     * 回调
     *
     */
    public function notify()
    {
        /**
         * @var \App\Plugin\Upload\AbstractUpload
         */
        $upload = make('upload');
        $upload->notify();
    }

    /**
     * 上传
     *
     */
    public function upload()
    {
        /**
         * @var \App\Plugin\Upload\AbstractUpload
         */
        $upload = make('upload');
        $upload->upload();
    }
}
