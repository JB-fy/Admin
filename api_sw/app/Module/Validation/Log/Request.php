<?php

declare(strict_types=1);

namespace App\Module\Validation\Log;

use App\Module\Validation\AbstractValidation;

class Request extends AbstractValidation
{
    protected array $rule = [
        'logId' => 'sometimes|required|integer|min:1',
        'requestUrl' => 'sometimes|required|url',
        'minRunTime' => 'sometimes|required|numeric|min:0',
        'maxRunTime' => 'sometimes|required|numeric|min:0|gte:minRunTime',
    ];

    protected array $scene = [];
}
