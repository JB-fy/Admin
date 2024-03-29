<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Platform;

use App\Module\Db\Model\AbstractModel;

/**
 * @property string $configKey 配置Key
 * @property string $configValue 配置值
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Config extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'platform_config';
    protected string $primaryKey = 'configKey';
    protected string $keyType = 'string';
    public bool $incrementing = false;

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['configKey', 'configValue', 'updatedAt', 'createdAt'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = [];
}
