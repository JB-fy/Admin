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
    'platformAdminSceneInfo' => function () {
        $allScene = getDao(App\Module\Db\Dao\Auth\Scene::class)->getList();
        $allScene = array_combine(array_column($allScene, 'sceneCode'), $allScene);
        return $allScene;
    },
    'platformAdminJwt' => function () {
        $config = getDao(App\Module\Db\Dao\Auth\Scene::class)->where(['sceneCode' => 'platformAdmin'])->getBuilder()->value('sceneConfig');
        $config = json_decode($config, true);
        return make(\App\Plugin\Jwt::class, ['config' => $config]);
    }
];
