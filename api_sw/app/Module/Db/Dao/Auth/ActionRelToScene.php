<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $actionId 权限操作ID
 * @property int $sceneId 权限场景ID
 * @property string $updateAt 更新时间
 * @property string $createAt 创建时间
 */
class ActionRelToScene extends AbstractDao
{
}
