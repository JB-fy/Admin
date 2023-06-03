<?php

declare(strict_types=1);

namespace App\Module\Validation;

class Common extends AbstractValidation
{
    protected array $rule =   [
        'field' => 'sometimes|required_if_null|array',
        'field.*' => 'sometimes|required',
        'filter' => 'sometimes|required_if_null|array',
        'filter.*' => 'sometimes|required',
        'order' => 'sometimes|required_if_null|array',
        'order.*' => 'sometimes|required|in:asc,desc,ASC,DESC',
        'page' => 'sometimes|required|integer|min:1',
        'limit' => 'sometimes|required|integer|min:0',
    ];

    protected array $scene = [
        'list' => [],
        'tree' => [
            'only' => [
                'field',
                'field.*',
                'filter',
                'filter.*',
            ]
        ]
    ];
}
