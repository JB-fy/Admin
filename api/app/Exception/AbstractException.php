<?php

declare(strict_types=1);

namespace App\Exception;

abstract class AbstractException extends \Exception
{
    /**
     * 获取responseData
     *
     * @return array|string
     */
    abstract public function getResponseData(): array|string;

    /**
     * 获取responseBody
     *
     * @return string
     */
    abstract public function getResponseBody(): string;
}
