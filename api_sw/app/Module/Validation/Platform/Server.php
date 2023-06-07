<?php

declare(strict_types=1);

namespace App\Module\Validation\Platform;

use App\Module\Validation\AbstractValidation;

class Server extends AbstractValidation
{
    protected array $rule = [
        'serverId' => 'sometimes|required|integer|min:1',
        'networkIp' => 'sometimes|required|string',
        'localIp' => 'sometimes|required|string',
    ];

    protected array $scene = [];
}
