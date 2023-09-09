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
        $type = $this->request->input('type');
        switch ($type) {
            default:
                $option = [
                    'expireTime' => 15 * 60, //签名有效时间
                    'dir' => 'common/' . date('Ymd') . '/',    //上传的文件前缀
                    'minSize' => 0,    //限制上传的文件大小。单位：字节
                    'maxSize' => 100 * 1024 * 1024,    //限制上传的文件大小。单位：字节
                ];
                break;
        }
        /**
         * @var \App\Plugin\Upload\AbstractUpload
         */
        $upload = make('upload');
        $upload->createSign($option);
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
}
