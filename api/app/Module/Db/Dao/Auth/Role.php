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
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Role extends AbstractDao
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
            case 'sceneName':
                $this->joinOfAlone($key);
                $this->field['select'][] = getDao(Scene::class)->getTable() . '.' . $key;
                return true;
            case 'menuIdArr':
            case 'actionIdArr':
                $this->afterField[] = $key;

                //需要id字段
                $this->field['select'][] = $this->getTable() . '.' . $this->getKey();
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
            case 'checkAction': //判断是否有操作权限。参数：['actionCode'=>操作标识, 'sceneCode'=>场景标识, 'loginId'=>登录身份id]
                $this->joinOfAlone($key, $value);

                $sceneInfo = getContainer()->get(\App\Module\Logic\Auth\Scene::class)->getCurrentInfo();    //当开启切面\App\Aspect\Scene时有值
                $sceneId = $sceneInfo === null ? getDao(Scene::class)->where(['sceneCode' => $value['sceneCode']])->getBuilder()->value('sceneId') : $sceneInfo->sceneId;
                $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.sceneId', '=', $sceneId, 'and']];
                $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.isStop', '=', 0, 'and']];
                switch ($value['sceneCode']) {
                    case 'platformAdmin':
                        $this->where[] = ['method' => 'where', 'param' => [getDao(RoleRelOfPlatformAdmin::class)->getTable() . '.adminId', '=', $value['loginId'], 'and']];
                        break;
                }
                $actionDaoTable = getDao(Action::class)->getTable();
                $this->where[] = ['method' => 'where', 'param' => [$actionDaoTable . '.actionCode', '=', $value['actionCode'], 'and']];
                $this->where[] = ['method' => 'where', 'param' => [$actionDaoTable . '.isStop', '=', 0, 'and']];
                $this->where[] = ['method' => 'where', 'param' => [getDao(ActionRelToScene::class)->getTable() . '.sceneId', '=', $sceneId, 'and']];
                $sceneDaoTable = getDao(Scene::class)->getTable();
                //$this->where[] = ['method' => 'where', 'param' => [$sceneDaoTable . '.sceneCode', '=', $value['sceneCode'], 'and']];
                $this->where[] = ['method' => 'where', 'param' => [$sceneDaoTable . '.sceneId', '=', $sceneId, 'and']];
                $this->where[] = ['method' => 'where', 'param' => [$sceneDaoTable . '.isStop', '=', 0, 'and']];
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
            case 'checkAction':
                switch ($value['sceneCode']) {
                    case 'platformAdmin':
                        $roleRelOfPlatformAdminDao = getDao(RoleRelOfPlatformAdmin::class);
                        $roleRelOfPlatformAdminDaoTable = $roleRelOfPlatformAdminDao->getTable();
                        if (!isset($this->join[$roleRelOfPlatformAdminDaoTable])) {
                            $this->join[$roleRelOfPlatformAdminDaoTable] = [
                                'method' => 'leftJoin',
                                'param' => [
                                    $roleRelOfPlatformAdminDaoTable,
                                    $roleRelOfPlatformAdminDaoTable . '.roleId',
                                    '=',
                                    $this->getTable() . '.roleId'
                                ]
                            ];
                        }
                        break;
                    default:
                        break;
                }
                $roleRelToActionDao = getDao(RoleRelToAction::class);
                $roleRelToActionDaoTable = $roleRelToActionDao->getTable();
                if (!isset($this->join[$roleRelToActionDaoTable])) {
                    $this->join[$roleRelToActionDaoTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $roleRelToActionDaoTable,
                            $roleRelToActionDaoTable . '.roleId',
                            '=',
                            $this->getTable() . '.roleId'
                        ]
                    ];
                }
                $actionDao = getDao(Action::class);
                $actionDaoTable = $actionDao->getTable();
                if (!isset($this->join[$actionDaoTable])) {
                    $this->join[$actionDaoTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $actionDaoTable,
                            $actionDaoTable . '.actionId',
                            '=',
                            $roleRelToActionDaoTable . '.actionId'
                        ]
                    ];
                }
                $actionRelToSceneDao = getDao(ActionRelToScene::class);
                $actionRelToSceneDaoTable = $actionRelToSceneDao->getTable();
                if (!isset($this->join[$actionRelToSceneDaoTable])) {
                    $this->join[$actionRelToSceneDaoTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $actionRelToSceneDaoTable,
                            $actionRelToSceneDaoTable . '.actionId',
                            '=',
                            $actionDaoTable . '.actionId'
                        ]
                    ];
                }
                $sceneDao = getDao(Scene::class);
                $sceneDaoTable = $sceneDao->getTable();
                if (!isset($this->join[$sceneDaoTable])) {
                    $this->join[$sceneDaoTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $sceneDaoTable,
                            $sceneDaoTable . '.sceneId',
                            '=',
                            $actionRelToSceneDaoTable . '.sceneId'
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
            case 'menuIdArr':
                $info->{$key} = getDao(RoleRelToMenu::class)->where(['roleId' => $info->{$this->getKey()}])->getBuilder()->pluck('menuId')->toArray();
                return true;
            case 'actionIdArr':
                $info->{$key} = getDao(RoleRelToAction::class)->where(['roleId' => $info->{$this->getKey()}])->getBuilder()->pluck('actionId')->toArray();
                return true;
        }
        return false;
    }
}
