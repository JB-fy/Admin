<?php

declare(strict_types=1);

namespace app\exception;

abstract class AbstractException extends \Exception
{
    /**
     * 获取response
     *
     * @return \support\Response
     */
    abstract public function getResponse(): \support\Response;
}
