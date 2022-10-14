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

use app\module\db\table\auth\AuthScene;
use Psr\Container\ContainerInterface;

return [
    'systemAdminJwt' => function (ContainerInterface $container) {
        $config = container(AuthScene::class, true)->where(['sceneCode' => 'systemAdmin'])->getBuilder()->value('sceneConfig');
        $config = json_decode($config, true);
        return $container->make(app\plugin\Jwt::class, ['config' => $config]);
    },
];
