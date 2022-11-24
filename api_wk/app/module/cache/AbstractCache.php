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

    public function __construct(array $connectionSelectData = [])
    {
        $this->connection($connectionSelectData);
    }

    /**
     * 解析连接（多个redis情况下，用于确定使用哪个连接）
     *
     * @param array $connectionSelectData
     * @return void
     */
    public function connection(array $connectionSelectData)
    {
        //选择逻辑
        //$this->connection = ''; //设置当前使用的连接
        $this->cache = container(\support\Redis::class)->connection($this->connection);
    }
}
