<?php

declare(strict_types=1);

namespace App\Plugin\IdCard;


class IdCardOfAliyun extends AbstractIdCard
{
    protected $config = [];

    /**
     * 实名认证
     *
     * @param string $idCardName
     * @param string $idCardNo
     * @return array
     */
    public function auth(string $idCardName, string $idCardNo): array
    {
        $client = getHttpClient(['timeout' => 5]);
        $option = [
            'headers' => ['Authorization' => 'APPCODE ' . $this->config['appcode']],
            'query' => ['name' => $idCardName, 'cardno' => $idCardNo]
            // 'json' => ['key' => 'value'],
            // 'form_params' => ['key' => 'value']
        ];
        try {
            $url = $this->config['host'] . $this->config['path'];
            $response = $client->get($url, $option);
            // $response = $client->post($url, $option);
            $resData = $response->getBody()->getContents();
            $resData = json_decode($resData, true);
        } catch (\Throwable $th) {
            throwFailJson(79999999, $th->getMessage());
        }
        if (!(isset($resData['resp']['code']) && $resData['resp']['code'] == 0)) {
            throwFailJson(79999999, $resData['resp']['desc'] ?? '');
        }

        $idCardInfo = [
            'gender' => 0,
            'address' => $resData['data']['address'],
            'birthday' => $resData['data']['birthday'],
        ];
        switch ($resData['data']['sex']) {
            case '男':
                $idCardInfo['gender'] = 1;
                break;
            case '女':
                $idCardInfo['gender'] = 2;
                break;
        }
        return $idCardInfo;
    }
}
