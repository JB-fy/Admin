<?php

declare(strict_types=1);

namespace App\Controller\Platform\Platform;

use App\Controller\AbstractController;
use App\Module\Db\Dao\Platform\Admin as PlatformAdmin;

class Admin extends AbstractController
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
        $this->checkAuth(__FUNCTION__, $sceneCode);

        /**--------参数处理 开始--------**/
        $allowField = $this->getAllowField(PlatformAdmin::class);
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
        $allowField = $this->getAllowField(PlatformAdmin::class);
        $allowField = array_merge($allowField, ['roleIdArr']);
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
        //不能修改平台超级管理员
        if (in_array(getConfig('app.superPlatformAdminId'), $data['idArr'])) {
            throwFailJson(30000000);
        }
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
        //不能删除平台超级管理员
        if (in_array(getConfig('app.superPlatformAdminId'), $data['idArr'])) {
            throwFailJson(30000001);
        }
        $this->checkAuth(__FUNCTION__, $sceneCode);

        $this->service->delete(['id' => $data['idArr']]);
    }
}
