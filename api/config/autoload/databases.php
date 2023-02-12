<?php

declare(strict_types=1);

/**
 * This file is part of Hyperf.
 *
 * @link     https://www.hyperf.io
 * @document https://hyperf.wiki
 * @contact  group@hyperf.io
 * @license  https://github.com/hyperf/hyperf/blob/master/LICENSE
 */
return [
    'default' => [
        'driver' => env('DB_DRIVER', 'mysql'),
        //'host' => env('DB_HOST', '127.0.0.1'),
        'write' => [
            'host' => explode(',', env('DB_HOST_WRITE', '127.0.0.1')),
        ],
        'read' => [
            'host' => explode(',', env('DB_HOST_READ', '127.0.0.1')),
        ],
        'sticky' => env('DB_STICKY', true),
        'database' => env('DB_DATABASE', 'hyperf'),
        'port' => env('DB_PORT', 3306),
        'username' => env('DB_USERNAME', 'root'),
        'password' => env('DB_PASSWORD', ''),
        'charset' => env('DB_CHARSET', 'utf8'),
        'collation' => env('DB_COLLATION', 'utf8_unicode_ci'),
        'prefix' => env('DB_PREFIX', ''),
        'timezone' => env('DB_TIMEZONE', '+8:00'),
        'pool' => [
            'min_connections' => (int) env('DB_MIN_CONNECTIONS', 1),
            'max_connections' => (int) env('DB_MAX_CONNECTIONS', 10),
            'connect_timeout' => (float) env('DB_CONNECT_TIMEOUT', 10.0),
            'wait_timeout' => (float) env('DB_WAIT_TIMEOUT', 3.0),
            'heartbeat' => (int) env('DB_HEARTBEAT', -1),
            'max_idle_time' => (float) env('DB_MAX_IDLE_TIME', 60.0),
        ],
        'commands' => [
            'gen:model' => [
                'path' => 'app/Module/Db/Model',
                'inheritance' => 'AbstractModel',
                'uses' => 'App\Module\Db\Model\AbstractModel',
                'force_casts' => true,
                'refresh_fillable' => true,
                'with_comments' => true,
                'table_mapping' => [
                    'auth_action:Auth\Action',
                    'auth_action_rel_to_scene:Auth\ActionRelToScene',
                    'auth_menu:Auth\Menu',
                    'auth_role:Auth\Role',
                    'auth_role_rel_of_platform_admin:Auth\RoleRelOfPlatformAdmin',
                    'auth_role_rel_to_action:Auth\RoleRelToAction',
                    'auth_role_rel_to_menu:Auth\RoleRelToMenu',
                    'auth_scene:Auth\Scene',

                    'log_request:Log\Request',

                    'platform_admin:Platform\Admin',
                    'platform_config:Platform\Config',
                    'platform_server:Platform\Server',
                ],
            ],
        ],
    ],
];
