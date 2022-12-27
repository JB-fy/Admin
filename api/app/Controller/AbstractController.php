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
     * 参数验证并处理
     * 
     * @param string $sceneName
     * @return array
     */
    final protected function validate(string $sceneName): array
    {
        $data = $this->request->all();
        switch ($sceneName) {
            case 'list':
                if (!empty($data)) {
                    //$data =  $this->container->get(\App\Module\Validation\CommonList::class)->make($data)->validated();  //不存在的字段不验证。相当于加sometimes规则
                    $data =  $this->container->get(\App\Module\Validation\CommonList::class)->make($data)->validate();
                    !isset($data['page']) ?: $data['page'] = (int)$data['page'];
                    !isset($data['limit']) ?: $data['limit'] = (int)$data['limit'];

                    if (!empty($data['where'])) {
                        $data['where'] = $this->validation->make($data['where'], $sceneName)->validate();
                    }
                }
                break;
            case 'create':
                $data = $this->validation->make($data, $sceneName)->validate();
                $data = $this->handleData($data);
                break;
            case 'update':
                $data = $this->validation->make($data, $sceneName)->validate();
                if (count($data) < 2) { //更新除了id还必须有其他参数，所以至少需要两个参数
                    throwFailJson('89999998');
                }
                $data = $this->handleData($data);
                break;
            case 'info':
            case 'delete':
            default:
                $data = $this->validation->make($data, $sceneName)->validate();
                break;
        }
        return $data;
    }

    /**
     * 创建更新时的参数处理（通用。需要特殊处理的，子类重新定义即可）
     *
     * @param array $data
     * @return array
     */
    protected function handleData(array $data): array
    {
        return $data;
    }
}
