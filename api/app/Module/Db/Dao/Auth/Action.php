<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $actionId 权限操作ID
 * @property int $pid 父ID（主要用于归类，方便查看。否则可以不要）
 * @property string $actionName 名称
 * @property string $actionCode 标识（代码中用于判断权限）
 * @property string $pidPath 层级路径
 * @property int $level 层级
 * @property string $remark 备注
 * @property int $sort 排序值（从小到大排序，默认50，范围0-100）
 * @property int $isStop 是否停用：0否 1是
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Action extends AbstractDao
{
    /**
     * 解析field（独有的）
     *
     * @param string $key
     * @return boolean
     */
    protected function fieldOfAlone(string $key): bool
    {
        switch ($key) {
            case 'sceneIdArr':
                $this->afterField[] = 'sceneIdArr';

                //需要id字段
                $this->field['select'][] = $this->getTable() . '.' . $this->getKey();
                return true;
        }
        return false;
    }

    /**
     * 获取数据后，再处理的字段（独有的）
     *
     * @param string $key
     * @param object $info
     * @return boolean
     */
    protected function afterFieldOfAlone(string $key, object &$info): bool
    {
        switch ($key) {
            case 'sceneIdArr':
                $info->sceneIdArr = getDao(ActionRelToScene::class)->where(['actionId' => $info->{$this->getKey()}])->getBuilder()->pluck('sceneId')->toArray();
                return true;
        }
        return false;
    }
}
