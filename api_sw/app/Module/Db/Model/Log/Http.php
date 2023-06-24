<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Log;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $httpId Http日志ID
 * @property string $url 地址
 * @property string $header 请求头
 * @property string $reqData 请求数据
 * @property string $resData 响应数据
 * @property string $runTime 运行时间（单位：毫秒）
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Http extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'log_http';
    protected string $primaryKey = 'httpId';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['httpId', 'url', 'header', 'reqData', 'resData', 'runTime', 'updatedAt', 'createdAt'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['httpId' => 'integer'];
}
