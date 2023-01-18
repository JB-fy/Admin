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
    //全部场景列表（即表auth_scene数据）
    'allScene' => function (ContainerInterface $container) {
        $allScene = getDao(\App\Module\Db\Dao\Auth\Scene::class)->getList();
        $allScene = array_combine(array_column($allScene, 'sceneCode'), $allScene);
        foreach ($allScene as &$v) {
            $v->sceneConfig = $v->sceneConfig === null ? [] : json_decode($v->sceneConfig, true);
        }
        return $allScene;
    },
    //平台管理员JWT插件
    'platformAdminJwt' => function (ContainerInterface $container) {
        /* $config = getDao(\App\Module\Db\Dao\Auth\Scene::class)->where(['sceneCode' => 'platformAdmin'])->getBuilder()->value('sceneConfig');
        $config = json_decode($config, true); */
        $sceneInfo = $container->get(\App\Module\Logic\Auth\Scene::class)->getInfo('platformAdmin');
        return make(\App\Plugin\Jwt::class, ['config' => $sceneInfo->sceneConfig]);
    },
    //全部平台配置（即表platform_config数据）
    'allPlatformConfig' => function (ContainerInterface $container) {
        $allPlatformConfig = getDao(\App\Module\Db\Dao\Platform\Config::class)->getBuilder()->pluck('configValue', 'configKey');
        return $allPlatformConfig;
    },
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
];
