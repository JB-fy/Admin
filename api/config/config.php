<?php

use Hyperf\Di\Annotation\AspectCollector;

return [
    'annotations' => [
        'scan' => [
            'paths' => [
                BASE_PATH . '/app',
            ],
            'ignore_annotations' => [
                'mixin',
            ],
            'class_map' => [],
            'collectors' => [
                AspectCollector::class
            ],
        ],
    ],
    'aspects' => [
        //\app\aspect\Access::class,  //跨域组件可以去掉。可直接在nginx中设置，或在support\Response中全局设置
        \app\aspect\Language::class,
        \app\aspect\LogOfRequest::class,
        \app\aspect\AuthScene::class,
        \app\aspect\AuthSceneOfSystemAdmin::class,
        //\app\aspect\translator\AuthMenu::class,
    ]
];
