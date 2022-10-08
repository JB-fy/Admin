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
     * 解析field
     *
     * @param array $field  格式：['字段',...]
     * @return self
     */
    public function parseField(array $field): self
    {
        $defaultField = ['id', '*', ...$this->getAllColumn()];
        $intersectField = array_intersect($field, $defaultField);
        parent::parseField($intersectField);    //默认字段交由父类方法处理

        $diffField = array_diff($field, $defaultField);
        foreach ($diffField as $v) {
            $this->parseJoin($v);
            switch ($v) {
                case 'roleName':
                    $this->field['select'][] = container(AuthRole::class, true)->getTableAlias() . '.' . $v;
                    break;
            }
        }
        return $this;
    }

    /**
     * 解析join
     *
     * @param string $key   键，用于确定关联表
     * @param [type] $value 值，用于确定关联表
     * @return void
     */
    public function parseJoin(string $key, $value = null)
    {
        switch ($key) {
            case 'roleName':
                $tableAuthRoleRelOfSystemAdmin = container(AuthRoleRelOfSystemAdmin::class, true);
                $tableAuthRoleRelOfSystemAdminAlias = $tableAuthRoleRelOfSystemAdmin->getTableAlias();
                if (!isset($this->join[$tableAuthRoleRelOfSystemAdminAlias])) {
                    $this->join[$tableAuthRoleRelOfSystemAdminAlias] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $tableAuthRoleRelOfSystemAdmin->getTable() . ' AS ' . $tableAuthRoleRelOfSystemAdminAlias,
                            $tableAuthRoleRelOfSystemAdminAlias . '.adminId',
                            '=',
                            $this->getTableAlias() . '.' . $this->getPrimaryKey()
                        ]
                    ];
                }
                $tableAuthRole = container(AuthRole::class, true);
                $tableAuthRoleAlias = $tableAuthRole->getTableAlias();
                if (!isset($this->join[$tableAuthRoleAlias])) {
                    $this->join[$tableAuthRoleAlias] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $tableAuthRole->getTable() . ' AS ' . $tableAuthRoleAlias,
                            $tableAuthRoleAlias . '.roleId',
                            '=',
                            $tableAuthRoleRelOfSystemAdminAlias . '.roleId'
                        ]
                    ];
                }
                break;
        }
    }
}
