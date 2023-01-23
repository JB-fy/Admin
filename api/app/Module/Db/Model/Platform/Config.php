<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Platform;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $configId 配置ID
 * @property string $configKey 配置项Key
 * @property string $configValue 配置项值
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
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
    protected array $fillable = ['configId', 'configKey', 'configValue', 'updateTime', 'createTime'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['configId' => 'integer'];
}
