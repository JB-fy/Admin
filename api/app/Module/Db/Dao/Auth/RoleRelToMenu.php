<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $roleId 权限角色ID
 * @property int $menuId 权限菜单ID
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class RoleRelToMenu extends AbstractDao
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'auth_role_rel_to_menu';

    /**
     * The connection name for the model.
     */
    protected ?string $connection = 'default';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['roleId', 'menuId', 'updateTime', 'createTime'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['roleId' => 'integer', 'menuId' => 'integer'];
}
