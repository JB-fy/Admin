<?php

declare(strict_types=1);

namespace App\Listener;

use Hyperf\Crontab\Event\FailToExecute;
use Hyperf\Event\Annotation\Listener;
use Hyperf\Event\Contract\ListenerInterface;

//#[Listener]
class CrontabListener implements ListenerInterface
{
    public function listen(): array
    {
        return [
            FailToExecute::class,
        ];
    }

    public function process(object $event): void
    {
        if ($event instanceof FailToExecute) {
            var_dump($event->crontab->getName());
            var_dump($event->throwable->getMessage());
        }
    }
}
