<?php

declare(strict_types=1);

namespace App\Aspect;

use Hyperf\Di\Annotation\Inject;
use Psr\Container\ContainerInterface;

abstract class AbstractAspect extends \Hyperf\Di\Aop\AbstractAspect
{
    #[Inject]
    protected ContainerInterface $container;
}
