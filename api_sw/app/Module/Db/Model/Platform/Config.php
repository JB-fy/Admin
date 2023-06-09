<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Platform;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $configId 配置ID
 * @property string $configKey 配置项Key
 * @property string $configValue 配置项值（设置大点。以后可能需要保存富文本内容，如公司简介或协议等等）
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Config extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'platform_config';
    protected string $primaryKey = 'configId';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['configId', 'configKey', 'configValue', 'updatedAt', 'createdAt'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['configId' => 'integer'];
}
