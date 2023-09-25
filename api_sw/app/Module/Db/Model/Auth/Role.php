<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Auth;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $roleId 权限角色ID
 * @property int $sceneId 权限场景ID
 * @property int $tableId 关联表ID（0表示平台创建，其它值根据authSceneId对应不同表，表示是哪个表内哪个机构或个人创建）
 * @property string $roleName 名称
 * @property int $isStop 停用：0否 1是
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Role extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'auth_role';
    protected string $primaryKey = 'roleId';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['roleId', 'sceneId', 'tableId', 'roleName', 'isStop', 'updatedAt', 'createdAt'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['roleId' => 'integer', 'sceneId' => 'integer', 'tableId' => 'integer', 'isStop' => 'integer'];
}
