<?php

declare(strict_types=1);

namespace App\Plugin;

class Jwt
{
    protected $config = [];

    public function __construct($config)
    {
        $this->config = $config;
    }

    /**
     * 获取配置参数
     * 
     * @return array
     */
    public function getConfig(): array
    {
        return $this->config;
    }

    /**
     * 生成token
     * 
     * @param array $payload
     * @return string
     */
    public function createToken(array $payload): string
    {
        $payload['expireTime'] = time() + $this->config['expireTime'];
        $header = [
            'signType' => $this->config['signType'],
            'type' => 'JWT'
        ];
        $encodeHeader = $this->encode(json_encode($header, JSON_UNESCAPED_UNICODE));
        $encodePayload = $this->encode(json_encode($payload, JSON_UNESCAPED_UNICODE));
        $token = $encodeHeader . '.' . $encodePayload . '.' . $this->encode($this->sign($encodeHeader . '.' . $encodePayload, $this->config['signKey'], $this->config['signType']));
        return $token;
    }

    /**
     * 验证token
     * 
     * @param string $token
     * @return array
     */
    public function verifyToken(string $token): array
    {
        $tokenList = explode('.', $token);
        if (count($tokenList) != 3) {
            throwFailJson(39994001);
        }
        list($encodeHeader, $encodePayload, $sign) = $tokenList;
        //获取算法
        $decodeHeader = json_decode($this->decode($encodeHeader), true);
        if (empty($decodeHeader['signType'])) {
            throwFailJson(39994001);
        }
        //签名验证
        if (self::encode($this->sign($encodeHeader . '.' . $encodePayload, $this->config['signKey'], $decodeHeader['signType'])) !== $sign) {
            throwFailJson(39994001);
        }
        $payload = json_decode($this->decode($encodePayload), true);
        //过期时间小宇当前服务器时间验证失败
        if (isset($payload['expireTime']) && $payload['expireTime'] < time()) {
            throwFailJson(39994001);
        }
        return $payload;
    }

    /**
     * base64编码
     *
     * @param string $str
     * @return string
     */
    protected function encode(string $str): string
    {
        return str_replace('=', '', strtr(base64_encode($str), '+/', '-_'));
    }

    /**
     * base64解码
     *
     * @param string $str
     * @return string
     */
    protected function decode(string $str): string
    {
        $remainder = strlen($str) % 4;
        if ($remainder) {
            $str .= str_repeat('=', 4 - $remainder);
        }
        return base64_decode(strtr($str, '-_', '+/'));
    }

    /**
     * 签名
     *
     * @param string $str
     * @param string $signKey
     * @param string $signType
     * @return string
     */
    protected function sign(string $str, string $signKey, string $signType = 'HS256'): string
    {
        switch ($signType) {
            case 'HS256':
            case 'HS384':
            case 'HS512':
                $sign = hash_hmac(str_replace('HS', 'SHA', $signType), $str, $signKey, true);
                if (!$sign) {
                    throwFailJson(39999002);
                }
                return $sign;
            case 'RS256':
            case 'RS384':
            case 'RS512':
                $sign = '';
                if (!openssl_sign($str, $sign, $signKey, str_replace('HS', 'SHA', $signType))) {
                    throwFailJson(39999002);
                }
                return $sign;
            default:
                throwFailJson(39999003);
        }
    }
}
