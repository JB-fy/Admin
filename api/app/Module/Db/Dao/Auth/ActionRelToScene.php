<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;
use Hyperf\Di\Annotation\Inject;

/**
 * @property int $actionId 权限操作ID
 * @property int $sceneId 权限场景ID
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class ActionRelToScene extends AbstractDao
{
    #[Inject(value: \App\Module\Db\Model\Auth\ActionRelToScene::class)]
    protected $model;
}
