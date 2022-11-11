<?php

declare(strict_types=1);

namespace app\module\validate;

class CommonList extends AbstractValidate
{
    protected $rule =   [
        'field|字段' => 'array',
        'where|条件' => 'array',
        'order|排序' => 'array',
        //'order.*|排序方式' => 'in:asc,desc,ASC,DESC',
        /* 'page|页码' => 'integer|min:1',  //框架integral规则不支持php8
        'limit|条目' => 'integer|min:0', */
    ];

    protected $message  =   [
        'field.array' => '字段必须是数组',
        'where.array' => '条件必须是数组',
        'order.array' => '排序必须是数组',
        'page.integer' => '页码必须是整数且大于0',
        'limit.integer' => '条目必须是整数且大于等于0',
    ];
}
