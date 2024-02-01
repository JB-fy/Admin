<?php

declare(strict_types=1);

namespace App\Controller\Platform\Login;

use App\Controller\AbstractController;

class PlatformAdmin extends AbstractController
{
    /**
     * 获取登录加密盐(前端登录操作用于加密密码后提交)
     *
     * @return void
     */
    public function salt()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $data = $this->validate(__FUNCTION__, $sceneCode);
        $this->service->salt($data['loginName']);
    }

    /**
     * 登录
     *
     * @return void
     */
    public function login()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $data = $this->validate(__FUNCTION__, $sceneCode);
        $this->service->login($data['loginName'], $data['password']);
    }
}
