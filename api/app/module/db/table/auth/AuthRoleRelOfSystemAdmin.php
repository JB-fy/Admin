<?php

declare(strict_types=1);

namespace app\module\db\table\auth;

use app\module\db\table\AbstractTable;
use DI\Annotation\Inject;

class AuthRoleRelOfSystemAdmin extends AbstractTable
{
    /**
     * @Inject
     * @var \app\module\db\model\auth\AuthRoleRelOfSystemAdmin
     */
    protected $model;
}
