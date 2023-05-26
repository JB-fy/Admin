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
        if (isset($data['menuIdArr']) && count($data['menuIdArr']) != getDao(Menu::class)->filter(['id' => $data['menuIdArr'], 'sceneId' => $data['sceneId']])->getBuilder()->count()) {
            //$count = getDao(Menu::class)->filter(['id' => $data['menuIdArr'], 'sceneId' => $data['sceneId'] ?? $oldInfo->sceneId])->getInfo();
            throwFailJson(89999998);
        }
        if (isset($data['actionIdArr']) && count($data['actionIdArr']) != getDao(ActionRelToScene::class)->filter(['actionId' => $data['actionIdArr'], 'sceneId' => $data['sceneId']])->getBuilder()->count()) {
            throwFailJson(89999998);
        }

        $id = $this->getDao()->insert($data)->saveInsert();
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
     * @param array $where
     * @return void
     */
    public function update(array $data, array $where)
    {
        if (isset($data['menuIdArr']) || isset($data['actionIdArr'])) {
            $oldInfo = $this->getDao()->filter($where)->getInfo();
            if (isset($data['menuIdArr'])) {
                if (count($data['menuIdArr']) != getDao(Menu::class)->filter(['id' => $data['menuIdArr'], 'sceneId' => $data['sceneId'] ?? $oldInfo->sceneId])->getBuilder()->count()) {
                    throwFailJson(89999998);
                }
                $this->container->get(AuthRole::class)->saveRelMenu($data['menuIdArr'], $oldInfo->roleId);
                $this->getDao()->filter($where)->update($data)->saveUpdate();    //有可能只改menuIdArr
            }
            if (isset($data['actionIdArr'])) {
                if (count($data['actionIdArr']) != getDao(ActionRelToScene::class)->filter(['actionId' => $data['actionIdArr'], 'sceneId' => $data['sceneId'] ?? $oldInfo->sceneId])->getBuilder()->count()) {
                    throwFailJson(89999998);
                }
                $this->container->get(AuthRole::class)->saveRelAction($data['actionIdArr'], $oldInfo->roleId);
                $this->getDao()->filter($where)->update($data)->saveUpdate();    //有可能只改actionIdArr
            }
        } else {
            $result = $this->getDao()->filter($where)->update($data)->saveUpdate();
            if (empty($result)) {
                throwFailJson();
            }
        }
        throwSuccessJson();
    }

    /**
     * 删除
     *
     * @param array $where
     * @return void
     */
    public function delete(array $where)
    {
        $idArr = $this->getIdArr($where);
        $result = $this->getDao()->filter($where)->delete();
        if (empty($result)) {
            throwFailJson();
        }
        getDao(RoleRelToMenu::class)->filter(['roleId' => $idArr])->delete();
        getDao(RoleRelToAction::class)->filter(['roleId' => $idArr])->delete();
        getDao(RoleRelOfPlatformAdmin::class)->filter(['roleId' => $idArr])->delete();
        throwSuccessJson();
    }
}
