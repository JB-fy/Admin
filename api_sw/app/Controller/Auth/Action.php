<?php

declare(strict_types=1);

namespace App\Controller\Auth;

use App\Controller\AbstractController;
use App\Module\Db\Dao\Auth\Action as AuthAction;

class Action extends AbstractController
{
    /**
     * 列表
     *
     * @return void
     */
    public function list()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $data = $this->validate(__FUNCTION__, $sceneCode);
                $isAuth = $this->checkAuth(__FUNCTION__, $sceneCode, false);

                /**--------参数处理 开始--------**/
                if ($isAuth) {
                    $allowField = $this->getAllowField(AuthAction::class);
                } else {
                    $allowField = ['id', 'keyword', 'actionId', 'actionName'];
                }
                if (empty($data['field'])) {
                    $data['field'] = $allowField;
                } else {
                    $data['field'] = array_intersect($data['field'], $allowField);
                    if (empty($data['field'])) {
                        $data['field'] = $allowField;
                    }
                }
                /**--------参数处理 结束--------**/

                $this->service->listWithCount(...$data);
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
        $sceneCode = $this->scene->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $data = $this->validate(__FUNCTION__, $sceneCode);
                $this->checkAuth(__FUNCTION__, $sceneCode);

                /**--------参数处理 开始--------**/
                $allowField = $this->getAllowField(AuthAction::class);
                $allowField = array_merge($allowField, ['sceneIdArr']);
                if (empty($data['field'])) {
                    $data['field'] = $allowField;
                } else {
                    $data['field'] = array_intersect($data['field'], $allowField);
                    if (empty($data['field'])) {
                        $data['field'] = $allowField;
                    }
                }
                /**--------参数处理 结束--------**/

                $this->service->info(['id' => $data['id']], $data['field']);
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
        $sceneCode = $this->scene->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $data = $this->validate(__FUNCTION__, $sceneCode);
                $this->checkAuth(__FUNCTION__, $sceneCode);

                $this->service->create($data);
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
        $sceneCode = $this->scene->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $data = $this->validate(__FUNCTION__, $sceneCode);
                $this->checkAuth(__FUNCTION__, $sceneCode);

                $this->service->update($data, ['id' => $data['idArr']]);
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
        $sceneCode = $this->scene->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $data = $this->validate(__FUNCTION__, $sceneCode);
                $this->checkAuth(__FUNCTION__, $sceneCode);

                $this->service->delete(['id' => $data['idArr']]);
                break;
        }
    }
}
