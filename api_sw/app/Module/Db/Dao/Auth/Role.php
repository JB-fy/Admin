<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $roleId 权限角色ID
 * @property int $sceneId 权限场景ID
 * @property int $tableId 关联表ID（0表示平台创建，其它值根据authSceneId对应不同表，表示是哪个表内哪个机构或个人创建）
 * @property string $roleName 名称
 * @property int $isStop 停用：0否 1是
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Role extends AbstractDao
{
    /**
     * 解析field（单个）
     *
     * @param string $key
     * @param [type] $value
     * @return void
     */
    protected function parseFieldOne(string $key, $value = null): void
    {
        switch ($key) {
            case 'sceneName':
                $this->builder->addSelect(getDao(Scene::class)->getTable() . '.' . $key);
                $this->parseJoin(getDao(Scene::class)->getTable());
                break;
            case 'menuIdArr':
            case 'actionIdArr':
                //需要id字段
                $this->builder->addSelect($this->getTable() . '.' . $this->getKey());
                $this->afterField[] = $key;
                break;
            case 'tableName':
                $this->builder->addSelect($this->getTable() . '.tableId');
                $this->builder->addSelect(getDao(Scene::class)->getTable() . '.sceneCode');
                $this->parseJoin(getDao(Scene::class)->getTable());
                $this->afterField[] = $key;
                break;
            default:
                parent::parseFieldOne($key, $value);
        }
    }

    /**
     * 获取数据后，再处理的字段（单个）
     *
     * @param string $key
     * @param object $info
     * @return void
     */
    protected function parseAfterField(string $key, object &$info): void
    {
        switch ($key) {
            case 'menuIdArr':
                $info->{$key} = getDao(RoleRelToMenu::class)->parseFilter(['roleId' => $info->{$this->getKey()}])->getBuilder()->pluck('menuId')->toArray();
                break;
            case 'actionIdArr':
                $info->{$key} = getDao(RoleRelToAction::class)->parseFilter(['roleId' => $info->{$this->getKey()}])->getBuilder()->pluck('actionId')->toArray();
                break;
            case 'tableName':
                if ($info->tableId == 0) {
                    $info->{$key} = '平台';
                    break;
                }
                switch ($info->sceneCode) {
                    case 'platform':
                        break;
                }
                break;
            default:
                parent::parseAfterField($key, $info);
        }
    }

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
            case 'sceneCode':
                $this->builder->where(getDao(Scene::class)->getTable() . '.' . $key, $operator ?? '=', $value, $boolean ?? 'and');
                $this->parseJoin(getDao(Scene::class)->getTable());
                break;
            default:
                parent::parseFilterOne($key, $operator, $value, $boolean);
        }
    }

    /**
     * 解析join（单个）
     *
     * @param string $joinCode
     * @return void
     */
    protected function parseJoinOne(string $joinAlias): void
    {
        switch ($joinAlias) {
            case getDao(Scene::class)->getTable():
                $this->builder->leftJoin($joinAlias, $joinAlias . '.sceneId', '=', $this->getTable() . '.sceneId');
                break;
            default:
                parent::parseJoinOne($joinAlias);
        }
    }
}
