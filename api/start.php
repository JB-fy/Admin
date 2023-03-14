#!/usr/bin/env php
<?php
require_once __DIR__ . '/vendor/autoload.php';
use Hyperf\AopIntegration\ClassLoader;
ClassLoader::init();
support\App::run();
