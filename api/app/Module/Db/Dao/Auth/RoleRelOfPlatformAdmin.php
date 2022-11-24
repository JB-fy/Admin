<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;
use Hyperf\Di\Annotation\Inject;

/**
 * @property int $roleId 权限角色ID
 * @property int $adminId 平台管理员ID
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class RoleRelOfPlatformAdmin extends AbstractDao
{
    #[Inject(value: \App\Module\Db\Model\Auth\RoleRelOfPlatformAdmin::class)]
    protected $model;
}
