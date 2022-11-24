<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Platform;

use App\Module\Db\Dao\AbstractDao;
use App\Module\Db\Dao\Auth\Role;
use App\Module\Db\Dao\Auth\RoleRelOfPlatformAdmin;
use Hyperf\Di\Annotation\Inject;

/**
 * @property int $adminId 管理员ID
 * @property string $account 账号
 * @property string $phone 电话号码
 * @property string $password 密码（md5保存）
 * @property string $nickname 昵称
 * @property string $avatar 头像
 * @property int $isStop 是否停用：0否 1是
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Admin extends AbstractDao
{
    #[Inject(value: \App\Module\Db\Model\Platform\Admin::class)]
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
                $this->field['select'][] = getDao(Role::class)->getTable() . '.' . $key;
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
            case 'loginStr':
                if (is_numeric($value)) {
                    $this->where[] = ['method' => 'where', 'param' => ['phone', $operator ?? '=', $value, $boolean ?? 'and']];
                } else {
                    $this->where[] = ['method' => 'where', 'param' => ['account', $operator ?? '=', $value, $boolean ?? 'and']];
                }
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
                $roleRelOfPlatformAdminDao = getDao(RoleRelOfPlatformAdmin::class);
                $roleRelOfPlatformAdminDaoTable = $roleRelOfPlatformAdminDao->getTable();
                if (!isset($this->join[$roleRelOfPlatformAdminDaoTable])) {
                    $this->join[$roleRelOfPlatformAdminDaoTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $roleRelOfPlatformAdminDaoTable,
                            $roleRelOfPlatformAdminDaoTable . '.adminId',
                            '=',
                            $this->getTable() . '.' . $this->getKey()
                        ]
                    ];
                }
                $roleDao = getDao(Role::class);
                $roleDaoTable = $roleDao->getTable();
                if (!isset($this->join[$roleDaoTable])) {
                    $this->join[$roleDaoTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $roleDaoTable,
                            $roleDaoTable . '.roleId',
                            '=',
                            $roleRelOfPlatformAdminDaoTable . '.roleId'
                        ]
                    ];
                }
                return true;
        }
        return false;
    }
}
