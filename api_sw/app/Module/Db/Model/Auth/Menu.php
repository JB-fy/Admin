<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Auth;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $menuId 权限菜单ID
 * @property int $sceneId 权限场景ID（只能是auth_scene表中sceneType为0的菜单类型场景）
 * @property int $pid 父ID
 * @property string $menuName 名称
 * @property int $level 层级
 * @property string $pidPath 层级路径
 * @property string $extraData 额外数据。（json格式：{"title（多语言时设置，未设置以menuName返回）": {"语言标识":"标题",...},"icon": "图标","url": "链接地址",...}）
 * @property int $sort 排序值（从小到大排序，默认50，范围0-100）
 * @property int $isStop 是否停用：0否 1是
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Menu extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'auth_menu';
    protected string $primaryKey = 'menuId';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['menuId', 'sceneId', 'pid', 'menuName', 'level', 'pidPath', 'extraData', 'sort', 'isStop', 'updateTime', 'createTime'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['menuId' => 'integer', 'sceneId' => 'integer', 'pid' => 'integer', 'level' => 'integer', 'sort' => 'integer', 'isStop' => 'integer'];
}
