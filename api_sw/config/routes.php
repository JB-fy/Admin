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
//不做日志记录
Router::addGroup('', function () {
    Router::addGroup('/platformAdmin', function () {
        //需验证登录身份
        Router::addGroup('', function () {
            Router::addGroup('/log/request', function () {
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Log\Request::class, 'list']);
            });
        }, ['middleware' => [\App\Middleware\ScenePlatformAdminLogin::class]]);
    }, ['middleware' => [\App\Middleware\ScenePlatformAdmin::class]]);
});

//做日志记录
Router::addGroup('', function () {
    Router::addGroup('/platformAdmin', function () {
        //无需验证登录身份
        Router::addGroup('', function () {
            Router::addGroup('/login', function () {
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/encryptStr', [\App\Controller\Login::class, 'encryptStr']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '', [\App\Controller\Login::class, 'login']);
            });
        });

        //需验证登录身份
        Router::addGroup('', function () {
            Router::addGroup('/upload', function () {
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/sign', [\App\Controller\Upload::class, 'sign']);
            });

            Router::addGroup('/login', function () {
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Login::class, 'info']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Login::class, 'update']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/menuTree', [\App\Controller\Login::class, 'menuTree']);
            });

            Router::addGroup('/auth/action', function () {
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Auth\Action::class, 'list']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Auth\Action::class, 'info']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Auth\Action::class, 'create']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Auth\Action::class, 'update']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/del', [\App\Controller\Auth\Action::class, 'delete']);
            });

            Router::addGroup('/auth/menu', function () {
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Auth\Menu::class, 'list']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Auth\Menu::class, 'info']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Auth\Menu::class, 'create']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Auth\Menu::class, 'update']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/del', [\App\Controller\Auth\Menu::class, 'delete']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/tree', [\App\Controller\Auth\Menu::class, 'tree']);
            });

            Router::addGroup('/auth/role', function () {
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Auth\Role::class, 'list']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Auth\Role::class, 'info']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Auth\Role::class, 'create']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Auth\Role::class, 'update']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/del', [\App\Controller\Auth\Role::class, 'delete']);
            });

            Router::addGroup('/auth/scene', function () {
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Auth\Scene::class, 'list']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Auth\Scene::class, 'info']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Auth\Scene::class, 'create']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Auth\Scene::class, 'update']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/del', [\App\Controller\Auth\Scene::class, 'delete']);
            });

            Router::addGroup('/platform/admin', function () {
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Platform\Admin::class, 'list']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Platform\Admin::class, 'info']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Platform\Admin::class, 'create']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Platform\Admin::class, 'update']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/del', [\App\Controller\Platform\Admin::class, 'delete']);
            });

            Router::addGroup('/platform/config', function () {
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/get', [\App\Controller\Platform\Config::class, 'get']);
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/save', [\App\Controller\Platform\Config::class, 'save']);
            });

            Router::addGroup('/platform/server', function () {
                Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Platform\Server::class, 'list']);
            });
        }, ['middleware' => [\App\Middleware\ScenePlatformAdminLogin::class]]);
    }, ['middleware' => [\App\Middleware\ScenePlatformAdmin::class]]);
}, ['middleware' => [\App\Middleware\LogRequest::class]]);
/**--------平台后台接口 结束--------**/
