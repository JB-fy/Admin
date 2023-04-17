<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Log;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $logId 请求日志ID
 * @property string $requestUrl 请求地址
 * @property string $requestHeader 请求头
 * @property string $requestData 请求数据
 * @property string $responseBody 响应体
 * @property string $runTime 运行时间（单位：毫秒）
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Request extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'log_request';
    protected string $primaryKey = 'logId';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['logId', 'requestUrl', 'requestHeader', 'requestData', 'responseBody', 'runTime', 'updateTime', 'createTime'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['logId' => 'integer'];
}
