<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $roleId 权限角色ID
 * @property int $sceneId 权限场景ID
 * @property int $tableId 关联表ID（0表示平台创建，其他值根据authSceneId对应不同表，表示是哪个表内哪个机构或个人创建）
 * @property string $roleName 名称
 * @property int $isStop 是否停用：0否 1是
 * @property string $updateAt 更新时间
 * @property string $createAt 创建时间
 */
class Role extends AbstractDao
{
    /**
     * 解析field（独有的）
     *
     * @param string $key
     * @return boolean
     */
    protected function parseFieldOfAlone(string $key): bool
    {
        switch ($key) {
            case 'sceneName':
                $this->builder->addSelect(getDao(Scene::class)->getTable() . '.' . $key);

                $this->parseJoinOfAlone('scene');
                return true;
            case 'menuIdArr':
            case 'actionIdArr':
                //需要id字段
                $this->builder->addSelect($this->getTable() . '.' . $this->getKey());

                $this->afterField[] = $key;
                return true;
        }
        return false;
    }

    /**
     * 解析join（独有的）
     *
     * @param string $key   键，用于确定关联表
     * @param [type] $value 值，用于确定关联表
     * @return boolean
     */
    protected function parseJoinOfAlone(string $key, $value = null): bool
    {
        switch ($key) {
            case 'scene':
                $sceneDao = getDao(Scene::class);
                $sceneDaoTable = $sceneDao->getTable();
                if (!in_array($sceneDaoTable, $this->joinCode)) {
                    $this->joinCode[] = $sceneDaoTable;
                    $this->builder->leftJoin($sceneDaoTable, $sceneDaoTable . '.sceneId', '=', $this->getTable() . '.sceneId');
                }
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
            case 'menuIdArr':
                $info->{$key} = getDao(RoleRelToMenu::class)->parseFilter(['roleId' => $info->{$this->getKey()}])->getBuilder()->pluck('menuId')->toArray();
                return true;
            case 'actionIdArr':
                $info->{$key} = getDao(RoleRelToAction::class)->parseFilter(['roleId' => $info->{$this->getKey()}])->getBuilder()->pluck('actionId')->toArray();
                return true;
        }
        return false;
    }
}
