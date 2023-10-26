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
            'maxSize' => 100 * 1024 * 1024,    //限制上传的文件大小。单位：字节。本地上传（Local.php）需要同时设置配置文件api_sw/config/autoload/server.php中的OPTION_UPLOAD_MAX_FILESIZE字段
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
    abstract public function sign(array $option);

    /**
     * 回调
     *
     * @return void
     */
    abstract public function notify();
}
