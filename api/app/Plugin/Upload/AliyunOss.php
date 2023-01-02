<?php

declare(strict_types=1);

namespace App\Plugin\Upload;

class AliyunOss extends AbstractUpload
{
    /* protected $config = [    //类生成实例后必须含有以下几个字段
        'accessId' => 'LTAI5tHx81H64BRJA971DPZF',
        'accessKey' => 'nJyNpTtUuIgZqx21FF4G2zi0WHOn51',
        'host' => 'http://4724382110.oss-cn-hongkong.aliyuncs.com'
    ]; */

    /**
     * 创建签名
     *
     * @param array $option
     * @return void
     */
    public function createSign(array $option = [])
    {
        /*--------初始化配置 开始--------*/
        mt_srand();
        $defaultOption = [
            'isCallback' => true, //是否回调服务器
            'expireTime' => 15 * 60, //签名有效时间。单位：秒
            'dir' => 'common/' . date('Y/m/d/His') . mt_rand(1000, 9999) . '_',    //上传的文件前缀
            'minSize' => 0,    //限制上传的文件大小。单位：字节
            'maxSize' => 100 * 1024 * 1024,    //限制上传的文件大小。单位：字节
        ];
        $option = array_merge($defaultOption, $option);
        /*--------初始化配置 结束--------*/

        $signInfo = [
            'accessid' => $this->config['accessId'],
            'host' => $this->config['host'],
            'dir' => $option['dir'],
            'expire' => time() + $option['expireTime'],
        ];
        if ($option['isCallback']) {
            $callbackUrl = getRequestUrl() . '/upload/notify';
            $callback_param = [
                'callbackUrl' => $callbackUrl,
                'callbackBody' => 'filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}',
                'callbackBodyType' => 'application/x-www-form-urlencoded'
            ];
            $base64_callback_body = base64_encode(json_encode($callback_param));
            $signInfo['callback'] = $base64_callback_body;
        }

        // $mydatetime = new \DateTime(date('c', $signInfo['expire']));
        // $expiration = $mydatetime->format(\DateTime::ISO8601);
        $expiration = date(DATE_ISO8601, $signInfo['expire']);
        $expiration = substr($expiration, 0, strpos($expiration, '+')) . 'Z';
        $signInfo['policy'] = base64_encode(json_encode([
            'expiration' => $expiration,
            'conditions' => [
                ['content-length-range', $option['minSize'], $option['maxSize']],
                ['starts-with', '$key', $signInfo['dir']]
            ]
        ]));
        $signInfo['signature'] = base64_encode(hash_hmac('sha1', $signInfo['policy'], $this->config['accessKey'], true));

        throwSuccessJson($signInfo);
    }

    /**
     * 回调
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
            throwFailJson('40000000');
        }
        if ($pubKeyUrlBase64 == '') {
            throwFailJson('40000001');
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
            throwFailJson('40000002');
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
            throwFailJson('40000003');
        }
        $data = $request->post();
        $data['url'] = $this->config['host'] . '/' . $data['filename'];
        throwSuccessJson($data);
    }
}
