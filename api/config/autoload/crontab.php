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
    'enable' => env('ENABLE_CRONTAB', false),//是否开启定时任务
    'crontab' => [//通过配置文件定义的定时任务
        /**--------Callback类型定时任务（默认） 开始--------**/
        (new Crontab())
            ->setName('LogRequestPartition')
            ->setCallback([App\Crontab\LogRequest::class, 'partition'])
            //->setRule('0 */15 3 * * 0')//星期天的凌晨3点每15分钟执行一次（多次执行防止新增分区失败）
            ->setRule('*/5 * * * * *')
            //->setSingleton(true)
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
