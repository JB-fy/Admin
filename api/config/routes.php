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

Router::get('/', [\App\Controller\Index::class, 'index']);

Router::addRoute(['GET', 'POST', 'OPTIONS'], '/test', [\App\Controller\Test::class, 'index']);

Router::addGroup('/login', function () {
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/encryptStr', [\App\Controller\Login::class, 'encryptStr']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '', [\App\Controller\Login::class, 'login']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Login::class, 'info']);
    // Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Login::class, 'update']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/menuTree', [\App\Controller\Login::class, 'menuTree']);
});

Router::addGroup('/auth/action', function () {
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Auth\Action::class, 'list']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Auth\Action::class, 'info']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Auth\Action::class, 'create']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Auth\Action::class, 'update']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/delete', [\App\Controller\Auth\Action::class, 'delete']);
});

Router::addGroup('/auth/menu', function () {
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Auth\Menu::class, 'list']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Auth\Menu::class, 'info']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Auth\Menu::class, 'create']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Auth\Menu::class, 'update']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/delete', [\App\Controller\Auth\Menu::class, 'delete']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/tree', [\App\Controller\Auth\Menu::class, 'tree']);
});

Router::addGroup('/auth/role', function () {
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Auth\Role::class, 'list']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Auth\Role::class, 'info']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Auth\Role::class, 'create']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Auth\Role::class, 'update']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/delete', [\App\Controller\Auth\Role::class, 'delete']);
});

Router::addGroup('/auth/scene', function () {
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/list', [\App\Controller\Auth\Scene::class, 'list']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Auth\Scene::class, 'info']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/create', [\App\Controller\Auth\Scene::class, 'create']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Auth\Scene::class, 'update']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/delete', [\App\Controller\Auth\Scene::class, 'delete']);
});

Router::get('/favicon.ico', function () {
    return '';
});
