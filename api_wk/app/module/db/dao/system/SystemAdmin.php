<?php

declare(strict_types=1);

namespace app\module\db\dao\system;

use app\module\db\dao\AbstractDao;
use DI\Annotation\Inject;
use app\module\db\dao\auth\AuthRole;
use app\module\db\dao\auth\AuthRoleRelOfSystemAdmin;

class SystemAdmin extends AbstractDao
{
    /**
     * @Inject
     * @var \app\module\db\model\system\SystemAdmin
     */
    protected $model;

    /**
     * 解析field（独有的）
     *
     * @param string $key
     * @return boolean
     */
    protected function fieldOfAlone(string $key): bool
    {
        switch ($key) {
            case 'roleName':
                $this->joinOfAlone($key);
                $this->field['select'][] = container(AuthRole::class, true)->getTable() . '.' . $key;
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
            case 'roleName':
                $daoAuthRoleRelOfSystemAdmin = container(AuthRoleRelOfSystemAdmin::class, true);
                $daoAuthRoleRelOfSystemAdminTable = $daoAuthRoleRelOfSystemAdmin->getTable();
                if (!isset($this->join[$daoAuthRoleRelOfSystemAdminTable])) {
                    $this->join[$daoAuthRoleRelOfSystemAdminTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $daoAuthRoleRelOfSystemAdminTable,
                            $daoAuthRoleRelOfSystemAdminTable . '.adminId',
                            '=',
                            $this->getTable() . '.' . $this->getKey()
                        ]
                    ];
                }
                $tableAuthRole = container(AuthRole::class, true);
                $tableAuthRoleName = $tableAuthRole->getTable();
                if (!isset($this->join[$tableAuthRoleName])) {
                    $this->join[$tableAuthRoleName] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $tableAuthRoleName,
                            $tableAuthRoleName . '.roleId',
                            '=',
                            $daoAuthRoleRelOfSystemAdminTable . '.roleId'
                        ]
                    ];
                }
                return true;
        }
        return false;
    }
}
