<?php

declare(strict_types=1);

namespace App\Module\Validation;

class CommonList extends AbstractValidation
{
    protected array $rule =   [
        'field' => 'array',
        'where' => 'array',
        'order' => 'array',
        //'order.*' => 'in:asc,desc,ASC,DESC',
        'page' => 'integer|min:1',
        'limit' => 'integer|min:0',
    ];
}
