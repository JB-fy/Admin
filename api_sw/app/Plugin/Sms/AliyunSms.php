<?php

declare(strict_types=1);

namespace App\Plugin\Sms;

use AlibabaCloud\SDK\Dysmsapi\V20170525\Dysmsapi;
use \Exception;
use AlibabaCloud\Tea\Exception\TeaError;
use AlibabaCloud\Tea\Utils\Utils;
use Darabonba\OpenApi\Models\Config;
use AlibabaCloud\SDK\Dysmsapi\V20170525\Models\SendSmsRequest;
use AlibabaCloud\Tea\Utils\Utils\RuntimeOptions;

class AliyunSms extends AbstractSms
{
    protected $config = [];

    /**
     * 发送短信
     *
     * @param string $phone
     * @param string $code
     * @return void
     */
    public function send(string $phone, string $code)
    {
        return $this->sendSms([$phone], json_encode(['code' => $code], JSON_UNESCAPED_UNICODE));
    }

    /**
     * 发送短信(批量)
     *
     * @param array $phoneArr
     * @param string $templateParam
     * @return void
     */
    public function sendSms(array $phoneArr, string $templateParam)
    {
        $client = $this->createClient();
        $sendSmsRequest = new SendSmsRequest([
            'phoneNumbers' => 'your_value',
            'signName' => $this->config['signName'],
            'templateCode' => $this->config['templateCode'],
            'templateParam' => $templateParam,
            // 'templateParam' => '{"code": "1234"}',
        ]);
        try {
            $result = $client->sendSmsWithOptions($sendSmsRequest, new RuntimeOptions([]));
            $result = $result->toMap();
            if (!(isset($result['body']['Code']) && $result['body']['Code'] == 'OK')) {
                throwFailJson(79999999, '阿里云SMS错误：' . $result['body']['Message']);
            }
        } catch (Exception $error) {
            if (!($error instanceof TeaError)) {
                $error = new TeaError([], $error->getMessage(), $error->getCode(), $error);
            }
            $errMsg = Utils::assertAsString($error->message);
            throwFailJson(79999999, '阿里云SMS错误：' . $errMsg);
        }
    }

    /**
     * 使用AK&SK初始化账号Client
     * @param string $accessKeyId
     * @param string $accessKeySecret
     * @return Dysmsapi Client
     */
    public function createClient()
    {
        $config = new Config([
            'accessKeyId' => $this->config['accessKeyId'],
            'accessKeySecret' => $this->config['accessKeySecret']
        ]);
        $config->endpoint = 'dysmsapi.aliyuncs.com';
        return new Dysmsapi($config);
    }
}
