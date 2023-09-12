<?php

declare(strict_types=1);

namespace App\Plugin\Upload;

class Local extends AbstractUpload
{
    /**
     * 上传
     *
     * @return void
     */
    public function upload()
    {
        $request = getRequest();
        $data = $request->post();

        if (time() > $data['expire']) {
            throwFailJson(79999999, '签名过期');
        }
        $signData = [
            'dir' => $data['dir'],
            'expire' => $data['expire'],
            'minSize' => $data['minSize'],
            'maxSize' => $data['maxSize'],
            'rand' => $data['rand'],
        ];
        if ($data['sign'] != $this->createSign($signData)) {
            throwFailJson(79999999, '签名错误');
        }

        $file = $request->file('file');
        $fileSize = $file->getSize();
        if ($data['minSize'] > 0 && $data['minSize'] >  $fileSize) {
            throwFailJson(79999999, '文件不能小于' . ($data['minSize'] / (1024 * 1024)) . 'MB');
        }
        if ($data['maxSize'] > 0 && $data['maxSize'] <  $fileSize) {
            throwFailJson(79999999, '文件不能大于' . ($data['maxSize'] / (1024 * 1024)) . 'MB');
        }

        $dirPath = $this->config['fileSaveDir'] . $data['dir'];
        if (!is_dir($dirPath)) {
            mkdir($dirPath, 0777, true);
        }
        if (empty($data['key'])) {
            $filename = $data['dir'] . round(microtime(true) * 1000) . '_' . mt_rand(10000000, 99999999) . '.' . $file->getExtension();
        } else {
            $filename = $data['key'];
        }
        $filePath = $this->config['fileSaveDir'] . $filename;
        $file->moveTo($filePath);
        if (!$file->isMoved()) {
            throwFailJson(79999999, '文件保存失败，请检查目录及权限');
        }
        chmod($filePath, 0666);

        $resData = [
            'url' => $this->config['fileUrlPrefix'] . '/' . $filename
        ];
        throwSuccessJson($resData);
    }

    /**
     * 签名
     *
     * @param array $option
     * @return void
     */
    public function sign($uploadFileType = '')
    {
        switch ($uploadFileType) {
            default:
                $option = [
                    'dir' => 'common/' . date('Ymd') . '/',    //上传的文件目录
                    'expire' => time() + 15 * 60, //签名有效时间戳。单位：秒
                    'minSize' => 0,    //限制上传的文件大小。单位：字节
                    'maxSize' => 100 * 1024 * 1024,    //限制上传的文件大小。单位：字节。需要同时设置配置文件api_sw/config/autoload/server.php中的OPTION_UPLOAD_MAX_FILESIZE字段
                ];
                break;
        }

        $signInfo = [
            'uploadUrl' => $this->config['url'],
            'host' => $this->config['fileUrlPrefix'],
            'dir' => $option['dir'],
            'expire' => $option['expire'],
            'isRes' =>  1,
        ];

        $uploadData = [
            'dir' =>     $option['dir'],
            'expire' =>  $option['expire'],
            'minSize' => $option['minSize'],
            'maxSize' => $option['maxSize'],
            'rand' =>    randStr(8),
        ];
        $uploadData['sign'] = $this->createSign($uploadData);

        $signInfo['uploadData'] = $uploadData;
        throwSuccessJson($signInfo);
    }

    /**
     * 回调（web前端直传用）
     *
     * @return void
     */
    public function notify()
    {
    }

    /**
     * 生成签名
     *
     * @param array $signData
     * @return string
     */
    protected function createSign(array $signData): string
    {
        sort($signData);
        $str = '';
        foreach ($signData as $key => $value) {
            $str .= $key . '=' . $value . '&';
        }
        $str .= 'key=' . $this->config['signKey'];
        return md5($str);
    }
}
