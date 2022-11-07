<?php

declare(strict_types=1);

namespace app\module\db\model\auth;

use app\module\db\model\AbstractModel;

class AuthAction extends AbstractModel
{
    protected $table = 'auth_action';
    protected $primaryKey = 'actionId';
}
