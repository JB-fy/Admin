<?php

declare(strict_types=1);

namespace app\module\cache;

abstract class AbstractCache
{
    protected string $connection = '';  //默认连接

    /**
     * @var \support\Redis
     */
    protected $cache;

    final public function __construct(array $connectionSelectData = [])
    {
        $this->parseConnection($connectionSelectData);
    }

    /**
     * 解析连接（多个redis情况下，用于确定使用哪个连接）
     *
     * @param array $connectionSelectData
     * @return void
     */
    public function parseConnection(array $connectionSelectData)
    {
        //选择逻辑
        //$this->connection = ''; //设置当前使用的连接
        $this->cache = container(\support\Redis::class)->connection($this->connection);
    }
}
