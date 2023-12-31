<?php

declare(strict_types=1);

namespace App\Controller\Platform\My;

use App\Controller\AbstractController;

class Menu extends AbstractController
{
    /**
     * 菜单树
     *
     * @return void
     */
    public function tree()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getCurrentInfo($sceneCode);
        $sceneInfo = $this->container->get(\App\Module\Logic\Auth\Scene::class)->getCurrentSceneInfo();
        $filter = [
            'selfMenu' => [
                'sceneCode' => $sceneCode,
                'sceneId' => $sceneInfo->sceneId,
                'loginId' => $loginInfo->loginId
            ]
        ];
        $field = [
            'tree',
            'showMenu'
        ];
        $this->container->get(\App\Module\Service\Auth\Menu::class)->tree($filter, $field);
    }
}
