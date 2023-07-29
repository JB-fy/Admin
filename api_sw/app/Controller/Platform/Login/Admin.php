<?php

declare(strict_types=1);

namespace App\Controller\Platform\Login;

use App\Controller\AbstractController;
use Hyperf\Di\Annotation\Inject;

class Admin extends AbstractController
{
    #[Inject(value: \App\Module\Service\Login\Login::class)]
    protected $service;

    #[Inject(value:\App\Module\Validation\Login\Login::class)]
    protected \App\Module\Validation\AbstractValidation $validation;

    /**
     * 获取登录加密字符串(前端登录操作用于加密密码后提交)
     *
     * @return void
     */
    public function salt()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $data = $this->validate(__FUNCTION__, $sceneCode);
        $this->service->salt($data['account'], $sceneCode);
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
        $this->service->PlatformAdmin($data['account'], $data['password']);
    }
}
