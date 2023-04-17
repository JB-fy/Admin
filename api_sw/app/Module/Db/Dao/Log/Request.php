<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Log;

use App\Module\Db\Dao\AbstractDao;

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
class Request extends AbstractDao
{
    /**
     * 解析where（独有的）
     *
     * @param string $key
     * @param string|null $operator
     * @param [type] $value
     * @param string|null $boolean
     * @return boolean
     */
    protected function whereOfAlone(string $key, string $operator = null, $value, string $boolean = null): bool
    {
        switch ($key) {
            case 'minRunTime':
                $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.runTime', $operator ?? '>=', $value, $boolean ?? 'and']];
                return true;
            case 'maxRunTime':
                $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.runTime', $operator ?? '<=', $value, $boolean ?? 'and']];
                return true;
        }
        return false;
    }
}
