<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Platform;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $serverId 服务器ID
 * @property string $networkIp 公网IP
 * @property string $localIp 内网IP
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Server extends AbstractDao
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'platform_server';

    /**
     * The connection name for the model.
     */
    protected ?string $connection = 'default';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['serverId', 'networkIp', 'localIp', 'updateTime', 'createTime'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['serverId' => 'integer'];
}
