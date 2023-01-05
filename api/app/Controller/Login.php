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
        $data = $this->validate(__FUNCTION__); //参数验证并处理
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $this->service->encryptStr($data['account'], $sceneCode);
                break;
            default:
                throwFailJson('39999999');
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
        $data = $this->validate(__FUNCTION__); //参数验证并处理
        $sceneCode = $this->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platformAdmin':
                $this->service->login($data['account'], $data['password'], $sceneCode);
                break;
            default:
                throwFailJson('39999999');
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
                throwFailJson('39999999');
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
                throwFailJson('39999999');
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
                /* if ($loginInfo->adminId == 1) {
                    $where = [
                        'sceneId' => $this->request->sceneInfo->sceneId,
                        'isStop' => 0
                    ];
                } else {
                    $where = [
                        'adminId' => $loginInfo->adminId,
                        'isStop' => 0
                    ];
                } */
                $where = [
                    'sceneId' => 1,
                    'isStop' => 0
                ];
                $field = [
                    'menuTree',
                    'showMenu'
                ];
                $this->container->get(Menu::class)->tree($field, $where);
                break;
            default:
                throwFailJson('39999999');
                break;
        }
    }
}
