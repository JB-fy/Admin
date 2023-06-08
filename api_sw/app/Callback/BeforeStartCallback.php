<?php

declare(strict_types=1);

namespace App\Callback;

use Hyperf\Contract\ConfigInterface;
use Hyperf\Di\Annotation\Inject;
use Psr\Container\ContainerInterface;

//必须先在config/autoload/server.php内callbacks字段增加Event::ON_BEFORE_START => [\App\Callback\BeforeStartCallback::class, 'onBeforeStart']
class BeforeStartCallback
{
    #[Inject]
    protected ContainerInterface $container;

    #[Inject]
    protected ConfigInterface $config;

    public function onBeforeStart()
    {
        /**--------设置当前服务器IP并记录 开始--------**/
        $serverNetworkIp = getServerNetworkIp();
        $serverLocalIp = getServerLocalIp();
        $this->config->set('server.networkIp', $serverNetworkIp);   //设置服务器外网ip
        $this->config->set('server.localIp', $serverLocalIp);   //设置服务器内网ip
        try {
            getDao(\App\Module\Db\Dao\Platform\Server::class)->getBuilder()->updateOrInsert(['networkIp' => $serverNetworkIp], ['localIp' => $serverLocalIp]);
        } catch (\Throwable $th) {
        }
        /**--------设置当前服务器IP并记录 结束--------**/

        /**--------将数据库内的配置设置到config中（方便使用） 开始--------**/
        //场景列表（即表auth_scene数据）
        $allScene = getDao(\App\Module\Db\Dao\Auth\Scene::class)->list();
        $allScene = array_combine(array_column($allScene, 'sceneCode'), $allScene);
        foreach ($allScene as &$v) {
            $v->sceneConfig = $v->sceneConfig === null ? [] : json_decode($v->sceneConfig, true);
        }
        $this->config->set('inDb.authScene', $allScene);

        //平台配置（即表platform_config数据）
        $allPlatformConfig = getDao(\App\Module\Db\Dao\Platform\Config::class)->getBuilder()->pluck('configValue', 'configKey');
        $this->config->set('inDb.platformConfig', $allPlatformConfig);
        /**--------将数据库内的配置设置到config中（方便使用） 结束--------**/

        /**--------数据库表分区 开始--------**/
        $this->container->get(\App\Crontab\LogRequest::class)->partition(); //请求日志
        /**--------数据库表分区 结束--------**/
    }
}
