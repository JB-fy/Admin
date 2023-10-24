<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;
use Hyperf\DbConnection\Db;

/**
 * @property int $menuId 权限菜单ID
 * @property int $sceneId 权限场景ID（只能是auth_scene表中sceneType为0的菜单类型场景）
 * @property int $pid 父ID
 * @property string $menuName 名称
 * @property string $menuIcon 图标
 * @property string $menuUrl 链接
 * @property int $level 层级
 * @property string $idPath 层级路径
 * @property string $extraData 额外数据。（json格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}）
 * @property int $sort 排序值（从小到大排序，默认50，范围0-100）
 * @property int $isStop 停用：0否 1是
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Menu extends AbstractDao
{
    protected array $jsonField = ['extraData']; //json类型字段。这些字段创建|更新时，需要特殊处理

    /**
     * 解析update（单个）
     *
     * @param string $key
     * @param [type] $value
     * @return void
     */
    protected function parseUpdateOne(string $key, $value): void
    {
        switch ($key) {
            case 'idPathOfChild':  //更新所有子孙级的idPath。参数：['newVal'=>父级新idPath, 'oldVal'=>父级旧idPath]
                $this->update[$this->getTable() . '.idPath'] = Db::raw('REPLACE(' . $this->getTable() . '.idPath, \'' . $value['oldVal'] . '\', \'' . $value['newVal'] . '\')');
                break;
            case 'levelOfChild':    //更新所有子孙级的level。参数：['newVal'=>父级新level, 'oldVal'=>父级旧level]
                $this->update[$this->getTable() . '.level'] = Db::raw($this->getTable() . '.level + ' . ($value['newVal'] - $value['oldVal']));
                break;
            default:
                parent::parseUpdateOne($key, $value);
        }
    }

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
            case 'tree':    //树状需要以下字段和排序方式
                $this->builder->addSelect($this->getTable() . '.' . $this->getKey());
                $this->builder->addSelect($this->getTable() . '.' . 'pid');
                $this->parseOrderOne('tree', null);    //排序方式
                break;
            case 'showMenu':    //前端显示菜单需要以下字段，且title需要转换
                $this->builder->addSelect($this->getTable() . '.' . 'menuName');
                $this->builder->addSelect($this->getTable() . '.' . 'menuIcon');
                $this->builder->addSelect($this->getTable() . '.' . 'menuUrl');
                $this->builder->addSelect($this->getTable() . '.' . 'extraData->i18n AS i18n');
                //$this->builder->addSelect(Db::raw('JSON_UNQUOTE(JSON_EXTRACT(extraData, "$.i18n")) AS i18n')); //mysql不能直接转成对象返回
                $this->afterField[] = 'showMenu';
                break;
            case 'sceneName':
                $this->builder->addSelect(getDao(Scene::class)->getTable() . '.' . $key);
                $this->parseJoin(getDao(Scene::class)->getTable());
                break;
            case 'pMenuName':
                $this->builder->addSelect('p_' . $this->getTable() . '.menuName AS pMenuName');
                $this->parseJoin('p_' . $this->getTable());
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
            case 'showMenu':
                $info->i18n = $info->i18n ? json_decode($info->i18n, true) : ['title' => ['zh-cn' => $info->menuName]];
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
            case 'selfMenu': //获取当前登录身份可用的菜单。参数：['sceneCode'=>场景标识, 'sceneId'=>场景id, 'loginId'=>登录身份id]
                $this->builder->where($this->getTable() . '.sceneId', '=', $value['sceneId'], 'and');
                $this->builder->where($this->getTable() . '.isStop', '=', 0, 'and');
                switch ($value['sceneCode']) {
                    case 'platform':
                        if ($value['loginId'] === getConfig('app.superPlatformAdminId')) { //平台超级管理员，不再需要其它条件
                            break;
                        }
                        $this->builder->where(getDao(Role::class)->getTable() . '.isStop', '=', 0, 'and');
                        $this->parseJoin(getDao(RoleRelToMenu::class)->getTable());
                        $this->parseJoin(getDao(Role::class)->getTable());

                        $this->builder->where(getDao(RoleRelOfPlatformAdmin::class)->getTable() . '.adminId', '=', $value['loginId'], 'and');
                        $this->parseJoin(getDao(RoleRelOfPlatformAdmin::class)->getTable());

                        $this->parseGroup(['id']);
                        break;
                }
                break;
            default:
                parent::parseFilterOne($key, $operator, $value, $boolean);
        }
    }

    /**
     * 解析order（单个）
     *
     * @param string $key
     * @param [type] $value
     * @return void
     */
    protected function parseOrderOne(string $key, $value): void
    {
        switch ($key) {
            case 'tree':
                $this->builder->orderBy($this->getTable() . '.' . 'pid', 'ASC');
                $this->builder->orderBy($this->getTable() . '.' . 'sort', 'ASC');
                $this->builder->orderBy($this->getTable() . '.' . 'menuId', 'ASC');
                break;
            default:
                parent::parseOrderOne($key, $value);
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
            case 'p_' . $this->getTable():
                $this->builder->leftJoin($this->getTable() . ' AS ' . $joinAlias, $joinAlias . '.menuId', '=', $this->getTable() . '.pid');
                break;
            case getDao(RoleRelToMenu::class)->getTable():
                $this->builder->leftJoin($joinAlias, $joinAlias . '.menuId', '=', $this->getTable() . '.menuId');
                break;
            case getDao(Role::class)->getTable():
                $this->builder->leftJoin($joinAlias, $joinAlias . '.roleId', '=', getDao(RoleRelToMenu::class)->getTable() . '.roleId');
                break;
            case getDao(RoleRelOfPlatformAdmin::class)->getTable():
                $this->builder->leftJoin($joinAlias, $joinAlias . '.roleId', '=', getDao(RoleRelToMenu::class)->getTable() . '.roleId');
                break;
            default:
                parent::parseJoinOne($joinAlias);
        }
    }
}
