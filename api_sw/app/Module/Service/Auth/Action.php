<?php

declare(strict_types=1);

namespace App\Module\Service\Auth;

use App\Module\Db\Dao\Auth\ActionRelToScene;
use App\Module\Db\Dao\Auth\RoleRelToAction;
use App\Module\Logic\Auth\Action as AuthAction;
use App\Module\Service\AbstractService;

class Action extends AbstractService
{
    /**
     * 创建
     *
     * @param array $data
     * @return void
     */
    public function create(array $data)
    {
        $id = $this->getDao()->parseInsert($data)->insert();
        if (empty($id)) {
            throwFailJson();
        }
        if (isset($data['sceneIdArr'])) {
            $this->container->get(AuthAction::class)->saveRelScene($data['sceneIdArr'], $id);
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
        if (isset($data['sceneIdArr'])) {
            $idArr = $this->getIdArr($filter);
            $this->getDao()->parseFilter($filter)->parseUpdate($data)->update();    //有可能只改sceneIdArr
            foreach ($idArr as $id) {
                $this->container->get(AuthAction::class)->saveRelScene($data['sceneIdArr'], $id);
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
        getDao(ActionRelToScene::class)->parseFilter(['actionId' => $idArr])->delete();
        getDao(RoleRelToAction::class)->parseFilter(['actionId' => $idArr])->delete();
        throwSuccessJson();
    }
}
