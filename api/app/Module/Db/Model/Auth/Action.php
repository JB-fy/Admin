<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Auth;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $id 权限操作ID
 * @property int $pid 父ID（主要用于归类，方便查看。否则可以不要）
 * @property string $actionName 名称
 * @property string $actionCode 标识（代码中用于判断权限）
 * @property string $pidPath 层级路径
 * @property int $level 层级
 * @property string $remark 备注
 * @property int $sort 排序值（从小到大排序，默认50，范围0-100）
 * @property int $isStop 是否停用：0否 1是
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Action extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'auth_action';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['id', 'pid', 'actionName', 'actionCode', 'pidPath', 'level', 'remark', 'sort', 'isStop', 'updateTime', 'createTime'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['id' => 'integer', 'pid' => 'integer', 'level' => 'integer', 'sort' => 'integer', 'isStop' => 'integer'];
}
