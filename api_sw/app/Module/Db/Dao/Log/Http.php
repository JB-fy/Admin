<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Log;

use App\Module\Db\Dao\AbstractDao;

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
class Http extends AbstractDao
{
    /**
     * 解析filter（独有的）
     *
     * @param string $key
     * @param string|null $operator
     * @param [type] $value
     * @param string|null $boolean
     * @return boolean
     */
    protected function parseFilterOfAlone(string $key, string $operator = null, $value, string $boolean = null): bool
    {
        switch ($key) {
            case 'minRunTime':
                $this->builder->where($this->getTable() . '.runTime', $operator ?? '>=', $value, $boolean ?? 'and');
                return true;
            case 'maxRunTime':
                $this->builder->where($this->getTable() . '.runTime', $operator ?? '<=', $value, $boolean ?? 'and');
                return true;
        }
        return false;
    }
}
