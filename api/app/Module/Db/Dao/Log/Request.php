<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Log;

use App\Module\Db\Dao\AbstractDao;
use Hyperf\Di\Annotation\Inject;

/**
 * @property int $id 请求日志ID
 * @property string $requestUrl 请求地址
 * @property string $requestData 请求数据
 * @property string $requestHeader 请求头
 * @property string $responseBody 响应体
 * @property string $runTime 运行时间（单位：毫秒）
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Request extends AbstractDao
{
    #[Inject(value: \App\Module\Db\Model\Log\Request::class)]
    protected $model;
}
