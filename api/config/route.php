<?php

/**
 * This file is part of webman.
 *
 * Licensed under The MIT License
 * For full copyright and license information, please see the MIT-LICENSE.txt
 * Redistributions of files must retain the above copyright notice.
 *
 * @author    walkor<walkor@workerman.net>
 * @copyright walkor<walkor@workerman.net>
 * @link      http://www.workerman.net/
 * @license   http://www.opensource.org/licenses/mit-license.php MIT License
 */

use Webman\Route;

/*Route::get('/favicon.ico', function () {
    return '';
});*/

Route::get('/test', [\app\controller\Test::class, 'index']);

Route::get('/', [\app\controller\Index::class, 'index']);

Route::group('/login', function () {
    Route::add(['GET', 'POST', 'OPTIONS'], '/getEncryptStr', [\app\controller\Login::class, 'getEncryptStr']);
    Route::add(['GET', 'POST', 'OPTIONS'], '', [\app\controller\Login::class, 'login']);
    Route::add(['GET', 'POST', 'OPTIONS'], '/getInfo', [\app\controller\Login::class, 'getInfo']);
    // Route::add(['GET', 'POST', 'OPTIONS'], '/updateInfo', [\app\controller\Login::class, 'updateInfo']);
    Route::add(['GET', 'POST', 'OPTIONS'], '/getMenuTree', [\app\controller\Login::class, 'getMenuTree']);
});

//当找不到路由时，处理方法
Route::fallback(function () {
    //return redirect('/');
    throwFailJson('000404');
});
//关闭默认路由，好处：路由集中，方便管理。且aop插件要求控制器必须在这里至少注册过一个路由，才能让该控制器内所有路由aop生效。
Route::disableDefaultRoute();

// Router::addRoute(['POST', 'OPTIONS'], '/admin/admin/getList', [\App\Controller\Admin\Admin::class, 'getList']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/admin/save', [\App\Controller\Admin\Admin::class, 'save']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/admin/del', [\App\Controller\Admin\Admin::class, 'del']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/authGroup/getList', [\App\Controller\Admin\AuthGroup::class, 'getList']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/authGroup/save', [\App\Controller\Admin\AuthGroup::class, 'save']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/authGroup/del', [\App\Controller\Admin\AuthGroup::class, 'del']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/authMenu/getList', [\App\Controller\Admin\AuthMenu::class, 'getList']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/authMenu/save', [\App\Controller\Admin\AuthMenu::class, 'save']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/authMenu/del', [\App\Controller\Admin\AuthMenu::class, 'del']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/authMenu/getTree', [\App\Controller\Admin\AuthMenu::class, 'getTree']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/authMenu/getSelfTree', [\App\Controller\Admin\AuthMenu::class, 'getSelfTree']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/config/save', [\App\Controller\Admin\Config::class, 'save']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/config/get', [\App\Controller\Admin\Config::class, 'get']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/config/getConst', [\App\Controller\Admin\Config::class, 'getConst']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/index/index', [\App\Controller\Admin\Index::class, 'index']);
// Router::addRoute(['POST', 'OPTIONS'], '/admin/logRequest/getList', [\App\Controller\Admin\LogRequest::class, 'getList']);