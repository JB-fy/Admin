<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Contract\ContainerInterface;
use Hyperf\Di\Annotation\Inject;

abstract class AbstractAspect extends \Hyperf\Di\Aop\AbstractAspect
{
    #[Inject]
    protected ContainerInterface $container;
}
