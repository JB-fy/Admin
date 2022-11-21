<?php

declare(strict_types=1);

namespace App\Controller;

use App\Module\Db\Dao\Auth\Scene;

class Test extends AbstractController
{
    public function index()
    {
        $a = make(Scene::class)->getBuilder()->get()->toArray();
        //var_dump($a);
        //sleep(10);
        //throwSuccessJson([]);
        //throwRaw('哈哈阿');
        throwFailJson('000999');
    }
}
