<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Auth;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $sceneId 权限场景ID
 * @property string $sceneCode 标识（代码中用于识别调用接口的所在场景，做对应的身份鉴定及权力鉴定。如已在代码中使用，不建议更改）
 * @property string $sceneName 名称
 * @property string $sceneConfig 配置。JSON格式：{"signType": "算法","signKey": "密钥","expireTime": 过期时间,...}
 * @property int $isStop 停用：0否 1是
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class Scene extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'auth_scene';
    protected string $primaryKey = 'sceneId';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['sceneId', 'sceneCode', 'sceneName', 'sceneConfig', 'isStop', 'updatedAt', 'createdAt'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['sceneId' => 'integer', 'isStop' => 'integer'];
}
