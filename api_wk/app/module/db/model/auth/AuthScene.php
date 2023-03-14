<?php

declare(strict_types=1);

namespace app\module\db\model\auth;

use app\module\db\model\AbstractModel;

class AuthScene extends AbstractModel
{
    protected $table = 'auth_scene';
    protected $primaryKey = 'sceneId';
}
