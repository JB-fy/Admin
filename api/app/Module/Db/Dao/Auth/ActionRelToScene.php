<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $actionId 权限操作ID
 * @property int $sceneId 权限场景ID
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class ActionRelToScene extends AbstractDao
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'auth_action_rel_to_scene';

    /**
     * The connection name for the model.
     */
    protected ?string $connection = 'default';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['actionId', 'sceneId', 'updateTime', 'createTime'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['actionId' => 'integer', 'sceneId' => 'integer'];
}
