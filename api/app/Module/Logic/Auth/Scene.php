<?php

declare(strict_types=1);

namespace App\Module\Logic\Auth;

use App\Module\Logic\AbstractLogic;
use Hyperf\Context\Context;
use Psr\Http\Message\ServerRequestInterface;

class Scene extends AbstractLogic
{
    /**
     * 当前请求对象设置场景信息
     * 
     * @param object $sceneInfo
     * @return void
     */
    public function setRequestSceneInfo(object $sceneInfo)
    {
        $request = Context::get(ServerRequestInterface::class);
        $request->sceneInfo = $sceneInfo;
        Context::set(ServerRequestInterface::class, $request); //重新设置请求对象，改变协程上下文内的请求对象
    }
}
