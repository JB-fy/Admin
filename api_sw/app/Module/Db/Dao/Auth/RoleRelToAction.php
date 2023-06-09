<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $roleId 权限角色ID
 * @property int $actionId 权限操作ID
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class RoleRelToAction extends AbstractDao
{
}
