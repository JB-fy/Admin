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
        $id = $this->getDao()->insert($data)->saveInsert();
        if (empty($id)) {
            throwFailJson('999999');
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
        $this->getDao()->where($where)->update($data)->saveUpdate();
        //上面可能结果为0，而只修改sceneIdArr
        if (isset($data['sceneIdArr'])) {
            $id = isset($where['id']) ? $where['id'] : $this->getDao()->where($where)->getBuilder()->value('actionId');
            $this->container->get(AuthAction::class)->saveRelScene($data['sceneIdArr'], $id);
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
        $id = isset($where['id']) ? $where['id'] : $this->getDao()->where($where)->getBuilder()->pluck('actionId')->toArray();
        $result = $this->getDao()->where($where)->delete();
        if (empty($result)) {
            throwFailJson('999999');
        }
        getDao(ActionRelToScene::class)->where(['actionId' => $id])->delete();
        throwSuccessJson();
    }
}
