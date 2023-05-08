<?php

declare(strict_types=1);

namespace App\Module\Logic\Auth;

use App\Module\Logic\AbstractLogic;

class Scene extends AbstractLogic
{
    /**
     * 在当前请求中，获取场景标识
     * 
     * @return string|null
     */
    public function getCurrentSceneCode(): ?string
    {
        return getRequest()->header('Scene');
    }

    /**
     * 在当前请求中，设置场景信息
     * 
     * @param object $info
     * @return void
     */
    public function setCurrentInfo(object $info)
    {
        \Hyperf\Context\Context::set('sceneInfo', $info);
    }

    /**
     * 在当前请求中，获取场景信息
     * 
     * @return object|null
     */
    public function getCurrentInfo(): object|null
    {
        return \Hyperf\Context\Context::get('sceneInfo');
    }
}
