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

Router::addGroup('/login',function (){
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/encryptStr', [\App\Controller\Login::class, 'encryptStr']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '', [\App\Controller\Login::class, 'login']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/info', [\App\Controller\Login::class, 'info']);
    // Router::addRoute(['GET', 'POST', 'OPTIONS'], '/update', [\App\Controller\Login::class, 'update']);
    Router::addRoute(['GET', 'POST', 'OPTIONS'], '/menuTree', [\App\Controller\Login::class, 'menuTree']);
});

Router::addGroup('/auth/scene/', function () {
    Router::addRoute(['GET', 'POST', 'OPTIONS'], 'list', [\App\Controller\Auth\AuthScene::class, 'list']);
});

Router::get('/favicon.ico', function () {
    return '';
});
