<?php

declare(strict_types=1);

use Hyperf\Crontab\Crontab;

/**
 *  定时器说明
 *  * * * * * *
 *  | | | | | |--星期（0-6，星期天为0）
 *  | | | | |----月（1-12）
 *  | | | |------日（1-31）
 *  | | |--------时（0-23）
 *  | |----------分（0-59）
 *  |------------秒（0-59）
 */
return [
    'enable' => env('CRONTAB_ENABLE', true), //是否开启定时任务
    //通过配置文件定义的定时任务
    'crontab' => [
        /**--------Callback类型定时任务（默认） 开始--------**/
        (new Crontab())
            ->setName('LogHttpPartition')
            ->setCallback([App\Crontab\LogHttp::class, 'partition'])
            ->setRule('0 0 3 * * 1') //星期一的凌晨3点执行（方便人工检查是否成功，防止失败隔天该表无法正常使用）
            //->setRule('*/5 * * * * *')
            //->setSingleton(true)  //解决单机任务并发问题，同时只会运行1个。但集群时无用
            /* ->setOnOneServer(true)  //多实例（server）部署项目时，则只有一个实例会被触发。原理：通过在redis内做缓存实现互斥锁
            ->setMutexPool('default') //当setOnOneServer(true)时，使用哪个redis连接池
            ->setMutexExpires(3600)   //当setOnOneServer(true)时，redis内key的缓存时间。同时如果定时任务执行完，但解除互斥锁失败时，互斥锁也会在这个时间之后自动解除。 */
            //->setExecuteTime(\Carbon\Carbon::now())
            //->setEnable(false)
            ->setMemo('请求日志表每周新增分区'),
        /**--------Callback类型定时任务（默认）结束--------**/

        /**--------Command类型定时任务 开始--------**/
        /*(new Crontab())
            ->setType('command')
            ->setName('Bar')
            ->setRule('* * * * *')
            ->setCallback([
                'command' => 'swiftmailer:spool:send',
                // (optional) arguments
                'fooArgument' => 'barValue',
                // (optional) options
                '--message-limit' => 1,
            ]),*/
        /**--------Command类型定时任务 结束--------**/
    ],
];
