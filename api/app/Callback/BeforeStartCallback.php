<?php

declare(strict_types=1);

namespace App\Callback;

use Hyperf\Di\Annotation\Inject;
use Psr\Container\ContainerInterface;

//必须先在config/autoload/server.php内callbacks字段增加Event::ON_BEFORE_START => [\App\Callback\BeforeStartCallback::class, 'onBeforeStart']
class BeforeStartCallback
{
    #[Inject]
    protected ContainerInterface $container;

    public function onBeforeStart()
    {
        /**--------执行请求日志表分区任务 开始--------**/
        $this->container->get(\App\Crontab\LogRequest::class)->partition();
        /**--------执行请求日志表分区任务 结束--------**/
    }
}
