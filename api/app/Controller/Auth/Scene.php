<?php

declare(strict_types=1);

namespace App\Controller\Auth;

use App\Controller\AbstractController;
use Hyperf\Di\Annotation\Inject;

class Scene extends AbstractController
{
    #[Inject()]
    protected \App\Module\Service\Auth\Scene $service;

    /* #[Inject]
    protected \App\Module\Validation\Login $validation; */

    public function list()
    {
        /**--------参数验证并处理 开始--------**/
        $data = $this->request->all();
        /* $data = [
            'field' => $request->input('field', []),
            'where' => $request->input('where', []),
            'order' => $request->input('order', []),
            'page' => $request->input('page', 1),
            'limit' => $request->input('limit', 10),
        ]; */
        $this->container->get(\App\Module\Validate\CommonList::class)->make($data)->validate();

        /* $filterRules = [
            'id' => 'integer|min:1',
            'excId' => 'integer|min:1',
            'authMenuId' => 'integer|min:1',
            'scene' => 'in:' . implode(',', array_keys($this->translator->trans('const.scene'))),
            'pid' => 'integer|min:0',
            'menuName' => 'alpha_dash|between:1,30',
            'isStop' => 'in:' . implode(',', array_keys($this->translator->trans('const.yesOrNo'))),
        ];
        $this->validation->make($data, 'encryptStr')->validate(); */

        /* if (!isset($data['order']) || empty($data['order'])) {
            $data['order'] = [
                'id' => 'desc',
            ];
        } */
        /**--------参数验证并处理 结束--------**/

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

                $this->service->list(...$data);
                //$this->service->list($data['field'], $data['where'], $data['order'], $data['page'], $data['limit']);
                break;
            default:
                throwFailJson('001001');
                break;
        }
    }
}
