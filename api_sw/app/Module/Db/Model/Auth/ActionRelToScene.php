<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Auth;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $actionId 权限操作ID
 * @property int $sceneId 权限场景ID
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class ActionRelToScene extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'auth_action_rel_to_scene';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['actionId', 'sceneId', 'updatedAt', 'createdAt'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['actionId' => 'integer', 'sceneId' => 'integer'];
}
