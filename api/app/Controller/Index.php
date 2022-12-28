<?php

declare(strict_types=1);

namespace App\Controller;

class Index extends AbstractController
{
    public function index()
    {
        return $this->response->redirect('/view/admin/index/index.html');
    }
}
