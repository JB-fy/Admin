<?php

declare(strict_types=1);

namespace app\module\db\dao\system;

use app\module\db\dao\AbstractDao;
use DI\Annotation\Inject;

class SystemConfig extends AbstractDao
{
    /**
     * @Inject
     * @var \app\module\db\model\system\SystemConfig
     */
    protected $model;
}
