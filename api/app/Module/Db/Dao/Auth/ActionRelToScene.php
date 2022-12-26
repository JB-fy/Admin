<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $actionId 权限操作ID
 * @property int $sceneId 权限场景ID
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class ActionRelToScene extends AbstractDao
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
            case 'actionId':
            case 'sceneId':
                if (is_array($value)) {
                    if (count($value) === 1) {
                        $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.' . $key, $operator ?? '=', $value[0], $boolean ?? 'and']];
                    } else {
                        $this->where[] = ['method' => 'whereIn', 'param' => [$this->getTable() . '.' . $key, $value, $boolean ?? 'and']];
                    }
                } else {
                    $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.' . $key, $operator ?? '=', $value, $boolean ?? 'and']];
                }
                return true;
        }
        return false;
    }
}
