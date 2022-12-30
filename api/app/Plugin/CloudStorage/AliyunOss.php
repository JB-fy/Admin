<?php

declare(strict_types=1);

namespace App\Plugin\CloudStorage;

class AliyunOss extends AbstractCloudStorage
{
    protected $config = [];

    public function __construct($config)
    {
        $this->config = $config;
    }

    /**
     * 创建签名信息
     *
     * @param $option
     */
    public function createSignInfo($option)
    {
        $callbackUrl = ApplicationContext::getContainer()->get(CommonLogic::class)->getBaseUrl() . 'api/aliyunOss/notify';

        $callback_param = [
            'callbackUrl' => $callbackUrl,
            'callbackBody' => 'filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}',
            'callbackBodyType' => 'application/x-www-form-urlencoded'
        ];
        $base64_callback_body = base64_encode(json_encode($callback_param));

        $end = time() + 300; //设置该policy超时时间是300s. 即这个policy过了这个有效时间，将不能访问

        $mydatetime = new \DateTime(date('c', $end));
        $expiration = $mydatetime->format(\DateTime::ISO8601);
        $expiration = substr($expiration, 0, strpos($expiration, '+')) . 'Z';

        $uploadType && in_array($uploadType, array_keys($this->config['savePath'])) ? null : $uploadType = 'default';
        mt_srand();
        switch ($uploadType) {
            case 'icon':
                $dir = $this->config['savePath'][$uploadType] . date('/Ym/dHis') . mt_rand(1000, 9999) . '_';
                //最大文件大小.用户可以自己设置
                $condition = [
                    0 => 'content-length-range',
                    1 => 0,
                    2 => 10 * 1024 * 1024
                ];
                $conditions[] = $condition;
                break;
            default:
                $dir = $this->config['savePath'][$uploadType] . date('/Y/m/d/His') . mt_rand(1000, 9999) . '_';
                //最大文件大小.用户可以自己设置
                $condition = [
                    0 => 'content-length-range',
                    1 => 0,
                    2 => 100 * 1024 * 1024
                ];
                $conditions[] = $condition;
                break;
        }
        //表示用户上传的数据,必须是以$dir开始, 不然上传会失败,这一步不是必须项,只是为了安全起见,防止用户通过policy上传到别人的目录
        $start = [
            0 => 'starts-with',
            1 => '$key',
            2 => $dir
        ];
        $conditions[] = $start;

        $policy = json_encode([
            'expiration' => $expiration,
            'conditions' => $conditions
        ]);
        $base64_policy = base64_encode($policy);
        $string_to_sign = $base64_policy;
        $signature = base64_encode(hash_hmac('sha1', $string_to_sign, $this->config['accessKey'], true));

        $response = [
            'accessid' => $this->config['accessId'],
            'host' => $this->config['host'],
            'policy' => $base64_policy,
            'signature' => $signature,
            'expire' => $end,
            'callback' => $base64_callback_body,
            'dir' => $dir, //这个参数是设置用户上传指定的前缀
        ];
        throw new ApiException('成功', 0, $response);
    }

    /**
     * Oss上传成回调
     *
     * @throws ApiException
     */
    public function notify()
    {
        // 1.获取OSS的签名header和公钥url header
        $request = ApplicationContext::getContainer()->get(RequestInterface::class);
        $authorizationBase64 = $request->getHeader('authorization')[0] ?? '';
        $pubKeyUrlBase64 = $request->getHeader('x-oss-pub-key-url')[0] ?? '';

        if ($authorizationBase64 == '') {
            throw new ApiException('阿里云OSS回调Header签名不能为空', 9999);
        }
        if ($pubKeyUrlBase64 == '') {
            throw new ApiException('阿里云OSS回调Header公钥Url不能为空', 9999);
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
            throw new ApiException('阿里云OSS回调Header公钥Url错误', 9999);
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
            throw new ApiException('阿里云OSS验证签名错误', 9999);
        }
        $data = $request->post();
        $data['filename'] = $this->config['host'] . '/' . $data['filename'];
        throw new ApiException('成功', 0, $data);
    }
}

// use App\Exception\ApiException;
// use App\Model\Logic\CommonLogic;
// use Hyperf\HttpServer\Contract\RequestInterface;
// use Hyperf\Utils\ApplicationContext;

// class AliyunOss
// {
    
// }