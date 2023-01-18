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
        $option = [];
        $type = $this->request->input('type');
        switch ($type) {
            case 'type1':
                /* $option = [
                    'callbackEnable' => true, //是否回调服务器
                    'expireTime' => 5 * 60, //签名有效时间
                    'dir' => 'common/' . date('Y/m/d/His') . mt_rand(1000, 9999) . '_',    //上传的文件前缀
                    'minSize' => 0,    //限制上传的文件大小。单位：字节
                    'maxSize' => 100 * 1024 * 1024,    //限制上传的文件大小。单位：字节
                ]; */
                break;
        }
        /**
         * @var \App\Plugin\Upload\AbstractUpload
         */
        $upload = $this->container->get('upload');
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
        $upload = $this->container->get('upload');
        $upload->notify();
    }
}
