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
    //云存储
    'cloudStorage' => function (\Psr\Container\ContainerInterface $container) {
        //$config = $this->config->get('app.aliyunOssConfig');
        $config = [
            'CM_ALIYUN_OSS_ACCESS_ID' => 'LTAI5tHx81H64BRJA971DPZF',
            'CM_ALIYUN_OSS_ACCESS_KEY' => 'nJyNpTtUuIgZqx21FF4G2zi0WHOn51',
            'CM_ALIYUN_OSS_HOST' => 'http://4724382110.oss-cn-hongkong.aliyuncs.com',
        ];
        return make(\App\Plugin\CloudStorage\AliyunOss::class, ['config' => $config]);
    },
    //平台后台场景信息
    'platformAdminSceneInfo' => function (\Psr\Container\ContainerInterface $container) {
        //$allScene = getDao(App\Module\Db\Dao\Auth\Scene::class)->getList();
        //$allScene = array_combine(array_column($allScene, 'sceneCode'), $allScene);
        $sceneInfo = getDao(\App\Module\Db\Dao\Auth\Scene::class)->where(['sceneCode' => 'platformAdmin'])->getInfo();
        $sceneInfo->sceneConfig = $sceneInfo->sceneConfig === null ? [] : json_decode($sceneInfo->sceneConfig, true);
        return $sceneInfo;
    },
    //平台管理员JWT签名
    'platformAdminJwt' => function (\Psr\Container\ContainerInterface $container) {
        /* $sceneInfo = $container->get(\App\Module\Logic\Auth\Scene::class)->getInfo('platformAdmin');
        $config = $sceneInfo->sceneConfig; */
        $config = getDao(\App\Module\Db\Dao\Auth\Scene::class)->where(['sceneCode' => 'platformAdmin'])->getBuilder()->value('sceneConfig');
        $config = json_decode($config, true);
        return make(\App\Plugin\Jwt::class, ['config' => $config]);
    }
];
