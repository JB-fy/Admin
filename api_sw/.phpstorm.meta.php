<?php

namespace PHPSTORM_META {
    // Reflect
    override(\Psr\Container\ContainerInterface::get(0), map(['' => '@']));
    override(\Hyperf\Context\Context::get(0), map(['' => '@']));
    override(\make(0), map(['' => '@']));
    override(\di(0), map(['' => '@']));
    override(\getDao(0), map(['' => '@']));
    override(\getCache(0), map(['' => '@']));
}
