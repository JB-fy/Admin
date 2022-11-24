<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;
use Hyperf\Di\Annotation\Inject;

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
class Action extends AbstractDao
{
    #[Inject(value: \App\Module\Db\Model\Auth\Action::class)]
    protected $model;
}
