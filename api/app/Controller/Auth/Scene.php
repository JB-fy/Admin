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
        $data = $this->validate(__FUNCTION__); //参数验证并处理
        switch (getRequestScene()) {
            case 'platformAdmin':
                $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getInfo('platformAdmin');
                /**--------验证权限 开始--------**/
                /* try {
                    $authActionCode = 'authSceneLook';
                    $this->container->get(AuthService::class)->checkAuth($loginInfo, $authActionCode);
                    $isAuth = true;
                } catch (ApiException $e) {
                    $isAuth = false;
                } */
                /**--------验证权限 结束--------**/

                /**--------参数过滤 结束--------**/
                /* if ($isAuth) {
                    $allowField = getDao(DaoAuthScene::class, true)->getAllColumn();
                    //$allowField = array_merge($allowField, ['pMenuName', 'menuActionJson']);
                } else {
                    //无查看权限时只能查看一些基本的字段
                    $allowField = ['sceneId', 'sceneName', 'scene'];
                } */

                $allowField = getDao(AuthScene::class)->getAllColumn();
                $allowField = array_merge($allowField, ['id']);
                $data['field'] = empty($data['field']) ? $allowField : array_intersect($data['field'], $allowField);    //过滤不可查看字段
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
                /* $authActionCode = 'authSceneInfo';
                $this->container->get(AuthService::class)->checkAuth($loginInfo, $authActionCode); */
                /**--------验证权限 结束--------**/

                $allowField = getDao(AuthScene::class)->getAllColumn();
                $allowField = array_merge($allowField, ['id']);
                $data['field'] = empty($data['field']) ? $allowField : array_intersect($data['field'], $allowField);    //过滤不可查看字段
                $this->service->info(['id' => $data['id']], $data['field']);
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
                /* $authActionCode = 'authSceneCreate';
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
                /* $authActionCode = 'authSceneUpdate';
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
                /* $authActionCode = 'authSceneInfo';
                $this->container->get(AuthService::class)->checkAuth($loginInfo, $authActionCode); */
                /**--------验证权限 结束--------**/

                $this->service->delete(['id' => $data['idArr']]);
                break;
            default:
                throwFailJson('001001');
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
