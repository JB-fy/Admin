<?php

declare(strict_types=1);

namespace app\module\db\dao\auth;

use app\module\db\dao\AbstractDao;
use DI\Annotation\Inject;

class AuthRoleRelToMenu extends AbstractDao
{
    /**
     * @Inject
     * @var \app\module\db\model\auth\AuthRoleRelToMenu
     */
    protected $model;
}
