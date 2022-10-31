<?php

declare(strict_types=1);

namespace app\module\db\model\system;

use app\module\db\model\AbstractModel;

class SystemLogOfRequest extends AbstractModel
{
    protected string $table = 'system_log_of_request';
    protected string $key = 'requestLogId';
}
