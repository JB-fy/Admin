<?php

declare(strict_types=1);

namespace app\controller;

use app\module\service\auth\AuthMenu as ServiceAuthMenu;
use app\module\service\Login as ServiceLogin;
use app\module\validate\Login as ValidateLogin;
use support\Request;

class Login extends AbstractController
{
    /**
     * 获取登录加密字符串(前端登录操作用于加密密码后提交)
     *
     * @param Request $request
     * @return void
     */
    public function encryptStr(Request $request)
    {
        switch ($request->authSceneInfo->sceneCode) {
            case 'systemAdmin':
                /**--------验证参数 开始--------**/
                $data = $request->all();
                container(ValidateLogin::class, true)->scene('encryptStr')->check($data);
                /**--------验证参数 结束--------**/

                container(ServiceLogin::class)->encryptStr($data['account'], 'systemAdmin');
                break;
            default:
                throwFailJson('001001');
                break;
        }
    }

    /**
     * 登录
     *
     * @param Request $request
     * @return void
     */
    public function login(Request $request)
    {
        switch ($request->authSceneInfo->sceneCode) {
            case 'systemAdmin':
                /**--------验证参数 开始--------**/
                $data = $request->all();
                container(ValidateLogin::class, true)->check($data);
                /**--------验证参数 结束--------**/

                container(ServiceLogin::class)->login($data['account'], $data['password'], 'systemAdmin');
                break;
            default:
                throwFailJson('001001');
                break;
        }
    }

    /**
     * 获取登录用户信息
     *
     * @param Request $request
     * @return void
     */
    public function info(Request $request)
    {
        switch ($request->authSceneInfo->sceneCode) {
            case 'systemAdmin':
                throwSuccessJson(['info' => $request->systemAdminInfo]);
                break;
            default:
                throwFailJson('001001');
                break;
        }
    }

    /**
     * 修改个人信息
     *
     * @param Request $request
     * @return void
     */
    // public function updateInfo(Request $request)
    // {
    //     switch ($request->authSceneInfo->sceneCode) {
    //         case 'systemAdmin':
    //             /**--------验证参数 开始--------**/
    //             $data = $request->all();
    //             container(ValidateLogin::class, true)->scene('encryptStr')->check($data);
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

    //             container(AdminService::class)->update($data, $request->systemAdminInfo->adminId);
    //             break;
    //         default:
    //             throwFailJson('001001');
    //             break;
    //     }
    // }

    /**
     * 获取后台用户菜单树
     *
     * @param Request $request
     * @return void
     */
    public function menuTree(Request $request)
    {
        switch ($request->authSceneInfo->sceneCode) {
            case 'systemAdmin':
                /* if ($request->systemAdminInfoo->adminId == 1) {
                    $where = [
                        'sceneId' => $request->authSceneInfo->sceneId,
                        'isStop' => 0
                    ];
                } else {
                    $where = [
                        'systemAdminId' => $request->systemAdminInfoo->adminId,
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
                container(ServiceAuthMenu::class)->tree($field, $where);
                break;
            default:
                throwFailJson('001001');
                break;
        }
    }
}
