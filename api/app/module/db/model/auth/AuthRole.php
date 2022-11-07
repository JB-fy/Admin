<?php

declare(strict_types=1);

namespace app\module\db\model\auth;

use app\module\db\model\AbstractModel;

class AuthRole extends AbstractModel
{
    protected $table = 'auth_role';
    protected $primaryKey = 'roleId';
}
