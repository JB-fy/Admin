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
        $uploadType = $this->request->input('uploadType');
        /**
         * @var \App\Plugin\Upload\AbstractUpload
         */
        $upload = make('upload');
        $upload->sign($this->createUploadOption($uploadType));
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

    /**
     * 生成各个通用的上传选项
     *
     * @param string $uploadType
     * @return array
     */
    public function createUploadOption($uploadType = ''): array
    {
        $option = [
            'dir' => 'common/' . date('Ymd') . '/',    //上传的文件目录
            'expire' => time() + 15 * 60, //签名有效时间戳。单位：秒
            'minSize' => 0,    //限制上传的文件大小。单位：字节
            'maxSize' => 100 * 1024 * 1024,    //限制上传的文件大小。单位：字节。本地上传（UploadOfLocal.php）需要同时设置配置文件api_sw/config/autoload/server.php中的OPTION_UPLOAD_MAX_FILESIZE字段
        ];
        /* switch ($uploadType) {
            case 'value':
                $option['dir'] = 'common/' . date('Ymd') . '/';
                $option['expire'] = time() + 15 * 60;
                $option['minSize'] = 0;
                $option['maxSize'] = 100 * 1024 * 1024;
                break;
        } */
        return $option;
    }
}
