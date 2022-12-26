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
            case 'pidPathOfChild':  //更新所有子孙级的pidPath。参数：[父级新pidPath，父级旧pidPath]
                $this->update[$this->getTable() . '.pidPath'] = Db::raw('REPLACE(' . $this->getTable() . '.pidPath, \'' . $value[1] . '\', \'' . $value[0] . '\')');
                return true;
            case 'levelOfChild':    //更新所有子孙级的level。参数：父级新level - 父级旧level。即父级新旧level差值
                $this->update[$this->getTable() . '.level'] = Db::raw($this->getTable() . '.level + ' . $value);
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
                $this->afterField[] = 'showMenu';

                $this->field['select'][] = $this->getTable() . '.' . 'menuName';
                //$this->field['select'][] = Db::raw('JSON_UNQUOTE(JSON_EXTRACT(extraData, "$.title")) AS title'); //不知道怎么直接转成对象返回
                $this->field['select'][] = $this->getTable() . '.' . 'extraData->title AS title';
                $this->field['select'][] = $this->getTable() . '.' . 'extraData->url AS url';
                $this->field['select'][] = $this->getTable() . '.' . 'extraData->icon AS icon';
                return true;
            case 'sceneName':
                $this->joinOfAlone($key);
                $this->field['select'][] = getDao(Scene::class)->getTable() . '.' . $key;
                return true;
            case 'pMenuName':
                $this->joinOfAlone($key);
                $this->field['select'][] = 'p_' . $this->getTable() . '.menuName AS pMenuName';
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
            case 'sceneName':
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
            case 'pMenuName':
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
