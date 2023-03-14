<?php

declare(strict_types=1);

namespace app\module\service\auth;

use app\module\service\AbstractService;

class AuthScene extends AbstractService
{
    protected $dao = \app\module\db\dao\auth\AuthScene::class;
}
