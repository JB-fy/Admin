<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;
use Hyperf\Di\Annotation\Inject;

/**
 * @property int $roleId 权限角色ID
 * @property int $actionId 权限操作ID
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class RoleRelToAction extends AbstractDao
{
    #[Inject(value: \App\Module\Db\Model\Auth\RoleRelToAction::class)]
    protected $model;
}
