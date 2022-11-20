<?php

declare(strict_types=1);

namespace App\Controller;

class Index extends AbstractController
{
    public function index()
    {
        //$url = $this->container->get(CommonLogic::class)->getBaseUrl() . 'view/admin/index/index.html';
        //return $this->response->redirect($url);
        return $this->response->redirect('/view/admin/index/index.html');
    }
}
