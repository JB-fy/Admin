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
 * @property string $pidPath 层级路径
 * @property string $extraData 额外数据。（json格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}）
 * @property int $sort 排序值（从小到大排序，默认50，范围0-100）
 * @property int $isStop 是否停用：0否 1是
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Menu extends AbstractDao
{
    protected array $jsonField = ['extraData']; //json类型字段。这些字段创建|更新时，需要特殊处理

    /**
     * 解析update（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @return boolean
     */
    protected function parseUpdateOfAlone(string $key, $value = null): bool
    {
        switch ($key) {
            case 'pidPathOfChild':  //更新所有子孙级的pidPath。参数：['newVal'=>父级新pidPath, 'oldVal'=>父级旧pidPath]
                $this->update[$this->getTable() . '.pidPath'] = Db::raw('REPLACE(' . $this->getTable() . '.pidPath, \'' . $value['oldVal'] . '\', \'' . $value['newVal'] . '\')');
                return true;
            case 'levelOfChild':    //更新所有子孙级的level。参数：['newVal'=>父级新level, 'oldVal'=>父级旧level]
                $this->update[$this->getTable() . '.level'] = Db::raw($this->getTable() . '.level + ' . ($value['newVal'] - $value['oldVal']));
                return true;
        }
        return false;
    }

    /**
     * 解析field（独有的）
     *
     * @param string $key
     * @return boolean
     */
    protected function parseFieldOfAlone(string $key): bool
    {
        switch ($key) {
            case 'menuTree':    //树状需要以下字段和排序方式
                $this->builder->addSelect($this->getTable() . '.' . $this->getKey());
                $this->builder->addSelect($this->getTable() . '.' . 'pid');

                $this->parseOrderOfAlone('menuTree');    //排序方式
                return true;
            case 'showMenu':    //前端显示菜单需要以下字段，且title需要转换
                $this->builder->addSelect($this->getTable() . '.' . 'menuName');
                $this->builder->addSelect($this->getTable() . '.' . 'menuIcon');
                $this->builder->addSelect($this->getTable() . '.' . 'menuUrl');

                $this->builder->addSelect($this->getTable() . '.' . 'extraData->i18n AS i18n');
                //$this->builder->addSelect(Db::raw('JSON_UNQUOTE(JSON_EXTRACT(extraData, "$.i18n")) AS i18n')); //mysql不能直接转成对象返回
                $this->afterField[] = 'showMenu';
                return true;
            case 'sceneName':
                $this->builder->addSelect(getDao(Scene::class)->getTable() . '.' . $key);

                $this->parseJoinOfAlone('scene');
                return true;
            case 'pMenuName':
                $this->builder->addSelect('p_' . $this->getTable() . '.menuName AS pMenuName');

                $this->parseJoinOfAlone('pMenu');
                return true;
        }
        return false;
    }

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
            case 'selfMenu': //获取当前登录身份可用的菜单。参数：['sceneCode'=>场景标识, 'sceneId'=>场景id, 'loginId'=>登录身份id]
                $this->builder->where($this->getTable() . '.sceneId', '=', $value['sceneId'], 'and');
                $this->builder->where($this->getTable() . '.isStop', '=', 0, 'and');
                switch ($value['sceneCode']) {
                    case 'platformAdmin':
                        if ($value['loginId'] === getConfig('app.superPlatformAdminId')) { //平台超级管理员，不再需要其他条件
                            return true;
                        }
                        $this->builder->where(getDao(Role::class)->getTable() . '.isStop', '=', 0, 'and');
                        $this->builder->where(getDao(RoleRelOfPlatformAdmin::class)->getTable() . '.adminId', '=', $value['loginId'], 'and');

                        $this->parseJoinOfAlone('roleRelToMenu');
                        $this->parseJoinOfAlone('role');
                        $this->parseJoinOfAlone('roleRelOfPlatformAdmin');
                        break;
                }

                $this->parseGroup(['id']);
                return true;
        }
        return false;
    }

    /**
     * 解析order（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @return boolean
     */
    protected function parseOrderOfAlone(string $key, $value = null): bool
    {
        switch ($key) {
            case 'menuTree':
                $this->builder->orderBy($this->getTable() . '.' . 'pid', 'ASC');
                $this->builder->orderBy($this->getTable() . '.' . 'sort', 'ASC');
                $this->builder->orderBy($this->getTable() . '.' . 'menuId', 'ASC');
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
            case 'pMenu':
                $pMenuDaoTable = 'p_' . $this->getTable();
                if (!in_array($pMenuDaoTable, $this->joinCode)) {
                    $this->joinCode[] = $pMenuDaoTable;
                    $this->builder->leftJoin($this->getTable() . ' AS ' . $pMenuDaoTable, $pMenuDaoTable . '.menuId', '=', $this->getTable() . '.pid');
                }
                return true;
            case 'roleRelToMenu':
                $roleRelToMenuDao = getDao(RoleRelToMenu::class);
                $roleRelToMenuDaoTable = $roleRelToMenuDao->getTable();
                if (!in_array($roleRelToMenuDaoTable, $this->joinCode)) {
                    $this->joinCode[] = $roleRelToMenuDaoTable;
                    $this->builder->leftJoin($roleRelToMenuDaoTable, $roleRelToMenuDaoTable . '.menuId', '=', $this->getTable() . '.menuId');
                }
                return true;
            case 'role':
                $roleRelToMenuDao = getDao(RoleRelToMenu::class);
                $roleRelToMenuDaoTable = $roleRelToMenuDao->getTable();

                $roleDao = getDao(Role::class);
                $roleDaoTable = $roleDao->getTable();
                if (!in_array($roleDaoTable, $this->joinCode)) {
                    $this->joinCode[] = $roleDaoTable;
                    $this->builder->leftJoin($roleDaoTable, $roleDaoTable . '.roleId', '=', $roleRelToMenuDaoTable . '.roleId');
                }
                return true;
            case 'roleRelOfPlatformAdmin':
                $roleRelToMenuDao = getDao(RoleRelToMenu::class);
                $roleRelToMenuDaoTable = $roleRelToMenuDao->getTable();

                $roleRelOfPlatformAdminDao = getDao(RoleRelOfPlatformAdmin::class);
                $roleRelOfPlatformAdminDaoTable = $roleRelOfPlatformAdminDao->getTable();
                if (!in_array($roleRelOfPlatformAdminDaoTable, $this->joinCode)) {
                    $this->joinCode[] = $roleRelOfPlatformAdminDaoTable;
                    $this->builder->leftJoin($roleRelOfPlatformAdminDaoTable, $roleRelOfPlatformAdminDaoTable . '.roleId', '=', $roleRelToMenuDaoTable . '.roleId');
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
            case 'showMenu':
                $info->i18n = $info->i18n ? json_decode($info->i18n, true) : ['title' => ['zh-cn' => $info->menuName]];
                return true;
        }
        return false;
    }
}
