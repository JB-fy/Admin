<?php

declare(strict_types=1);

namespace app\module\db\dao\auth;

use app\module\db\dao\AbstractDao;
use DI\Annotation\Inject;

class AuthAction extends AbstractDao
{
    /**
     * @Inject
     * @var \app\module\db\model\auth\AuthAction
     */
    protected $model;
}
