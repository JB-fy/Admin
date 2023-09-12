<?php

declare(strict_types=1);

namespace App\Plugin\Sms;

abstract class AbstractSms
{
    protected $config = [];

    public function __construct(array $config)
    {
        $this->config = $config;
    }


    /**
     * 发送短信
     *
     * @param string $phone
     * @param string $code
     * @return void
     */
    abstract public function send(string $phone, string $code);

    /**
     * 发送短信(批量)
     *
     * @param array $phoneArr
     * @param string $templateParam
     * @return void
     */
    abstract public function sendSms(array $phoneArr, string $templateParam);
}
