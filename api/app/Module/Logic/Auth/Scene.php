<?php

declare(strict_types=1);

namespace App\Module\Logic\Auth;

use App\Module\Logic\AbstractLogic;
use Hyperf\Context\Context;
use Hyperf\HttpServer\Contract\RequestInterface;

class Scene extends AbstractLogic
{
    /**
     * 在当前请求中，设置场景信息
     * 
     * @param object $info
     * @return void
     */
    public function setInfo(object $info)
    {
        $request = Context::get(\Psr\Http\Message\ServerRequestInterface::class);
        $request = $request->withAttribute('sceneInfo', $info);
        Context::set(\Psr\Http\Message\ServerRequestInterface::class, $request);
    }

    /**
     * 获取当前请求中的场景信息
     * 
     * @return object
     */
    public function getInfo(): object
    {
        return $this->container->get(RequestInterface::class)->getAttribute('sceneInfo');
    }
}
