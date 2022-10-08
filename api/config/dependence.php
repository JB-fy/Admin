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

use Psr\Container\ContainerInterface;

return [
    //'systemAdminJwt' => container(app\plugin\Jwt::class, true, ['config' => config('custom.auth.systemAdmin')])
    'systemAdminJwt' => function (ContainerInterface $container) {
        return $container->make(app\plugin\Jwt::class, ['config' => config('custom.auth.systemAdmin')]);
    },
];
