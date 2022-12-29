<?php

declare(strict_types=1);

namespace App\Module\Logic\Platform;

use App\Module\Db\Dao\Auth\RoleRelOfPlatformAdmin;
use App\Module\Logic\AbstractLogic;

class Admin extends AbstractLogic
{
    /**
     * 保存关联角色
     *
     * @param array $roleIdArr
     * @param integer $id
     * @return void
     */
    public function saveRelRole(array $roleIdArr, int $id = 0)
    {
        $roleIdArrOfOld = getDao(RoleRelOfPlatformAdmin::class)->where(['adminId' => $id])->getBuilder()->pluck('roleId')->toArray();
        /**----新增关联角色 开始----**/
        $insertRoleIdArr = array_diff($roleIdArr, $roleIdArrOfOld);
        if (!empty($insertRoleIdArr)) {
            $insertList = [];
            foreach ($insertRoleIdArr as $v) {
                $insertList[] = [
                    'adminId' => $id,
                    'roleId' => $v
                ];
            }
            getDao(RoleRelOfPlatformAdmin::class)->insert($insertList)->saveInsert();
        }
        /**----新增关联角色 结束----**/

        /**----删除关联角色 开始----**/
        $deleteRoleIdArr = array_diff($roleIdArrOfOld, $roleIdArr);
        if (!empty($deleteRoleIdArr)) {
            getDao(RoleRelOfPlatformAdmin::class)->where(['adminId' => $id, 'roleId' => $deleteRoleIdArr])->delete();
        }
        /**----删除关联角色 结束----**/
    }
}
