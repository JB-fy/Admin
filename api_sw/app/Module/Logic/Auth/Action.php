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
        $sceneIdArrOfOld = getDao(ActionRelToScene::class)->where(['actionId' => $id])->getBuilder()->pluck('sceneId')->toArray();
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
            getDao(ActionRelToScene::class)->insert($insertList)->saveInsert();
        }
        /**----新增关联场景 结束----**/

        /**----删除关联场景 开始----**/
        $deleteSceneIdArr = array_diff($sceneIdArrOfOld, $sceneIdArr);
        if (!empty($deleteSceneIdArr)) {
            getDao(ActionRelToScene::class)->where(['actionId' => $id, 'sceneId' => $deleteSceneIdArr])->delete();
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
        var_dump(555);
        $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getCurrentInfo($sceneCode);
        var_dump(6666);
        $where = [
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
                //$where['selfAction']['loginId'] = $loginInfo->adminId;
                break;
            default:
                break;
        }
        if (empty(getDao(AuthAction::class)->where($where)->getBuilder()->count())) {
            if ($isThrow) {
                throwFailJson('39990002');
            } else {
                return false;
            }
        }
        return true;
    }
}
