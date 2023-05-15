<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Platform;

use App\Module\Db\Dao\AbstractDao;
use App\Module\Db\Dao\Auth\RoleRelOfPlatformAdmin;

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
    /**
     * 解析field（独有的）
     *
     * @param string $key
     * @return boolean
     */
    protected function fieldOfAlone(string $key): bool
    {
        switch ($key) {
            case 'roleIdArr':
                //需要id字段
                $this->builder->addSelect($this->getTable() . '.' . $this->getKey());

                $this->afterField[] = $key;
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
            case 'accountOrPhone':
                if (is_numeric($value)) {
                    $this->builder->where($this->getTable() . '.' . 'phone', $operator ?? '=', $value, $boolean ?? 'and');
                } else {
                    $this->builder->where($this->getTable() . '.' . 'account', $operator ?? '=', $value, $boolean ?? 'and');
                }
                return true;
            case 'roleId':
                if (is_array($value)) {
                    if (count($value) === 1) {
                        $this->builder->where(getDao(RoleRelOfPlatformAdmin::class)->getTable() . '.' . $key, $operator ?? '=', array_shift($value), $boolean ?? 'and');
                    } else {
                        $this->builder->whereIn(getDao(RoleRelOfPlatformAdmin::class)->getTable() . '.' . $key, $value, $boolean ?? 'and');
                    }
                } else {
                    $this->builder->where(getDao(RoleRelOfPlatformAdmin::class)->getTable() . '.' . $key, $operator ?? '=', $value, $boolean ?? 'and');
                }

                $this->joinOfAlone('roleRelOfPlatformAdmin');
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
            case 'roleRelOfPlatformAdmin':
                $roleRelOfPlatformAdminDao = getDao(RoleRelOfPlatformAdmin::class);
                $roleRelOfPlatformAdminDaoTable = $roleRelOfPlatformAdminDao->getTable();
                if (!in_array($roleRelOfPlatformAdminDaoTable, $this->joinCode)) {
                    $this->joinCode[] = $roleRelOfPlatformAdminDaoTable;
                    $this->builder->leftJoin($roleRelOfPlatformAdminDaoTable, $roleRelOfPlatformAdminDaoTable . '.adminId', '=', $this->getTable() . '.' . $this->getKey());
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
            case 'roleIdArr':
                $info->{$key} = getDao(RoleRelOfPlatformAdmin::class)->where(['adminId' => $info->{$this->getKey()}])->getBuilder()->pluck('roleId')->toArray();
                return true;
        }
        return false;
    }
}
