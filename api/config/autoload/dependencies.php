<?php

declare(strict_types=1);

use Psr\Container\ContainerInterface;

/**
 * This file is part of Hyperf.
 *
 * @link     https://www.hyperf.io
 * @document https://hyperf.wiki
 * @contact  group@hyperf.io
 * @license  https://github.com/hyperf/hyperf/blob/master/LICENSE
 */
return [
    //上传组件
    'upload' => function (ContainerInterface $container) {
        /* $config = [
            'accessId' => 'LTAI5tHx81H64BRJA971DPZF',
            'accessKey' => 'nJyNpTtUuIgZqx21FF4G2zi0WHOn51',
            'host' => 'http://4724382110.oss-cn-hongkong.aliyuncs.com'
        ]; */
        $config = [
            'accessId' => 'LTAI5tSjYikt3bX33riHezmk',
            'accessKey' => 'k4uRZU6flv73yz1j4LJu9VY5eNlHas',
            'host' => 'https://gamemt.oss-cn-hangzhou.aliyuncs.com'
        ];
        return make(\App\Plugin\Upload\AliyunOss::class, ['config' => $config]);
    },
    //平台后台场景信息
    'platformAdminSceneInfo' => function (ContainerInterface $container) {
        //$allScene = getDao(App\Module\Db\Dao\Auth\Scene::class)->getList();
        //$allScene = array_combine(array_column($allScene, 'sceneCode'), $allScene);
        $sceneInfo = getDao(\App\Module\Db\Dao\Auth\Scene::class)->where(['sceneCode' => 'platformAdmin'])->getInfo();
        $sceneInfo->sceneConfig = $sceneInfo->sceneConfig === null ? [] : json_decode($sceneInfo->sceneConfig, true);
        return $sceneInfo;
    },
    //平台管理员JWT插件
    'platformAdminJwt' => function (ContainerInterface $container) {
        /* $config = getDao(\App\Module\Db\Dao\Auth\Scene::class)->where(['sceneCode' => 'platformAdmin'])->getBuilder()->value('sceneConfig');
        $config = json_decode($config, true); */
        $sceneInfo = getContainer()->get(\App\Module\Logic\Auth\Scene::class)->getInfo('platformAdmin');
        return make(\App\Plugin\Jwt::class, ['config' => $sceneInfo->sceneConfig]);
    }
];
