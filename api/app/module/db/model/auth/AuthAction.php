<?php

declare(strict_types=1);

namespace app\module\db\model\auth;

use app\module\db\model\AbstractModel;

class AuthAction extends AbstractModel
{
    protected string $table = 'auth_action';
    protected string $key = 'actionId';
}
