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
    'platformAdminSceneInfo' => function (\Psr\Container\ContainerInterface $container) {
        //$allScene = getDao(App\Module\Db\Dao\Auth\Scene::class)->getList();
        //$allScene = array_combine(array_column($allScene, 'sceneCode'), $allScene);
        $sceneInfo = getDao(\App\Module\Db\Dao\Auth\Scene::class)->where(['sceneCode' => 'platformAdmin'])->getInfo();
        $sceneInfo->sceneConfig = $sceneInfo->sceneConfig === null ? [] : json_decode($sceneInfo->sceneConfig, true);
        return $sceneInfo;
    },
    'platformAdminJwt' => function (\Psr\Container\ContainerInterface $container) {
        /* $sceneInfo = $container->get(\App\Module\Logic\Auth\Scene::class)->getInfo('platformAdmin');
        $config = $sceneInfo->sceneConfig; */
        $config = getDao(\App\Module\Db\Dao\Auth\Scene::class)->where(['sceneCode' => 'platformAdmin'])->getBuilder()->value('sceneConfig');
        $config = json_decode($config, true);
        return make(\App\Plugin\Jwt::class, ['config' => $config]);
    }
];
