<?php

declare(strict_types=1);

namespace App\Controller\Auth;

use App\Controller\AbstractController;
use App\Module\Db\Dao\Auth\Role as AuthRole;

class Role extends AbstractController
{
    /**
     * 列表
     *
     * @return void
     */
    public function list()
    {
        $data = $this->validate(__FUNCTION__); //参数验证并处理
        $sceneCode = getRequestScene();
        switch ($sceneCode) {
            case 'platformAdmin':
                $isAuth = $this->container->get(\App\Module\Logic\Auth\Role::class)->checkAuth('authRoleLook', $sceneCode, false);  //验证权限

                /**--------参数过滤 结束--------**/
                /* if ($isAuth) {
                    $allowField = getDao(AuthRole::class)->getAllColumn();
                    $allowField = array_merge($allowField, ['sceneName', 'pActionName']);
                } else {
                    //无查看权限时只能查看一些基本的字段
                    $allowField = ['menuId', 'menuName', 'menu'];
                } */

                $allowField = getDao(AuthRole::class)->getAllColumn();
                $allowField = array_merge($allowField, ['id', 'sceneName']);
                $data['field'] = empty($data['field']) ? $allowField : array_intersect($data['field'], $allowField);    //过滤不可查看字段
                /**--------参数过滤 结束--------**/

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
        $data = $this->validate(__FUNCTION__); //参数验证并处理
        $sceneCode = getRequestScene();
        switch ($sceneCode) {
            case 'platformAdmin':
                $this->container->get(\App\Module\Logic\Auth\Role::class)->checkAuth('authRoleLook', $sceneCode);  //验证权限

                $allowField = getDao(AuthRole::class)->getAllColumn();
                $allowField = array_merge($allowField, ['id', 'sceneName', 'menuIdArr', 'actionIdArr']);
                $data['field'] = empty($data['field']) ? $allowField : array_intersect($data['field'], $allowField);    //过滤不可查看字段
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
        $data = $this->validate(__FUNCTION__); //参数验证并处理
        $sceneCode = getRequestScene();
        switch ($sceneCode) {
            case 'platformAdmin':
                $this->container->get(\App\Module\Logic\Auth\Role::class)->checkAuth('authRoleCreate', $sceneCode);  //验证权限

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
        $data = $this->validate(__FUNCTION__); //参数验证并处理
        $sceneCode = getRequestScene();
        switch ($sceneCode) {
            case 'platformAdmin':
                $this->container->get(\App\Module\Logic\Auth\Role::class)->checkAuth('authRoleUpdate', $sceneCode);  //验证权限

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
        $data = $this->validate(__FUNCTION__); //参数验证并处理
        $sceneCode = getRequestScene();
        switch ($sceneCode) {
            case 'platformAdmin':
                $this->container->get(\App\Module\Logic\Auth\Role::class)->checkAuth('authRoleDelete', $sceneCode);  //验证权限

                $this->service->delete(['id' => $data['idArr']]);
                break;
            default:
                throwFailJson('39999999');
                break;
        }
    }
}
