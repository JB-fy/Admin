<?php

declare(strict_types=1);

namespace App\Controller\Auth;

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
        switch ($sceneCode) {
            case 'platformAdmin':
                $data = $this->validate(__FUNCTION__, $sceneCode);
                $isAuth = $this->checkAuth(__FUNCTION__, $sceneCode, false);

                /**--------参数处理 开始--------**/
                if ($isAuth) {
                    $allowField = $this->getAllowField(AuthMenu::class);
                    $allowField = array_merge($allowField, ['sceneName', 'pMenuName']);
                } else {
                    $allowField = ['menuId', 'menuName', 'id'];
                }
                $data['field'] = empty($data['field']) ? $allowField : array_intersect($data['field'], $allowField);
                /**--------参数处理 结束--------**/

                $this->service->listWithCount(...$data);
                break;
            default:
                throwFailJson(39999999);
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
                $allowField = $this->getAllowField(AuthMenu::class);
                $data['field'] = empty($data['field']) ? $allowField : array_intersect($data['field'], $allowField);
                /**--------参数处理 结束--------**/

                $this->service->info(['id' => $data['id']], $data['field']);
                break;
            default:
                throwFailJson(39999999);
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
            default:
                throwFailJson(39999999);
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
            default:
                throwFailJson(39999999);
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
            default:
                throwFailJson(39999999);
                break;
        }
    }

    /**
     * 获取菜单树
     *
     * @return void
     */
    public function tree()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $data = $this->validate(__FUNCTION__, $sceneCode);
                $isAuth = $this->checkAuth(__FUNCTION__, $sceneCode, false);

                /**--------参数处理 开始--------**/
                if ($isAuth) {
                    $allowField = $this->getAllowField(AuthMenu::class);
                    $allowField = array_merge($allowField, ['sceneName', 'pMenuName']);
                } else {
                    $allowField = ['menuId', 'menuName', 'id'];
                }
                $data['field'] = empty($data['field']) ? $allowField : array_intersect($data['field'], $allowField);

                $data['where'] = array_merge($data['where'], ['isStop' => 0]);  //补充条件
                $data['field'] = array_merge($data['field'], ['menuTree']); //补充字段（树状菜单所需）
                /**--------参数处理 结束--------**/

                $this->service->tree(...$data);
                break;
            default:
                throwFailJson(39999999);
                break;
        }
    }
}
