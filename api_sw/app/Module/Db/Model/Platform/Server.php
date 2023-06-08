<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Platform;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $serverId 服务器ID
 * @property string $networkIp 公网IP
 * @property string $localIp 内网IP
 * @property string $updateAt 更新时间
 * @property string $createAt 创建时间
 */
class Server extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'platform_server';
    protected string $primaryKey = 'serverId';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['serverId', 'networkIp', 'localIp', 'updateAt', 'createAt'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['serverId' => 'integer'];
}
