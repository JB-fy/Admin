<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Auth;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $actionId 权限操作ID
 * @property string $actionName 名称
 * @property string $actionCode 标识（代码中用于判断权限）
 * @property string $remark 备注
 * @property int $isStop 停用：0否 1是
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Action extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'auth_action';
    protected string $primaryKey = 'actionId';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['actionId', 'actionName', 'actionCode', 'remark', 'isStop', 'updatedAt', 'createdAt'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['actionId' => 'integer', 'isStop' => 'integer'];
}
