<?php

declare(strict_types=1);

namespace app\module\db\model\system;

use app\module\db\model\AbstractModel;

class SystemAdmin extends AbstractModel
{
    protected string $table = 'system_admin';
    protected string $primaryKey = 'adminId';
}
