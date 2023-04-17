<?php

declare(strict_types=1);

namespace App\Module\Logic\Auth;

use App\Module\Db\Dao\Auth\RoleRelToAction;
use App\Module\Db\Dao\Auth\RoleRelToMenu;
use App\Module\Logic\AbstractLogic;

class Role extends AbstractLogic
{
    /**
     * 保存关联菜单
     *
     * @param array $menuIdArr
     * @param integer $id
     * @return void
     */
    public function saveRelMenu(array $menuIdArr, int $id = 0)
    {
        $menuIdArrOfOld = getDao(RoleRelToMenu::class)->where(['roleId' => $id])->getBuilder()->pluck('menuId')->toArray();
        /**----新增关联菜单 开始----**/
        $insertMenuIdArr = array_diff($menuIdArr, $menuIdArrOfOld);
        if (!empty($insertMenuIdArr)) {
            $insertList = [];
            foreach ($insertMenuIdArr as $v) {
                $insertList[] = [
                    'roleId' => $id,
                    'menuId' => $v
                ];
            }
            getDao(RoleRelToMenu::class)->insert($insertList)->saveInsert();
        }
        /**----新增关联菜单 结束----**/

        /**----删除关联菜单 开始----**/
        $deleteMenuIdArr = array_diff($menuIdArrOfOld, $menuIdArr);
        if (!empty($deleteMenuIdArr)) {
            getDao(RoleRelToMenu::class)->where(['roleId' => $id, 'menuId' => $deleteMenuIdArr])->delete();
        }
        /**----删除关联菜单 结束----**/
    }

    /**
     * 保存关联操作
     *
     * @param array $actionIdArr
     * @param integer $id
     * @return void
     */
    public function saveRelAction(array $actionIdArr, int $id = 0)
    {
        $actionIdArrOfOld = getDao(RoleRelToAction::class)->where(['roleId' => $id])->getBuilder()->pluck('actionId')->toArray();
        /**----新增关联操作 开始----**/
        $insertActionIdArr = array_diff($actionIdArr, $actionIdArrOfOld);
        if (!empty($insertActionIdArr)) {
            $insertList = [];
            foreach ($insertActionIdArr as $v) {
                $insertList[] = [
                    'roleId' => $id,
                    'actionId' => $v
                ];
            }
            getDao(RoleRelToAction::class)->insert($insertList)->saveInsert();
        }
        /**----新增关联操作 结束----**/

        /**----删除关联操作 开始----**/
        $deleteActionIdArr = array_diff($actionIdArrOfOld, $actionIdArr);
        if (!empty($deleteActionIdArr)) {
            getDao(RoleRelToAction::class)->where(['roleId' => $id, 'actionId' => $deleteActionIdArr])->delete();
        }
        /**----删除关联操作 结束----**/
    }
}
