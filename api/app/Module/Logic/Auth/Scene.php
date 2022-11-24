<?php

declare(strict_types=1);

namespace App\Module\Logic\Auth;

use App\Module\Logic\AbstractLogic;
use Hyperf\Context\Context;
use Psr\Http\Message\ServerRequestInterface;

class Scene extends AbstractLogic
{
    /**
     * 设置请求对象场景信息
     * 
     * @param object $sceneInfo
     * @return void
     */
    public function setRequestSceneInfo(object $sceneInfo)
    {
        $request = Context::get(ServerRequestInterface::class);
        //$request->sceneInfo = $sceneInfo; //这样设置，\Hyperf\HttpServer\Contract\RequestInterface类的对象用$request->sceneInfo拿不到数据
        $request = $request->withAttribute('sceneInfo', $sceneInfo);
        Context::set(ServerRequestInterface::class, $request); //重新设置请求对象，改变协程上下文内的请求对象
    }

    /**
     * 获取请求对象场景信息
     * 
     * @return object
     */
    public function getRequestSceneInfo(): object
    {
        $request = Context::get(ServerRequestInterface::class);
        //return $request->sceneInfo;
        return $request->getAttribute('sceneInfo');
    }
}
