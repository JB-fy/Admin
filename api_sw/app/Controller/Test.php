<?php

declare(strict_types=1);

namespace App\Controller;

class Test extends AbstractController
{
    public function index()
    {
        //$data = $this->request->all();
        //$crontab = $this->container->get(\Hyperf\Contract\ConfigInterface::class)->get('crontab.crontab.LogHttpPartition');
        //$a = getDao(\App\Module\Db\Dao\Auth\Scene::class)->getBuilder()->get()->toArray();
        //$a = $this->container->get(\Hyperf\Redis\RedisFactory::class)->get('default')->set('aaaa', 'asda', 10);
        //$a = $this->container->get(\App\Module\Cache\Login::class);
        //$a = $this->container->get(\Hyperf\Contract\ConfigInterface::class)->get('app.cache.saltFormat');
        //$a = $this->container->get(\App\Module\Validation\Login::class)->make($this->request->all(), 'salt')->validate();
        //var_dump($a);
        //sleep(10);
        //\Swoole\Coroutine::sleep(10);
        //throwSuccessJson([]);
        //throwRaw('哈哈阿');
        throwFailJson(89999999);

        /* $cacheLogin = make(\App\Module\Cache\Login::class);
        $cacheLogin->setSaltKey('admin', 'aaaa');
        $salt = randStr(8);
        $cacheLogin->setSalt($salt); */
    }
}
