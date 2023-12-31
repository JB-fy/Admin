<?php

declare(strict_types=1);

namespace App\Plugin\Upload;

class UploadOfAliyunOss extends AbstractUpload
{
    /**
     * 上传
     *
     * @return void
     */
    public function upload()
    {
    }

    /**
     * 创建签名（web前端直传用）
     *
     * @param array $option
     * @return void
     */
    public function sign(array $option)
    {
        $bucketHost = $this->getBucketHost();
        $signInfo = [
            'uploadUrl' => $bucketHost,
            'host' => $bucketHost,
            'dir' => $option['dir'],
            'expire' => $option['expire'],
            'isRes' =>  0,
        ];

        $uploadData = [
            'OSSAccessKeyId' =>        $this->config['accessKeyId'],
            'success_action_status' => '200', //让服务端返回200,不然，默认会返回204
        ];
        $expiration = date(DATE_ISO8601, $signInfo['expire']);
        $expiration = substr($expiration, 0, strpos($expiration, '+')) . 'Z';
        $uploadData['policy'] = base64_encode(json_encode([
            'expiration' => $expiration,
            'conditions' => [
                ['content-length-range', $option['minSize'], $option['maxSize']],
                ['starts-with', '$key', $option['dir']]
            ]
        ]));
        $uploadData['signature'] = base64_encode(hash_hmac('sha1', $uploadData['policy'], $this->config['accessKeySecret'], true));
        //是否回调
        if (!empty($this->config['callbackUrl'])) {
            $callback_param = [
                // 'callbackUrl' => getRequestUrl() . '/upload/notify',
                'callbackUrl' => $this->config['callbackUrl'],
                'callbackBody' => 'filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}',
                'callbackBodyType' => 'application/x-www-form-urlencoded'
            ];
            $base64_callback_body = base64_encode(json_encode($callback_param));
            $uploadData['callback'] = $base64_callback_body;
            $signInfo['isRes'] = 1;
        }

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
        // 1.获取OSS的签名header和公钥url header
        $request = getRequest();
        $authorizationBase64 = $request->getHeader('authorization')[0] ?? '';
        $pubKeyUrlBase64 = $request->getHeader('x-oss-pub-key-url')[0] ?? '';
        if ($authorizationBase64 == '') {
            throwFailJson(79990000);
        }
        if ($pubKeyUrlBase64 == '') {
            throwFailJson(79990001);
        }

        // 2.获取OSS的签名
        $authorization = base64_decode($authorizationBase64);

        // 3.获取公钥
        $pubKeyUrl = base64_decode($pubKeyUrlBase64);
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $pubKeyUrl);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 10);
        $pubKey = curl_exec($ch);
        if ($pubKey == '') {
            throwFailJson(79990002);
        }

        // 4.获取回调body
        $body = $request->getBody()->getContents();

        // 5.拼接待签名字符串
        $path = $request->server('request_uri', '');
        $pos = strpos($path, '?');
        $authStr = ($pos === false ? urldecode($path) . "\n" . $body : urldecode(substr($path, 0, $pos)) . substr($path, $pos, strlen($path) - $pos) . "\n" . $body);

        // 6.验证签名
        $ok = openssl_verify($authStr, $authorization, $pubKey, OPENSSL_ALGO_MD5);
        if ($ok != 1) {
            throwFailJson(79990003);
        }

        $resData = [
            'url' => $this->getBucketHost() . '/' . $request->post('filename')
        ];
        throwSuccessJson($resData);
    }

    /**
     * 获取bucketHost（web前端直传用）
     *
     * @return string
     */
    protected function getBucketHost(): string
    {
        $scheme = strpos($this->config['host'], 'https://') === 0 ? 'https://' : 'http://';
        return substr_replace($this->config['host'], $scheme . $this->config['bucket'] . '.', 0, strlen($scheme));
    }
}
