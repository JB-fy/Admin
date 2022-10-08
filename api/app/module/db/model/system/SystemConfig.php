<?php

declare(strict_types=1);

namespace app\module\db\model\system;

use app\module\db\model\AbstractModel;

class SystemConfig extends AbstractModel
{
    protected string $table = 'system_config';
    protected string $primaryKey = 'configId';
}
