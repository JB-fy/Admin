<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Platform;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $id 配置ID
 * @property string $configKey 配置项Key
 * @property string $configValue 配置项值（设置大点。以后可能需要保存富文本内容，如公司简介或协议等等）
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Config extends AbstractDao
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'platform_config';

    /**
     * The connection name for the model.
     */
    protected ?string $connection = 'default';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['id', 'configKey', 'configValue', 'updateTime', 'createTime'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['id' => 'integer'];
}
