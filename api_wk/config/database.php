<?php

/**
 * This file is part of webman.
 *
 * Licensed under The MIT License
 * For full copyright and license information, please see the MIT-LICENSE.txt
 * Redistributions of files must retain the above copyright notice.
 *
 * @author    walkor<walkor@workerman.net>
 * @copyright walkor<walkor@workerman.net>
 * @link      http://www.workerman.net/
 * @license   http://www.opensource.org/licenses/mit-license.php MIT License
 */

return [
    // 默认数据库
    'default' => 'mysql',

    // 各种数据库配置
    'connections' => [
        'mysql' => [
            'driver' => getenv('DB_DRIVER'),
            //'host'        => getenv('DB_HOST'),
            'write' => [
                'host' => explode(',', getenv('DB_HOST_WRITE')),
            ],
            'read' => [
                'host' => explode(',', getenv('DB_HOST_READ')),
            ],
            'strict'      => getenv('DB_STICKY') === 'true' ? true : false,
            'port' => (int) getenv('DB_PORT'),
            'database' => getenv('DB_DATABASE'),
            'username' => getenv('DB_USERNAME'),
            'password' => getenv('DB_PASSWORD'),
            'charset'     => getenv('DB_CHARSET'),
            'collation' => getenv('DB_COLLATION'),
            'prefix'      => getenv('DB_PREFIX'),
            'timezone' => getenv('DB_TIMEZONE'),
            'unix_socket' => '',
            'engine'      => null,
            'modes' => ['STRICT_TRANS_TABLES', 'ERROR_FOR_DIVISION_BY_ZERO', 'NO_ENGINE_SUBSTITUTION', 'NO_ZERO_IN_DATE', 'NO_ZERO_DATE'],
        ],
    ],
];
