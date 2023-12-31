<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $sceneId 权限场景ID
 * @property string $sceneCode 标识（代码中用于识别调用接口的所在场景，做对应的身份鉴定及权力鉴定。如已在代码中使用，不建议更改）
 * @property string $sceneName 名称
 * @property string $sceneConfig 配置。JSON格式：{"signType": "算法","signKey": "密钥","expireTime": 过期时间,...}
 * @property int $isStop 停用：0否 1是
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Scene extends AbstractDao
{
    protected array $jsonField = ['sceneConfig']; //json类型字段。这些字段创建|更新时，需要特殊处理

    /**
     * 解析filter（单个）
     *
     * @param string $key
     * @param string|null $operator
     * @param [type] $value
     * @param string|null $boolean
     * @return void
     */
    protected function parseFilterOne(string $key, string $operator = null, $value, string $boolean = null): void
    {
        switch ($key) {
            case 'sceneName':
                if ($operator === null) {
                    $this->builder->where($this->getTable() . '.' . $key, 'like', '%' . $value . '%', $boolean ?? 'and');
                } else {
                    $this->builder->where($this->getTable() . '.' . $key, $operator, $value, $boolean ?? 'and');
                }
                break;
            default:
                parent::parseFilterOne($key, $operator, $value, $boolean);
        }
    }
}
