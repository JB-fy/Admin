<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Platform;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $configId 配置ID
 * @property string $configKey 配置Key
 * @property string $configValue 配置值
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Config extends AbstractDao
{
}
