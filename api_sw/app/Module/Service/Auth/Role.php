<?php

declare(strict_types=1);

namespace App\Module\Service\Auth;

use App\Module\Db\Dao\Auth\ActionRelToScene;
use App\Module\Db\Dao\Auth\Menu;
use App\Module\Db\Dao\Auth\RoleRelOfPlatformAdmin;
use App\Module\Db\Dao\Auth\RoleRelToAction;
use App\Module\Db\Dao\Auth\RoleRelToMenu;
use App\Module\Logic\Auth\Role as AuthRole;
use App\Module\Service\AbstractService;

class Role extends AbstractService
{
    /**
     * 创建
     *
     * @param array $data
     * @return void
     */
    public function create(array $data)
    {
        if (isset($data['menuIdArr']) && count($data['menuIdArr']) != getDao(Menu::class)->parseFilter(['id' => $data['menuIdArr'], 'sceneId' => $data['sceneId']])->getBuilder()->count()) {
            //$count = getDao(Menu::class)->parseFilter(['id' => $data['menuIdArr'], 'sceneId' => $data['sceneId'] ?? $oldInfo->sceneId])->info();
            throwFailJson(89999998);
        }
        if (isset($data['actionIdArr']) && count($data['actionIdArr']) != getDao(ActionRelToScene::class)->parseFilter(['actionId' => $data['actionIdArr'], 'sceneId' => $data['sceneId']])->getBuilder()->count()) {
            throwFailJson(89999998);
        }

        $id = $this->getDao()->parseInsert($data)->insert();
        if (empty($id)) {
            throwFailJson();
        }
        if (isset($data['menuIdArr'])) {
            $this->container->get(AuthRole::class)->saveRelMenu($data['menuIdArr'], $id);
        }
        if (isset($data['actionIdArr'])) {
            $this->container->get(AuthRole::class)->saveRelAction($data['actionIdArr'], $id);
        }
        throwSuccessJson();
    }

    /**
     * 更新
     *
     * @param array $data
     * @param array $filter
     * @return void
     */
    public function update(array $data, array $filter)
    {
        if (isset($data['menuIdArr']) || isset($data['actionIdArr'])) {
            $idArr = $this->getIdArr($filter);
            foreach ($idArr as $id) {
                $filterOne = ['roleId'=>$id];
                $oldInfo = $this->getDao()->parseFilter($filterOne)->info();
                $this->getDao()->parseFilter($filterOne)->parseUpdate($data)->update();    //有可能只改menuIdArr或actionIdArr
                if (isset($data['menuIdArr'])) {
                    if (count($data['menuIdArr']) != getDao(Menu::class)->parseFilter(['id' => $data['menuIdArr'], 'sceneId' => $data['sceneId'] ?? $oldInfo->sceneId])->getBuilder()->count()) {
                        throwFailJson(89999998);
                    }
                    $this->container->get(AuthRole::class)->saveRelMenu($data['menuIdArr'], $oldInfo->roleId);
                }
                if (isset($data['actionIdArr'])) {
                    if (count($data['actionIdArr']) != getDao(ActionRelToScene::class)->parseFilter(['actionId' => $data['actionIdArr'], 'sceneId' => $data['sceneId'] ?? $oldInfo->sceneId])->getBuilder()->count()) {
                        throwFailJson(89999998);
                    }
                    $this->container->get(AuthRole::class)->saveRelAction($data['actionIdArr'], $oldInfo->roleId);
                }
            }
        } else {
            $result = $this->getDao()->parseFilter($filter)->parseUpdate($data)->update();
            if (empty($result)) {
                throwFailJson();
            }
        }
        throwSuccessJson();
    }

    /**
     * 删除
     *
     * @param array $filter
     * @return void
     */
    public function delete(array $filter)
    {
        $idArr = $this->getIdArr($filter);
        $result = $this->getDao()->parseFilter($filter)->delete();
        if (empty($result)) {
            throwFailJson();
        }
        getDao(RoleRelToMenu::class)->parseFilter(['roleId' => $idArr])->delete();
        getDao(RoleRelToAction::class)->parseFilter(['roleId' => $idArr])->delete();
        getDao(RoleRelOfPlatformAdmin::class)->parseFilter(['roleId' => $idArr])->delete();
        throwSuccessJson();
    }
}
