<?php

declare(strict_types=1);

namespace App\Module\Validation;

class CommonList extends AbstractValidation
{
    protected array $rule =   [
        'field' => 'sometimes|required_if_null|array',
        'where' => 'sometimes|required_if_null|array',
        'order' => 'sometimes|required_if_null|array',
        'order.*' => 'sometimes|required|in:asc,desc,ASC,DESC',
        'page' => 'sometimes|required|integer|min:1',
        'limit' => 'sometimes|required|integer|min:0',
    ];
}
