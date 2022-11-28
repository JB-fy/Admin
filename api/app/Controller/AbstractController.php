<?php

declare(strict_types=1);

namespace App\Controller;

use Hyperf\Di\Annotation\Inject;
use Hyperf\HttpServer\Contract\RequestInterface;
use Psr\Container\ContainerInterface;

abstract class AbstractController
{
    #[Inject]
    protected ContainerInterface $container;

    #[Inject]
    protected RequestInterface $request;

    //#[Inject(value:\App\Module\Service\Login::class)]
    //protected \App\Module\Service\AbstractService $service;   //代码会报红
    protected $service;

    //#[Inject(value:\App\Module\Validation\Login::class)]
    protected \App\Module\Validation\AbstractValidation $validation;

    public function __construct()
    {
        //子类未定义时会自动设置。注意：目录的对应关系
        if (empty($this->service)) {
            $serviceClassName = str_replace('\\Controller\\', '\\Module\\Service\\', get_class($this));
            if (class_exists($serviceClassName)) {
                $this->service = $this->container->get($serviceClassName);
            }
        }
        if (empty($this->validation)) {
            $validationClassName = str_replace('\\Controller\\', '\\Module\\Validation\\', get_class($this));
            if (class_exists($validationClassName)) {
                $this->validation = $this->container->get($validationClassName);
            }
        }
    }

    /**
     * 列表参数验证并处理
     * 
     * @return array
     */
    final protected function listParamVatetion(): array
    {
        $data = $this->request->all();
        $this->container->get(\App\Module\Validation\CommonList::class)->make($data)->validate();
        !isset($data['page']) ?: $data['page'] = (int)$data['page'];
        !isset($data['limit']) ?: $data['limit'] = (int)$data['limit'];

        $this->validation->make($data['where'], 'list')->validate();
        return $data;
    }
}
