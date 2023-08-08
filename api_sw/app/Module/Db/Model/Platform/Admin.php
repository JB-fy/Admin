<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Platform;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $adminId 管理员ID
 * @property string $phone 电话号码
 * @property string $account 账号
 * @property string $password 密码（md5保存）
 * @property string $salt 加密盐
 * @property string $nickname 昵称
 * @property string $avatar 头像
 * @property int $isStop 停用：0否 1是
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Admin extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'platform_admin';
    protected string $primaryKey = 'adminId';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['adminId', 'phone', 'account', 'password', 'salt', 'nickname', 'avatar', 'isStop', 'updatedAt', 'createdAt'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['adminId' => 'integer', 'isStop' => 'integer'];
}
