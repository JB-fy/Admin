<?php

declare(strict_types=1);

namespace App\Module\Logic\Auth;

use App\Module\Logic\AbstractLogic;

class Scene extends AbstractLogic
{
    /**
     * 在当前请求中，设置场景信息
     * 
     * @param object $sceneInfo
     * @return void
     */
    public function setCurrentSceneInfo(object $sceneInfo)
    {
        \Hyperf\Context\Context::set('sceneInfo', $sceneInfo);
    }

    /**
     * 在当前请求中，获取场景信息
     * 
     * @return object|null
     */
    public function getCurrentSceneInfo(): object|null
    {
        return \Hyperf\Context\Context::get('sceneInfo');
    }

    /**
     * 在当前请求中，获取场景标识
     * 
     * @return string|null
     */
    public function getCurrentSceneCode(): ?string
    {
        $sceneInfo = \Hyperf\Context\Context::get('sceneInfo');
        return $sceneInfo?->sceneCode;
    }
}
