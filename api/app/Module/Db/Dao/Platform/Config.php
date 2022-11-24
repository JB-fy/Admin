<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Platform;

use App\Module\Db\Dao\AbstractDao;
use Hyperf\Di\Annotation\Inject;

/**
 * @property int $id 配置ID
 * @property string $configKey 配置项Key
 * @property string $configValue 配置项值（设置大点。以后可能需要保存富文本内容，如公司简介或协议等等）
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Config extends AbstractDao
{
    #[Inject(value: \App\Module\Db\Model\Platform\Config::class)]
    protected $model;
}
