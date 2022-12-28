<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $sceneId 权限场景ID
 * @property string $sceneCode 标识（代码中用于识别调用接口的所在场景，做对应的身份鉴定及权力鉴定。如已在代码中使用，不建议更改）
 * @property string $sceneName 名称
 * @property string $sceneConfig 配置（内容自定义。json格式：{"alg": "算法","key": "密钥","expTime": "签名有效时间",...}）
 * @property int $isStop 是否停用：0否 1是
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Scene extends AbstractDao
{
    /**
     * 解析insert（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @param integer $index
     * @return boolean
     */
    protected function insertOfAlone(string $key, $value, int $index = 0): bool
    {
        switch ($key) {
            case 'sceneConfig':
                if (!($value === '' || $value === null)) {
                    $this->insert[$index][$key] =  is_array($value) ? json_encode($value, JSON_UNESCAPED_UNICODE) : $value;
                }
                return true;
        }
        return false;
    }

    /**
     * 解析update（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @return boolean
     */
    protected function updateOfAlone(string $key, $value = null): bool
    {
        switch ($key) {
            case 'sceneConfig':
                if ($value === '' || $value === null) {
                    $this->update[$this->getTable() . '.' . $key] = null;
                } else {
                    $this->update[$this->getTable() . '.' . $key] = is_array($value) ? json_encode($value, JSON_UNESCAPED_UNICODE) : $value;
                }
                return true;
        }
        return false;
    }

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
            case 'sceneName':
                if ($operator === null) {
                    $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.' . $key, 'like', '%' . $value . '%', $boolean ?? 'and']];
                } else {
                    $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.' . $key, $operator, $value, $boolean ?? 'and']];
                }
                return true;
        }
        return false;
    }
}
