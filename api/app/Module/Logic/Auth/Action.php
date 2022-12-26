<?php

declare(strict_types=1);

namespace App\Module\Logic\Auth;

use App\Module\Db\Dao\Auth\ActionRelToScene;
use App\Module\Logic\AbstractLogic;

class Action extends AbstractLogic
{
    public function saveRelScene(array $sceneIdArr, int $id = 0)
    {
        $sceneIdArrOfOld = getDao(ActionRelToScene::class)->where(['actionId' => $id])->getBuilder()->pluck('sceneId')->toArray();
        /**----新增关联场景  开始----**/
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
        /**----新增关联场景  结束----**/

        /**----删除关联场景  开始----**/
        $deleteSceneIdArr = array_diff($sceneIdArrOfOld, $sceneIdArr);
        if (!empty($deleteSceneIdArr)) {
            getDao(ActionRelToScene::class)->where(['actionId' => $id, 'sceneId' => $deleteSceneIdArr])->delete();
        }
        /**----删除关联场景  结束----**/
    }
}
