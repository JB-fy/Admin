<?php

declare(strict_types=1);

namespace app\module\db\table\system;

use app\module\db\table\AbstractTable;
use DI\Annotation\Inject;

class SystemLogOfRequest extends AbstractTable
{
    /**
     * @Inject
     * @var \app\module\db\model\system\SystemLogOfRequest
     */
    protected $model;
}
