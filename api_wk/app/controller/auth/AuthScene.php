<?php

declare(strict_types=1);

namespace app\controller\auth;

use app\controller\AbstractController;
use app\module\db\dao\auth\AuthScene as DaoAuthScene;
use app\module\service\auth\AuthScene as ServiceAuthScene;
use app\module\validate\CommonList;
use support\Request;

class AuthScene extends AbstractController
{
    public function list(Request $request)
    {
        /**--------参数验证并处理 开始--------**/
        $data = $request->all();
        /* $data = [
            'field' => $request->input('field', []),
            'where' => $request->input('where', []),
            'order' => $request->input('order', []),
            'page' => $request->input('page', 1),
            'limit' => $request->input('limit', 10),
        ]; */
        container(CommonList::class, true)->check($data);

        /* $filterRules = [
            'id' => 'integer|min:1',
            'excId' => 'integer|min:1',
            'authMenuId' => 'integer|min:1',
            'scene' => 'in:' . implode(',', array_keys($this->translator->trans('const.scene'))),
            'pid' => 'integer|min:0',
            'menuName' => 'alpha_dash|between:1,30',
            'isStop' => 'in:' . implode(',', array_keys($this->translator->trans('const.yesOrNo'))),
        ];
        container(ValidateLogin::class, true)->scene('encryptStr')->check($data['where']); */

        /* if (!isset($data['order']) || empty($data['order'])) {
            $data['order'] = [
                'id' => 'desc',
            ];
        } */
        /**--------参数验证并处理 结束--------**/

        switch ($request->authSceneInfo->sceneCode) {
            case 'systemAdmin':
                $loginInfo = $request->systemAdminInfo;
                /**--------验证权限 开始--------**/
                /* try {
                    $authActionCode = 'authSceneLook';
                    container(AuthService::class)->checkAuth($loginInfo, $authActionCode);
                    $isAuth = true;
                } catch (ApiException $e) {
                    $isAuth = false;
                } */
                /**--------验证权限 结束--------**/

                /**--------参数处理 结束--------**/
                /* if ($isAuth) {
                    $allowField = container(DaoAuthScene::class, true)->getAllColumn();
                    //$allowField = array_merge($allowField, ['pMenuName', 'menuActionJson']);
                } else {
                    //无查看权限时只能查看一些基本的字段
                    $allowField = ['sceneId', 'sceneName', 'scene'];
                }

                $data['field'] = array_intersect($data['field'], $allowField); //过滤不可查看字段
                empty($data['field']) ? $data['field'] = $allowField : null; */
                /**--------参数处理 结束--------**/

                container(ServiceAuthScene::class)->list(...$data);
                //container(ServiceAuthScene::class)->list($data['field'], $data['where'], $data['order'], $data['page'], $data['limit']);
                break;
            default:
                throwFailJson('001001');
                break;
        }
    }
}
