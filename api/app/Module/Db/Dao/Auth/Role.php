<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;
use Hyperf\Di\Annotation\Inject;

/**
 * @property int $id 权限角色ID
 * @property int $sceneId 权限场景ID
 * @property int $tableId 关联表ID（0表示平台创建，其他值根据authSceneId对应不同表，表示是哪个表内哪个机构或个人创建）
 * @property string $roleName 名称
 * @property int $isStop 是否停用：0否 1是
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Role extends AbstractDao
{
    #[Inject(value: \App\Module\Db\Model\Auth\Role::class)]
    protected $model;
}
