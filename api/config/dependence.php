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

use app\module\db\dao\auth\AuthScene;
use Psr\Container\ContainerInterface;

return [
    /**
     * 这里定义的依赖注入，使用时必须特别注意。以下面为例说明
     *  container('systemAdminJwt')   第一次使用时，读取一次配置，随后缓存在容器内
     *  container('systemAdminJwt', true) 每次使用都会读取配置
     *  由于配置不会经常改动，不建议每次使用都读取配置
     *  但使用container('systemAdminJwt')，必须在更改配置后，重启webman服务
     */
    'systemAdminJwt' => function (ContainerInterface $container) {
        $config = container(AuthScene::class, true)->where(['sceneCode' => 'systemAdmin'])->getBuilder()->value('sceneConfig');
        $config = json_decode($config, true);
        return $container->make(app\plugin\Jwt::class, ['config' => $config]);
    },
];
