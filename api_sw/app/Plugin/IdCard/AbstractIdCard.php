<?php

declare(strict_types=1);

namespace App\Plugin\IdCard;

abstract class AbstractIdCard
{
    protected $config = [];

    public function __construct(array $config)
    {
        $this->config = $config;
    }


    /**
     * 实名认证
     *
     * @param string $idCardName
     * @param string $idCardNo
     * @return array
     */
    abstract public function auth(string $idCardName, string $idCardNo): array;
}
