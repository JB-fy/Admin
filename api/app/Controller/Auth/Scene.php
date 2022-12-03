<?php

declare(strict_types=1);

namespace App\Controller\Auth;

use App\Controller\AbstractController;

class Scene extends AbstractController
{
    public function list()
    {
        $data = $this->listVatetion(); //参数验证并处理

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
                    $allowField = container(DaoAuthScene::class, true)->getAllColumn();
                    //$allowField = array_merge($allowField, ['pMenuName', 'menuActionJson']);
                } else {
                    //无查看权限时只能查看一些基本的字段
                    $allowField = ['sceneId', 'sceneName', 'scene'];
                }

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
}
