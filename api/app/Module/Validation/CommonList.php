<?php

declare(strict_types=1);

namespace App\Module\Validation;

class CommonList extends AbstractValidation
{
    protected array $rule =   [
        'field' => 'sometimes|required_if_null|array',
        'where' => 'sometimes|required_if_null|array',
        'order' => 'sometimes|required_if_null|array',
        'order.*' => 'sometimes|required_if_null|in:asc,desc,ASC,DESC',
        'page' => 'sometimes|required_if_null|integer|min:1',
        'limit' => 'sometimes|required_if_null|integer|min:0',
    ];
}
