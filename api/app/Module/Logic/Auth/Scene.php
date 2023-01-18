<?php

declare(strict_types=1);

namespace App\Module\Logic\Auth;

use App\Module\Logic\AbstractLogic;
use Hyperf\Context\Context;

class Scene extends AbstractLogic
{
    /**
     * 获取对应场景信息
     * 
     * @param string $sceneCode
     * @return object|null
     */
    public function getInfo(string $sceneCode): object|null
    {
        if (env('AUTH_SCENE_DYNAMIC_ENABLE', false)) {
            $allSceneList = make('allSceneList');   //数据库更新会马上生效
        } else {
            $allSceneList = $this->container->get('allSceneList');    //数据库更新需要重启服务才会生效
        }
        return $allSceneList[$sceneCode];
    }

    /**
     * 在当前请求中，获取场景标识
     * 
     * @return string|null
     */
    public function getCurrentSceneCode(): ?string
    {
        return $this->container->get(\Hyperf\HttpServer\Contract\RequestInterface::class)->header('Scene');
    }

    /**
     * 在当前请求中，设置场景信息
     * 
     * @param object $info
     * @return void
     */
    public function setCurrentInfo(object $info)
    {
        $request = Context::get(\Psr\Http\Message\ServerRequestInterface::class);
        $request = $request->withAttribute('sceneInfo', $info);
        Context::set(\Psr\Http\Message\ServerRequestInterface::class, $request);
    }

    /**
     * 在当前请求中，获取场景信息
     * 
     * @return object|null
     */
    public function getCurrentInfo(): object|null
    {
        return $this->container->get(\Hyperf\HttpServer\Contract\RequestInterface::class)->getAttribute('sceneInfo');
    }
}
