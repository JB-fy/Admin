<?php

declare(strict_types=1);

namespace App\Controller\Platform\My;

use App\Controller\AbstractController;

class Profile extends AbstractController
{
    /**
     * 获取登录用户信息
     *
     * @return void
     */
    public function info()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getCurrentInfo($sceneCode);
        throwSuccessJson(['info' => $loginInfo]);
    }

    /**
     * 修改个人信息
     *
     * @return void
     */
    public function update()
    {
        /**--------参数验证并处理 开始--------**/
        $data = $this->request->all();
        $data = $this->container->get(\App\Module\Validation\Platform\Admin::class)->make($data, 'updateSelf')->validate();
        /**--------参数验证并处理 结束--------**/

        $sceneCode = $this->scene->getCurrentSceneCode();
        $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getCurrentInfo($sceneCode);
        $this->container->get(\App\Module\Service\Platform\Admin::class)->update($data, ['id' => $loginInfo->adminId]);
    }
}