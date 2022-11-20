<?php

declare(strict_types=1);

namespace App\Controller;

class Test extends AbstractController
{
    public function index()
    {
        //sleep(10);
        //throwSuccessJson([]);
        //throwRaw('哈哈阿');
        throwFailJson('000999');
    }
}
