<?php

declare(strict_types=1);

namespace App\Controller\Api;

use App\Controller\AbstractController;

class CloudStorage extends AbstractController
{
    /**
     * 获取签名
     *
     */
    public function getAccess()
    {
        $this->container->get('cloudStorage')->getAccess($this->request->post('uploadType'));
    }

    /**
     * 上传成功回调
     *
     */
    public function notify()
    {
        $this->container->get('cloudStorage')->notify();
    }
}
