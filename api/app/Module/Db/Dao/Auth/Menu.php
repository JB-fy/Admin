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
 * @property string $pidPath 层级路径
 * @property int $level 层级
 * @property string $extraData 扩展数据。（json格式：{"title（多语言时设置，未设置以menuName返回）": {"语言标识":"标题",...},"icon": "图标","url": "链接地址",...}）
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
    protected function updateOfAlone(string $key, $value = null): bool
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
    protected function fieldOfAlone(string $key): bool
    {
        switch ($key) {
            case 'menuTree':    //树状需要以下字段和排序方式
                $this->field['select'][] = $this->getTable() . '.' . $this->getKey();
                $this->field['select'][] = $this->getTable() . '.' . 'pid';

                $this->orderOfAlone('menuTree');    //排序方式
                return true;
            case 'showMenu':    //前端显示菜单需要以下字段，且title需要转换
                $this->field['select'][] = $this->getTable() . '.' . 'menuName';
                //$this->field['select'][] = Db::raw('JSON_UNQUOTE(JSON_EXTRACT(extraData, "$.title")) AS title'); //不知道怎么直接转成对象返回
                $this->field['select'][] = $this->getTable() . '.' . 'extraData->title AS title';
                $this->field['select'][] = $this->getTable() . '.' . 'extraData->url AS url';
                $this->field['select'][] = $this->getTable() . '.' . 'extraData->icon AS icon';

                $this->afterField[] = 'showMenu';
                return true;
            case 'sceneName':
                $this->field['select'][] = getDao(Scene::class)->getTable() . '.' . $key;

                $this->joinOfAlone('scene');
                return true;
            case 'pMenuName':
                $this->field['select'][] = 'p_' . $this->getTable() . '.menuName AS pMenuName';

                $this->joinOfAlone('pMenu');
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
            case 'selfMenu': //获取当前登录身份可用的菜单。参数：['sceneCode'=>场景标识, 'loginId'=>登录身份id]
                $sceneInfo = getContainer()->get(\App\Module\Logic\Auth\Scene::class)->getCurrentInfo();    //当开启切面\App\Aspect\Scene时有值
                $sceneId = $sceneInfo === null ? getDao(Scene::class)->where(['sceneCode' => $value['sceneCode']])->getBuilder()->value('sceneId') : $sceneInfo->sceneId;
                $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.sceneId', '=', $sceneId, 'and']];
                $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.isStop', '=', 0, 'and']];
                switch ($value['sceneCode']) {
                    case 'platformAdmin':
                        if ($value['loginId'] === getConfig('app.superPlatformAdminId')) { //平台超级管理员，不再需要其他条件
                            return true;
                        }
                        $this->where[] = ['method' => 'where', 'param' => [getDao(Role::class)->getTable() . '.isStop', '=', 0, 'and']];
                        $this->where[] = ['method' => 'where', 'param' => [getDao(RoleRelOfPlatformAdmin::class)->getTable() . '.adminId', '=', $value['loginId'], 'and']];

                        $this->joinOfAlone('roleRelToMenu');
                        $this->joinOfAlone('role');
                        $this->joinOfAlone('roleRelOfPlatformAdmin');
                        break;
                }

                $this->groupOfCommon('id');
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
    protected function orderOfAlone(string $key, $value = null): bool
    {
        switch ($key) {
            case 'menuTree':
                $this->order[] = ['method' => 'orderBy', 'param' => [$this->getTable() . '.' . 'pid', 'ASC']];
                $this->order[] = ['method' => 'orderBy', 'param' => [$this->getTable() . '.' . 'sort', 'ASC']];
                $this->order[] = ['method' => 'orderBy', 'param' => [$this->getTable() . '.' . 'menuId', 'ASC']];
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
    protected function joinOfAlone(string $key, $value = null): bool
    {
        switch ($key) {
            case 'scene':
                $sceneDao = getDao(Scene::class);
                $sceneDaoTable = $sceneDao->getTable();
                if (!isset($this->join[$sceneDaoTable])) {
                    $this->join[$sceneDaoTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $sceneDaoTable,
                            $sceneDaoTable . '.sceneId',
                            '=',
                            $this->getTable() . '.sceneId'
                        ]
                    ];
                }
                return true;
            case 'pMenu':
                $pMenuDaoTable = 'p_' . $this->getTable();
                if (!isset($this->join[$pMenuDaoTable])) {
                    $this->join[$pMenuDaoTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $this->getTable() . ' AS ' . $pMenuDaoTable,
                            $pMenuDaoTable . '.menuId',
                            '=',
                            $this->getTable() . '.pid'
                        ]
                    ];
                }
                return true;
            case 'roleRelToMenu':
                $roleRelToMenuDao = getDao(RoleRelToMenu::class);
                $roleRelToMenuDaoTable = $roleRelToMenuDao->getTable();
                if (!isset($this->join[$roleRelToMenuDaoTable])) {
                    $this->join[$roleRelToMenuDaoTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $roleRelToMenuDaoTable,
                            $roleRelToMenuDaoTable . '.menuId',
                            '=',
                            $this->getTable() . '.menuId'
                        ]
                    ];
                }
                return true;
            case 'role':
                $roleRelToMenuDao = getDao(RoleRelToMenu::class);
                $roleRelToMenuDaoTable = $roleRelToMenuDao->getTable();

                $roleDao = getDao(Role::class);
                $roleDaoTable = $roleDao->getTable();
                if (!isset($this->join[$roleDaoTable])) {
                    $this->join[$roleDaoTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $roleDaoTable,
                            $roleDaoTable . '.roleId',
                            '=',
                            $roleRelToMenuDaoTable . '.roleId'
                        ]
                    ];
                }
                return true;
            case 'roleRelOfPlatformAdmin':
                $roleRelToMenuDao = getDao(RoleRelToMenu::class);
                $roleRelToMenuDaoTable = $roleRelToMenuDao->getTable();

                $roleRelOfPlatformAdminDao = getDao(RoleRelOfPlatformAdmin::class);
                $roleRelOfPlatformAdminDaoTable = $roleRelOfPlatformAdminDao->getTable();
                if (!isset($this->join[$roleRelOfPlatformAdminDaoTable])) {
                    $this->join[$roleRelOfPlatformAdminDaoTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $roleRelOfPlatformAdminDaoTable,
                            $roleRelOfPlatformAdminDaoTable . '.roleId',
                            '=',
                            $roleRelToMenuDaoTable . '.roleId'
                        ]
                    ];
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
                $info->title = $info->title ? json_decode($info->title, true) : [];
                $info->icon = $info->icon ?? '';
                $info->url = $info->url ?? '';
                return true;
        }
        return false;
    }
}
