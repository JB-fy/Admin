<?php

declare(strict_types=1);

namespace App\Controller\Platform\Platform;

use App\Controller\AbstractController;

class Config extends AbstractController
{
    /**
     * 获取
     *
     * @return void
     */
    public function get()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $data = $this->validate(__FUNCTION__, $sceneCode);
        $this->checkAuth(__FUNCTION__, $sceneCode);
        $this->service->get(empty($data['configKeyArr']) ? [] : ['configKey' => $data['configKeyArr']]);
    }

    /**
     * 保存
     *
     * @return void
     */
    public function save()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $data = $this->validate(__FUNCTION__, $sceneCode);  //新增配置时，需要在验证文件内新增对应的配置验证。否则数据会被过滤掉
        $this->checkAuth(__FUNCTION__, $sceneCode);
        $this->service->save($data);
    }
}
