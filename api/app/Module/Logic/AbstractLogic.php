<?php

declare(strict_types=1);

namespace App\Module\Logic;

use Hyperf\Di\Annotation\Inject;
use Psr\Container\ContainerInterface;

abstract class AbstractLogic
{
    #[Inject]
    protected ContainerInterface $container;
}
