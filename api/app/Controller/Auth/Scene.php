<?php

declare(strict_types=1);

namespace App\Controller\Auth;

use App\Controller\AbstractController;
use App\Module\Db\Dao\Auth\Scene as AuthScene;

class Scene extends AbstractController
{
    /**
     * 列表
     *
     * @return void
     */
    public function list()
    {
        $data = $this->validate(__FUNCTION__);
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $isAuth = $this->checkAuth(__FUNCTION__, $sceneCode, false);

                /**--------参数处理 开始--------**/
                if ($isAuth) {
                    $allowField = $this->getAllowField(AuthScene::class);
                } else {
                    $allowField = ['sceneId', 'sceneName', 'id'];
                }
                $data['field'] = empty($data['field']) ? $allowField : array_intersect($data['field'], $allowField);
                /**--------参数处理 结束--------**/

                $this->service->listWithCount(...$data);
                break;
            default:
                throwFailJson('39999999');
                break;
        }
    }

    /**
     * 详情
     *
     * @return void
     */
    public function info()
    {
        $data = $this->validate(__FUNCTION__);
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $this->checkAuth(__FUNCTION__, $sceneCode);

                /**--------参数处理 开始--------**/
                $allowField = $this->getAllowField(AuthScene::class);
                $data['field'] = empty($data['field']) ? $allowField : array_intersect($data['field'], $allowField);
                /**--------参数处理 结束--------**/

                $this->service->info(['id' => $data['id']], $data['field']);
                break;
            default:
                throwFailJson('39999999');
                break;
        }
    }

    /**
     * 创建
     *
     * @return void
     */
    public function create()
    {
        $data = $this->validate(__FUNCTION__);
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $this->checkAuth(__FUNCTION__, $sceneCode);

                $this->service->create($data);
                break;
            default:
                throwFailJson('39999999');
                break;
        }
    }

    /**
     * 更新
     *
     * @return void
     */
    public function update()
    {
        $data = $this->validate(__FUNCTION__);
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $this->checkAuth(__FUNCTION__, $sceneCode);

                $this->service->update($data, ['id' => $data['id']]);
                break;
            default:
                throwFailJson('39999999');
                break;
        }
    }

    /**
     * 删除
     *
     * @return void
     */
    public function delete()
    {
        $data = $this->validate(__FUNCTION__);
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $this->checkAuth(__FUNCTION__, $sceneCode);

                $this->service->delete(['id' => $data['idArr']]);
                break;
            default:
                throwFailJson('39999999');
                break;
        }
    }

    /**
     * 创建|更新时的参数处理
     *
     * @param array $data
     * @return array
     */
    /* protected function handleData(array $data): array
    {
        //isset($data['sceneConfig']) && empty($data['sceneConfig']) ? $data['sceneConfig'] = null : null;    //传值且为空时，数据库该字段设置为null。更建议放在Dao类中处理
        return $data;
    } */
}
