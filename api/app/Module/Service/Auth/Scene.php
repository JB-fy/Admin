<?php

declare(strict_types=1);

namespace App\Module\Service\Auth;

use App\Module\Service\AbstractService;

class Scene extends AbstractService
{
    protected $daoClassName = \App\Module\Db\Dao\Auth\Scene::class;
}
