<?php

declare(strict_types=1);

namespace App\Plugin\Upload;

class UploadOfLocal extends AbstractUpload
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
    public function sign(array $option)
    {
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
        ksort($signData);
        $strArr = [];
        foreach ($signData as $key => $value) {
            if (is_array($value) || is_object($value)) {
                // $strArr[] = $key . '=' . json_encode($value, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE | JSON_HEX_TAG | JSON_HEX_APOS | JSON_HEX_AMP | JSON_HEX_QUOT);
                $strArr[] = $key . '=' . json_encode($value, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE | JSON_HEX_AMP);
            } elseif (is_bool($value)) {
                if ($value) {
                    $strArr[] = $key . '=true';
                } else {
                    $strArr[] = $key . '=false';
                }
            } else {
                $strArr[] = $key . '=' . $value;
            }
        }
        $strArr[] = 'key=' . $this->config['signKey'];
        return md5(implode('&', $strArr));
    }
}
