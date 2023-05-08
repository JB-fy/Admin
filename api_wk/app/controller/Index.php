<?php

declare(strict_types=1);

namespace app\controller;

use support\Request;

class Index extends AbstractController
{
    public function index(Request $request)
    {
        return redirect($request->url() . 'view/admin/index/index.html');
    }
}
