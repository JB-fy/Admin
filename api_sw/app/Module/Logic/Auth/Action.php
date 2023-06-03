<?php

declare(strict_types=1);

namespace App\Module\Logic\Auth;

use App\Module\Db\Dao\Auth\Action as AuthAction;
use App\Module\Db\Dao\Auth\ActionRelToScene;
use App\Module\Logic\AbstractLogic;

class Action extends AbstractLogic
{
    /**
     * 保存关联场景
     *
     * @param array $sceneIdArr
     * @param integer $id
     * @return void
     */
    public function saveRelScene(array $sceneIdArr, int $id = 0)
    {
        $sceneIdArrOfOld = getDao(ActionRelToScene::class)->parseFilter(['actionId' => $id])->getBuilder()->pluck('sceneId')->toArray();
        /**----新增关联场景 开始----**/
        $insertSceneIdArr = array_diff($sceneIdArr, $sceneIdArrOfOld);
        if (!empty($insertSceneIdArr)) {
            $insertList = [];
            foreach ($insertSceneIdArr as $v) {
                $insertList[] = [
                    'actionId' => $id,
                    'sceneId' => $v
                ];
            }
            getDao(ActionRelToScene::class)->parseInsert($insertList)->insert();
        }
        /**----新增关联场景 结束----**/

        /**----删除关联场景 开始----**/
        $deleteSceneIdArr = array_diff($sceneIdArrOfOld, $sceneIdArr);
        if (!empty($deleteSceneIdArr)) {
            getDao(ActionRelToScene::class)->parseFilter(['actionId' => $id, 'sceneId' => $deleteSceneIdArr])->delete();
        }
        /**----删除关联场景 结束----**/
    }

    /**
     * 判断操作权限
     *
     * @param string $actionCode
     * @param string $sceneCode
     * @param boolean $isThrow
     * @return boolean
     */
    public function checkAuth(string $actionCode, string $sceneCode, bool $isThrow = true): bool
    {
        $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getCurrentInfo($sceneCode);
        $filter = [
            'actionCode' => $actionCode,
            'selfAction' => [
                'sceneCode' => $sceneCode,
                'loginId' => $loginInfo->adminId
            ],
        ];
        switch ($sceneCode) {
            case 'platformAdmin':
                if ($loginInfo->adminId == getConfig('app.superPlatformAdminId')) { //平台超级管理员，无权限限制
                    return true;
                }
                //$filter['selfAction']['loginId'] = $loginInfo->adminId;
                break;
            default:
                break;
        }
        if (empty(getDao(AuthAction::class)->parseFilter($filter)->getBuilder()->count())) {
            if ($isThrow) {
                throwFailJson(39990002);
            } else {
                return false;
            }
        }
        return true;
    }
}
