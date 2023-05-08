<?php

declare(strict_types=1);

namespace App\Controller;

use App\Module\Service\Auth\Menu;

class Login extends AbstractController
{
    /**
     * 获取登录加密字符串(前端登录操作用于加密密码后提交)
     *
     * @return void
     */
    public function encryptStr()
    {
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $data = $this->validate(__FUNCTION__, $sceneCode);
                $this->service->encryptStr($data['account'], $sceneCode);
                break;
            default:
                throwFailJson(39999999);
                break;
        }
    }

    /**
     * 登录
     *
     * @return void
     */
    public function login()
    {
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $data = $this->validate(__FUNCTION__, $sceneCode);
                $this->service->login($data['account'], $data['password'], $sceneCode);
                break;
            default:
                throwFailJson(39999999);
                break;
        }
    }

    /**
     * 获取登录用户信息
     *
     * @return void
     */
    public function info()
    {
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getCurrentInfo($sceneCode);
                throwSuccessJson(['info' => $loginInfo]);
                break;
            default:
                throwFailJson(39999999);
                break;
        }
    }

    /**
     * 修改个人信息
     *
     * @return void
     */
    public function updateInfo()
    {
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                /**--------参数验证并处理 开始--------**/
                $data = $this->request->all();
                $data = $this->container->get(\App\Module\Validation\Platform\Admin::class)->make($data, 'updateSelf')->validate();
                /**--------参数验证并处理 结束--------**/

                $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getCurrentInfo($sceneCode);
                $this->container->get(\App\Module\Service\Platform\Admin::class)->update($data, ['id' => $loginInfo->adminId]);
                break;
            default:
                throwFailJson(39999999);
                break;
        }
    }

    /**
     * 获取后台用户菜单树
     *
     * @return void
     */
    public function menuTree()
    {
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getCurrentInfo($sceneCode);
                $where = [
                    'selfMenu' => [
                        'sceneCode' => $sceneCode,
                        'loginId' => $loginInfo->adminId
                    ]
                ];
                $field = [
                    'menuTree',
                    'showMenu'
                ];
                $this->container->get(Menu::class)->tree($field, $where);
                break;
            default:
                throwFailJson(39999999);
                break;
        }
    }
}
