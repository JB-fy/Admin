<?php

declare(strict_types=1);

namespace App\Module\Service\Auth;

use App\Module\Db\Dao\Auth\ActionRelToScene;
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
     * @param array $where
     * @return void
     */
    public function update(array $data, array $where)
    {
        if (isset($data['sceneIdArr'])) {
            $idArr = $this->getIdArr($where);
            foreach ($idArr as $id) {
                $this->container->get(AuthAction::class)->saveRelScene($data['sceneIdArr'], $id);
            }
            $this->getDao()->parseFilter($where)->parseUpdate($data)->update();    //有可能只改sceneIdArr
        } else {
            $result = $this->getDao()->parseFilter($where)->parseUpdate($data)->update();
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
        $result = $this->getDao()->parseFilter($where)->delete();
        if (empty($result)) {
            throwFailJson();
        }
        getDao(ActionRelToScene::class)->parseFilter(['actionId' => $idArr])->delete();
        throwSuccessJson();
    }
}
