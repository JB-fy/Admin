<?php

declare(strict_types=1);
/**
 * This file is part of Hyperf.
 *
 * @link     https://www.hyperf.io
 * @document https://hyperf.wiki
 * @contact  group@hyperf.io
 * @license  https://github.com/hyperf/hyperf/blob/master/LICENSE
 */

use Hyperf\HttpServer\Router\Router;

/**--------公共接口 开始--------**/
Router::get('/', [\App\Controller\Index::class, 'index']);

Router::addRoute(['GET', 'POST', 'OPTIONS'], '/test', [\App\Controller\Test::class, 'index']);

Router::addGroup('/upload', function () {
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/notify', [\App\Controller\Upload::class, 'notify']);
});

Router::get('/favicon.ico', function () {
    return '';
});
/**--------公共接口 结束--------**/

/**--------平台后台接口 开始--------**/
Router::addGroup('/platform', function () {
    //做日志记录
    Router::addGroup('', function () {
        Router::addGroup('', function () {
            //无需验证登录身份
            Router::addGroup('', function () {
                Router::addGroup('/login', function () {
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/salt', [\App\Controller\Platform\Login\PlatformAdmin::class, 'salt']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/login', [\App\Controller\Platform\Login\PlatformAdmin::class, 'login']);
                });
            });

            //需验证登录身份
            Router::addGroup('', function () {
                Router::addGroup('/upload', function () {
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/sign', [\App\Controller\Upload::class, 'sign']);
                });

                Router::addGroup('/my', function () {
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/profile/info', [\App\Controller\Platform\My\Profile::class, 'info']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/profile/update', [\App\Controller\Platform\My\Profile::class, 'update']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/menu/tree', [\App\Controller\Platform\My\Menu::class, 'tree']);
                });

                Router::addGroup('/auth/action', function () {
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Platform\Auth\Action::class, 'list']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Platform\Auth\Action::class, 'info']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Platform\Auth\Action::class, 'create']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Platform\Auth\Action::class, 'update']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/del', [\App\Controller\Platform\Auth\Action::class, 'delete']);
                });

                Router::addGroup('/auth/menu', function () {
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Platform\Auth\Menu::class, 'list']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Platform\Auth\Menu::class, 'info']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Platform\Auth\Menu::class, 'create']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Platform\Auth\Menu::class, 'update']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/del', [\App\Controller\Platform\Auth\Menu::class, 'delete']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/tree', [\App\Controller\Platform\Auth\Menu::class, 'tree']);
                });

                Router::addGroup('/auth/role', function () {
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Platform\Auth\Role::class, 'list']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Platform\Auth\Role::class, 'info']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Platform\Auth\Role::class, 'create']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Platform\Auth\Role::class, 'update']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/del', [\App\Controller\Platform\Auth\Role::class, 'delete']);
                });

                Router::addGroup('/auth/scene', function () {
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Platform\Auth\Scene::class, 'list']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Platform\Auth\Scene::class, 'info']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Platform\Auth\Scene::class, 'create']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Platform\Auth\Scene::class, 'update']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/del', [\App\Controller\Platform\Auth\Scene::class, 'delete']);
                });

                Router::addGroup('/platform/admin', function () {
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Platform\Platform\Admin::class, 'list']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Platform\Platform\Admin::class, 'info']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Platform\Platform\Admin::class, 'create']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Platform\Platform\Admin::class, 'update']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/del', [\App\Controller\Platform\Platform\Admin::class, 'delete']);
                });

                Router::addGroup('/platform/config', function () {
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/get', [\App\Controller\Platform\Platform\Config::class, 'get']);
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/save', [\App\Controller\Platform\Platform\Config::class, 'save']);
                });

                Router::addGroup('/platform/server', function () {
                    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Platform\Platform\Server::class, 'list']);
                });
            }, ['middleware' => [\App\Middleware\SceneLoginOfPlatform::class]]);
        }, ['middleware' => [\App\Middleware\Scene::class]]);
    }, ['middleware' => [\App\Middleware\LogHttp::class]]);
});
/**--------平台后台接口 结束--------**/
