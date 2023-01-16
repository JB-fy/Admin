<?php

declare(strict_types=1);

namespace App\Crontab;

use App\Module\Db\Dao\Log\Request;
use Hyperf\DbConnection\Db;
/* use Hyperf\Di\Annotation\Inject;
use Psr\Container\ContainerInterface;

class LogRequest
{
    #[Inject]
    protected ContainerInterface $container; */

class LogRequest
{
    /**
     * 请求日志表分区
     *
     * @return void
     */
    public function partition()
    {
        try {
            dbTablePartition(\App\Module\Db\Dao\Log\Request::class, 24 * 60 * 60, 9);
        } catch (\Throwable $th) {
            //出错时，做记录通知后台管理人员，让其联系技术人工处理
            var_dump($th->getMessage());
        }
    }
}
