<?php

declare(strict_types=1);

namespace App\Crontab;

class LogHttp
{
    /**
     * 请求日志表分区
     *
     * @return void
     */
    public function partition()
    {
        /* try {
            dbTablePartition(\App\Module\Db\Dao\Log\Http::class, 7);
        } catch (\Throwable $th) {
            //出错时，做记录通知后台管理人员，让其联系技术人工处理。
            //也可以不捕获错误，启用app/Listener/CrontabListener监听器统一处理定时器报错问题
            var_dump($th->getMessage());
        } */
    }
}
