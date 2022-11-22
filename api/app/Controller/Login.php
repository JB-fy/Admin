<?php

declare(strict_types=1);

namespace App\Controller;

//use app\module\service\auth\AuthMenu as ServiceAuthMenu;
use app\module\service\Login as ServiceLogin;
use App\Module\Validation\Login as ValidationLogin;

class Login extends AbstractController
{
    /**
     * 获取登录加密字符串(前端登录操作用于加密密码后提交)
     *
     * @return void
     */
    public function encryptStr()
    {
        switch ($this->request->authSceneInfo->sceneCode) {
            case 'platform':
                /**--------验证参数 开始--------**/
                $data = $this->request->all();
                $this->container->get(ValidationLogin::class)->scene('encryptStr')->check($data);
                /**--------验证参数 结束--------**/

                $this->container->get(ServiceLogin::class)->encryptStr($data['account'], 'platform');
                break;
            default:
                throwFailJson('001001');
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
        switch ($this->request->authSceneInfo->sceneCode) {
            case 'platform':
                /**--------验证参数 开始--------**/
                $data = $this->request->all();
                $this->container->get(ValidationLogin::class)->check($data);
                /**--------验证参数 结束--------**/

                $this->container->get(ServiceLogin::class)->login($data['account'], $data['password'], 'platform');
                break;
            default:
                throwFailJson('001001');
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
        switch ($this->request->authSceneInfo->sceneCode) {
            case 'platform':
                throwSuccessJson(['info' => $this->request->platformInfo]);
                break;
            default:
                throwFailJson('001001');
                break;
        }
    }

    /**
     * 修改个人信息
     *
     * @return void
     */
    // public function updateInfo()
    // {
    //     switch ($this->request->authSceneInfo->sceneCode) {
    //         case 'platform':
    //             /**--------验证参数 开始--------**/
    //             $data = $this->request->all();
    //             $this->container->get(ValidationLogin::class, true)->scene('encryptStr')->check($data);
    //             /**--------验证参数 结束--------**/

    //             /**--------验证参数 开始--------**/
    //             $data = [];

    //             $this->request->post('nickname') === null ? null : $data['nickname'] = $this->request->post('nickname');
    //             $this->request->post('newPassword') === null ? null : $data['newPassword'] = $this->request->post('newPassword');
    //             $this->request->post('checkNewPassword') === null ? null : $data['checkNewPassword'] = $this->request->post('checkNewPassword');
    //             $this->request->post('oldPassword') === null ? null : $data['oldPassword'] = $this->request->post('oldPassword');

    //             $rules = [
    //                 'nickname' => 'between:1,30',
    //                 'newPassword' => 'size:32|different:oldPassword|same:checkNewPassword',
    //                 'checkNewPassword' => 'required_with:newPassword|size:32',
    //                 'oldPassword' => 'required_with:newPassword|size:32',
    //             ];
    //             $this->validator->validate($data, $rules);
    //             if (isset($data['newPassword'])) {
    //                 $data['password'] = $data['newPassword'];
    //                 unset($data['newPassword']);
    //                 unset($data['checkNewPassword']);
    //             }
    //             /**--------验证参数 结束--------**/

    //             $this->container->get(AdminService::class)->update($data, $this->request->platformInfo->adminId);
    //             break;
    //         default:
    //             throwFailJson('001001');
    //             break;
    //     }
    // }

    /**
     * 获取后台用户菜单树
     *
     * @return void
     */
    // public function menuTree()
    // {
    //     switch ($this->request->authSceneInfo->sceneCode) {
    //         case 'platform':
    //             /* if ($this->request->platformInfo->adminId == 1) {
    //                 $where = [
    //                     'sceneId' => $this->request->authSceneInfo->sceneId,
    //                     'isStop' => 0
    //                 ];
    //             } else {
    //                 $where = [
    //                     'platformId' => $this->request->platformInfo->adminId,
    //                     'isStop' => 0
    //                 ];
    //             } */
    //             $where = [
    //                 'sceneId' => 1,
    //                 'isStop' => 0
    //             ];
    //             $field = [
    //                 'menuTree',
    //                 'showMenu'
    //             ];
    //             $this->container->get(ServiceAuthMenu::class)->tree($field, $where);
    //             break;
    //         default:
    //             throwFailJson('001001');
    //             break;
    //     }
    // }
}
