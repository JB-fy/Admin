<?php

declare(strict_types=1);

namespace App\Module\Logic\Auth;

use App\Module\Logic\AbstractLogic;
use Hyperf\Context\Context;
use Hyperf\HttpServer\Contract\RequestInterface;

class Scene extends AbstractLogic
{
    /**
     * 获取对应场景信息
     * 
     * @param string $sceneCode
     * @return object
     */
    public function getInfo(string $sceneCode): object
    {
        //return make($sceneCode . 'SceneInfo');   //数据库更改会变动
        return $this->container->get($sceneCode . 'SceneInfo');    //需要重启服务才会变动
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
     * 获取当前请求中的场景信息
     * 
     * @return object|null
     */
    public function getCurrentInfo(): object|null
    {
        return $this->container->get(RequestInterface::class)->getAttribute('sceneInfo');
    }
}
