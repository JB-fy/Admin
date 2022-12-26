<?php

declare(strict_types=1);

namespace App\Controller\Auth;

use App\Controller\AbstractController;

class Action extends AbstractController
{
    /**
     * 列表
     *
     * @return void
     */
    public function list()
    {
        $data = $this->validate(__FUNCTION__); //参数验证并处理
        switch (getRequestScene()) {
            case 'platformAdmin':
                $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getInfo('platformAdmin');
                /**--------验证权限 开始--------**/
                /* try {
                    $authActionCode = 'authActionLook';
                    $this->container->get(AuthService::class)->checkAuth($loginInfo, $authActionCode);
                    $isAuth = true;
                } catch (ApiException $e) {
                    $isAuth = false;
                } */
                /**--------验证权限 结束--------**/

                /**--------参数过滤 结束--------**/
                /* if ($isAuth) {
                    $allowField = getDao(AuthAction::class)->getAllColumn();
                    $allowField = array_merge($allowField, ['sceneName', 'pActionName']);
                } else {
                    //无查看权限时只能查看一些基本的字段
                    $allowField = ['menuId', 'menuName', 'menu'];
                } */

                /* $allowField = getDao(AuthAction::class)->getAllColumn();
                $allowField = array_merge($allowField, ['sceneName', 'pActionName']);
                $data['field'] = array_intersect($data['field'], $allowField); //过滤不可查看字段
                empty($data['field']) ? $data['field'] = $allowField : null; */
                /**--------参数过滤 结束--------**/

                $this->service->listWithCount(...$data);
                break;
            default:
                throwFailJson('001001');
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
        switch (getRequestScene()) {
            case 'platformAdmin':
                $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getInfo('platformAdmin');
                /**--------验证权限 开始--------**/
                /* $authActionCode = 'authActionInfo';
                $this->container->get(AuthService::class)->checkAuth($loginInfo, $authActionCode); */
                /**--------验证权限 结束--------**/

                $this->service->info(['id' => $data['id']]);
                break;
            default:
                throwFailJson('001001');
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
        switch (getRequestScene()) {
            case 'platformAdmin':
                $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getInfo('platformAdmin');
                /**--------验证权限 开始--------**/
                /* $authActionCode = 'authActionCreate';
                $this->container->get(AuthService::class)->checkAuth($loginInfo, $authActionCode); */
                /**--------验证权限 结束--------**/

                $this->service->create($data);
                break;
            default:
                throwFailJson('001001');
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
        switch (getRequestScene()) {
            case 'platformAdmin':
                $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getInfo('platformAdmin');
                /**--------验证权限 开始--------**/
                /* $authActionCode = 'authActionUpdate';
                $this->container->get(AuthService::class)->checkAuth($loginInfo, $authActionCode); */
                /**--------验证权限 结束--------**/

                $this->service->update($data, ['id' => $data['id']]);
                break;
            default:
                throwFailJson('001001');
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
        switch (getRequestScene()) {
            case 'platformAdmin':
                $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getInfo('platformAdmin');
                /**--------验证权限 开始--------**/
                /* $authActionCode = 'authActionInfo';
                $this->container->get(AuthService::class)->checkAuth($loginInfo, $authActionCode); */
                /**--------验证权限 结束--------**/

                $this->service->delete(['id' => $data['idArr']]);
                break;
            default:
                throwFailJson('001001');
                break;
        }
    }
}
