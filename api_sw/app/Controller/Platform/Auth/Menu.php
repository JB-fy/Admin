<?php

declare(strict_types=1);

namespace App\Controller\Platform\Auth;

use App\Controller\AbstractController;
use App\Module\Db\Dao\Auth\Menu as AuthMenu;

class Menu extends AbstractController
{
    /**
     * 列表
     *
     * @return void
     */
    public function list()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $data = $this->validate(__FUNCTION__, $sceneCode);
        $isAuth = $this->checkAuth(__FUNCTION__, $sceneCode, false);

        /**--------参数处理 开始--------**/
        if ($isAuth) {
            $allowField = $this->getAllowField(AuthMenu::class);
            $allowField = array_merge($allowField, ['sceneName', 'pMenuName']);
        } else {
            $allowField = ['id', 'label', 'menuId', 'menuName'];
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
    }

    /**
     * 详情
     *
     * @return void
     */
    public function info()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $data = $this->validate(__FUNCTION__, $sceneCode);
        $this->checkAuth(__FUNCTION__, $sceneCode);

        /**--------参数处理 开始--------**/
        $allowField = $this->getAllowField(AuthMenu::class);
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
    }

    /**
     * 创建
     *
     * @return void
     */
    public function create()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $data = $this->validate(__FUNCTION__, $sceneCode);
        $this->checkAuth(__FUNCTION__, $sceneCode);

        $this->service->create($data);
    }

    /**
     * 更新
     *
     * @return void
     */
    public function update()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $data = $this->validate(__FUNCTION__, $sceneCode);
        $this->checkAuth(__FUNCTION__, $sceneCode);

        $this->service->update($data, ['id' => $data['idArr']]);
    }

    /**
     * 删除
     *
     * @return void
     */
    public function delete()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $data = $this->validate(__FUNCTION__, $sceneCode);
        $this->checkAuth(__FUNCTION__, $sceneCode);

        $this->service->delete(['id' => $data['idArr']]);
    }

    /**
     * 菜单树
     *
     * @return void
     */
    public function tree()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        $data = $this->validate(__FUNCTION__, $sceneCode);
        $isAuth = $this->checkAuth(__FUNCTION__, $sceneCode, false);

        /**--------参数处理 开始--------**/
        if ($isAuth) {
            $allowField = $this->getAllowField(AuthMenu::class);
            $allowField = array_merge($allowField, ['sceneName', 'pMenuName']);
        } else {
            $allowField = ['menuId', 'menuName', 'id'];
        }
        if (empty($data['field'])) {
            $data['field'] = $allowField;
        } else {
            $data['field'] = array_intersect($data['field'], $allowField);
            if (empty($data['field'])) {
                $data['field'] = $allowField;
            }
        }

        $data['filter'] = array_merge($data['filter'], ['isStop' => 0]);
        $data['field'] = array_merge($data['field'], ['tree']);
        /**--------参数处理 结束--------**/

        $this->service->tree(...$data);
    }
}
