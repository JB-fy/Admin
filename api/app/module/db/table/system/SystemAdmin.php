<?php

declare(strict_types=1);

namespace app\module\db\table\system;

use app\module\db\table\AbstractTable;
use DI\Annotation\Inject;
use app\module\db\table\auth\AuthRole;
use app\module\db\table\auth\AuthRoleRelOfSystemAdmin;

class SystemAdmin extends AbstractTable
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
                $tableAuthRoleRelOfSystemAdmin = container(AuthRoleRelOfSystemAdmin::class, true);
                $tableAuthRoleRelOfSystemAdminName = $tableAuthRoleRelOfSystemAdmin->getTable();
                if (!isset($this->join[$tableAuthRoleRelOfSystemAdminName])) {
                    $this->join[$tableAuthRoleRelOfSystemAdminName] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $tableAuthRoleRelOfSystemAdminName,
                            $tableAuthRoleRelOfSystemAdminName . '.adminId',
                            '=',
                            $this->getTable() . '.' . $this->getPrimaryKey()
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
                            $tableAuthRoleRelOfSystemAdminName . '.roleId'
                        ]
                    ];
                }
                return true;
        }
        return false;
    }
}
