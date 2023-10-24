<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Platform;

use App\Module\Db\Dao\AbstractDao;
use App\Module\Db\Dao\Auth\RoleRelOfPlatformAdmin;

/**
 * @property int $adminId 管理员ID
 * @property string $account 账号
 * @property string $phone 电话号码
 * @property string $password 密码。md5保存
 * @property string $nickname 昵称
 * @property string $avatar 头像
 * @property int $isStop 停用：0否 1是
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Admin extends AbstractDao
{
    /**
     * 解析insert（单个）
     *
     * @param string $key
     * @param [type] $value
     * @param integer $index
     * @return void
     */
    protected function parseInsertOne(string $key, $value, int $index = 0): void
    {
        switch ($key) {
            case 'phone':
                $this->insertData[$index][$key] = $value;
                if ($value === '') {
                    $this->insertData[$index][$key] = null;
                }
                break;
            case 'account':
                $this->insertData[$index][$key] = $value;
                if ($value === '') {
                    $this->insertData[$index][$key] = null;
                }
                break;
            default:
                parent::parseInsertOne($key, $value, $index);
        }
    }

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
            case 'phone':
                $this->updateData[$key] = $value;
                if ($value === '') {
                    $this->updateData[$key] = null;
                }
                break;
            case 'account':
                $this->updateData[$key] = $value;
                if ($value === '') {
                    $this->updateData[$key] = null;
                }
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
            case 'roleIdArr':
                $this->builder->addSelect($this->getTable() . '.' . $this->getKey()); //需要id字段
                $this->afterField[] = $key;
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
            case 'roleIdArr':
                $info->{$key} = getDao(RoleRelOfPlatformAdmin::class)->parseFilter(['adminId' => $info->{$this->getKey()}])->getBuilder()->pluck('roleId')->toArray();
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
            case 'loginName':
                if (is_numeric($value)) {
                    $this->builder->where($this->getTable() . '.' . 'phone', $operator ?? '=', $value, $boolean ?? 'and');
                } else {
                    $this->builder->where($this->getTable() . '.' . 'account', $operator ?? '=', $value, $boolean ?? 'and');
                }
                break;
            case 'roleId':
                $joinAlias = getDao(RoleRelOfPlatformAdmin::class)->getTable();
                if (is_array($value)) {
                    if (count($value) === 1) {
                        $this->builder->where($joinAlias . '.' . $key, $operator ?? '=', array_shift($value), $boolean ?? 'and');
                    } else {
                        $this->builder->whereIn($joinAlias . '.' . $key, $value, $boolean ?? 'and');
                    }
                } else {
                    $this->builder->where($joinAlias . '.' . $key, $operator ?? '=', $value, $boolean ?? 'and');
                }
                $this->parseJoin($joinAlias);
                break;
            default:
                parent::parseFilterOne($key, $operator, $value, $boolean);
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
            case getDao(RoleRelOfPlatformAdmin::class)->getTable():
                $this->builder->leftJoin($joinAlias, $joinAlias . '.adminId', '=', $this->getTable() . '.' . $this->getKey());
                break;
            default:
                parent::parseJoinOne($joinAlias);
        }
    }
}
