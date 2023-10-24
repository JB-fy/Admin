<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $actionId 权限操作ID
 * @property string $actionName 名称
 * @property string $actionCode 标识（代码中用于判断权限）
 * @property string $remark 备注
 * @property int $isStop 停用：0否 1是
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Action extends AbstractDao
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
            case 'sceneIdArr':
                $this->builder->addSelect($this->getTable() . '.' . $this->getKey()); //需要id字段
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
            case 'sceneIdArr':
                $info->{$key} = getDao(ActionRelToScene::class)->parseFilter(['actionId' => $info->{$this->getKey()}])->getBuilder()->pluck('sceneId')->toArray();
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
            case 'sceneId':
                $joinAlias = getDao(ActionRelToScene::class)->getTable();
                if (is_array($value)) {
                    if (count($value) === 1) {
                        $this->builder->where($joinAlias . '.' . $key, $operator ?? '=', array_shift($value), $boolean ?? 'and');
                    } else {
                        $this->builder->whereIn($joinAlias . '.' . $key, $value, $boolean ?? 'and');
                    }
                } else {
                    $this->builder->where($joinAlias . '.' . $key, $operator ?? '=', $value, $boolean ?? 'and');
                }
                $this->parseJoin($joinAlias);
                break;
            case 'selfAction': //获取当前登录身份可用的操作。参数：['sceneCode'=>场景标识, 'sceneId'=>场景id, 'loginId'=>登录身份id]
                $this->builder->where($this->getTable() . '.isStop', '=', 0, 'and');
                $this->builder->where(getDao(ActionRelToScene::class)->getTable() . '.sceneId', '=', $value['sceneCode'], 'and');
                $this->parseJoin(getDao(ActionRelToScene::class)->getTable());
                $this->parseGroup(['id']);
                switch ($value['sceneCode']) {
                    case 'platform':
                        if ($value['loginId'] === getConfig('app.superPlatformAdminId')) { //平台超级管理员，不再需要其它条件
                            break;
                        }
                        $this->builder->where(getDao(Role::class)->getTable() . '.isStop', '=', 0, 'and');
                        $this->parseJoin(getDao(RoleRelToAction::class)->getTable());
                        $this->parseJoin(getDao(Role::class)->getTable());

                        $this->builder->where(getDao(RoleRelOfPlatformAdmin::class)->getTable() . '.adminId', '=', $value['loginId'], 'and');
                        $this->parseJoin(getDao(RoleRelOfPlatformAdmin::class)->getTable());
                        break;
                }
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
            case getDao(ActionRelToScene::class)->getTable():
                $this->builder->leftJoin($joinAlias, $joinAlias . '.actionId', '=', $this->getTable() . '.' . $this->getKey());
                break;
            case getDao(RoleRelToAction::class)->getTable():
                $this->builder->leftJoin($joinAlias, $joinAlias . '.actionId', '=', $this->getTable() . '.' . $this->getKey());
                break;
            case getDao(Role::class)->getTable():
                $this->builder->leftJoin($joinAlias, $joinAlias . '.roleId', '=', getDao(RoleRelToAction::class)->getTable() . '.roleId');
                break;
            case getDao(RoleRelOfPlatformAdmin::class)->getTable():
                $this->builder->leftJoin($joinAlias, $joinAlias . '.roleId', '=', getDao(RoleRelToAction::class)->getTable() . '.roleId');
                break;
            default:
                parent::parseJoinOne($joinAlias);
        }
    }
}
