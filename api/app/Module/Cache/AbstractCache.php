<?php

declare(strict_types=1);

namespace App\Module\Cache;

use Hyperf\Contract\ConfigInterface;
use Hyperf\Di\Annotation\Inject;
use Psr\Container\ContainerInterface;

abstract class AbstractCache
{
    #[Inject]
    protected ContainerInterface $container;

    #[Inject]
    protected ConfigInterface $config;

    #[Inject(value: \Hyperf\Redis\Redis::class)]
    protected \Hyperf\Redis\Redis|\Hyperf\Redis\RedisProxy $cache;  //默认redis的default连接库

    /**
     * 解析连接（多个redis情况下，用于确定使用哪个连接）
     *
     * @param array $connectionSelectData
     * @return void
     */
    public function connection(array $connectionSelectData)
    {
        //选择逻辑
        //$this->cache = $this->container->get(\Hyperf\Redis\RedisFactory::class)->get($this->connection);
    }
}
