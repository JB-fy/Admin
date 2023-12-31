<?php

declare(strict_types=1);

namespace App\Module\Service\Platform;

use App\Module\Db\Dao\Auth\RoleRelOfPlatformAdmin;
use App\Module\Logic\Platform\Admin as PlatformAdmin;
use App\Module\Service\AbstractService;

class Admin extends AbstractService
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
        if (isset($data['roleIdArr'])) {
            $this->container->get(PlatformAdmin::class)->saveRelRole($data['roleIdArr'], $id);
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
        $idArr = $this->getIdArr($filter);
        if (empty($idArr)) {
            throwFailJson(99999999);
        }

        if (isset($data['passwordToCheck'])) {
            if (count($idArr) > 1) { //不支持批量修改
                throwFailJson(89999996, trans('code.89999996', ['errField' => 'passwordToCheck']));
            }
            $oldInfo = $this->getDao()->parseFilter($filter)->info();
            if (md5($data['passwordToCheck'] . $oldInfo->salt) != $oldInfo->password) {
                throwFailJson(39990003);
            }
        }

        if (isset($data['roleIdArr'])) {
            $this->getDao()->parseFilter($filter)->parseUpdate($data)->update();    //有可能只改roleIdArr
            foreach ($idArr as $id) {
                $this->container->get(PlatformAdmin::class)->saveRelRole($data['roleIdArr'], $id);
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
        getDao(RoleRelOfPlatformAdmin::class)->parseFilter(['adminId' => $idArr])->delete();
        throwSuccessJson();
    }
}
