<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Platform;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $id 管理员ID
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
     * The table associated with the model.
     */
    protected ?string $table = 'platform_admin';

    /**
     * The connection name for the model.
     */
    protected ?string $connection = 'default';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['id', 'account', 'phone', 'password', 'nickname', 'avatar', 'isStop', 'updateTime', 'createTime'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['id' => 'integer', 'isStop' => 'integer'];
}
