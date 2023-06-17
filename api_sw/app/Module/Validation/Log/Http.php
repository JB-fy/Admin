<?php

declare(strict_types=1);

namespace App\Module\Validation\Log;

use App\Module\Validation\AbstractValidation;

class Http extends AbstractValidation
{
    protected array $rule = [
        'httpId' => 'sometimes|required|integer|min:1',
        'url' => 'sometimes|required|url',
        'minRunTime' => 'sometimes|required|numeric|min:0',
        'maxRunTime' => 'sometimes|required|numeric|min:0|gte:minRunTime',
    ];

    protected array $scene = [];
}
