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

    //操作标识前缀
    protected $actionCodePrefix;

    public function __construct()
    {
        $className = get_class($this);
        //子类未定义时会自动设置。注意：目录的对应关系
        if (empty($this->service)) {
            $serviceClassName = str_replace('\\Controller\\', '\\Module\\Service\\', $className);
            if (class_exists($serviceClassName)) {
                $this->service = $this->container->get($serviceClassName);
            }
        }
        if (empty($this->validation)) {
            $validationClassName = str_replace('\\Controller\\', '\\Module\\Validation\\', $className);
            if (class_exists($validationClassName)) {
                $this->validation = $this->container->get($validationClassName);
            }
        }
        if (empty($this->actionCodePrefix)) {
            $this->actionCodePrefix = lcfirst(str_replace('\\', '',  substr($className, strpos($className, '\\Controller\\') + strlen('\\Controller\\'))));
        }
    }

    /**
     * 在当前请求中，获取场景标识
     * 
     * @return string|null
     */
    public function getCurrentSceneCode(): ?string
    {
        return $this->container->get(\App\Module\Logic\Auth\Scene::class)->getCurrentSceneCode();
    }

    /**
     * 参数验证并处理
     * 
     * @param string $funcName
     * @return array
     */
    final protected function validate(string $funcName): array
    {
        $data = $this->request->all();
        switch ($funcName) {
            case 'list':
                if (!empty($data)) {
                    //$data =  $this->container->get(\App\Module\Validation\CommonList::class)->make($data)->validated();  //不存在的字段不验证。相当于加sometimes规则
                    $data =  $this->container->get(\App\Module\Validation\CommonList::class)->make($data)->validate();
                    !isset($data['page']) ?: $data['page'] = (int)$data['page'];
                    !isset($data['limit']) ?: $data['limit'] = (int)$data['limit'];

                    if (!empty($data['where'])) {
                        $data['where'] = $this->validation->make($data['where'], $funcName)->validate();
                    }
                }
                break;
            case 'create':
                $data = $this->validation->make($data, $funcName)->validate();
                $data = $this->handleData($data);
                break;
            case 'update':
                $data = $this->validation->make($data, $funcName)->validate();
                if (count($data) < 2) { //更新除了id还必须有其他参数，所以至少需要两个参数
                    throwFailJson('89999999');
                }
                $data = $this->handleData($data);
                break;
            case 'info':
            case 'delete':
            case 'get':
            case 'save':
            default:
                $data = $this->validation->make($data, $funcName)->validate();
                break;
        }
        return $data;
    }

    /**
     * 判断操作权限
     * 
     * @param string $funcName
     * @return array
     */
    final protected function checkAuth(string $funcName, string $sceneCode, bool $isThrow = true): bool
    {
        switch ($funcName) {
            case 'get':
            case 'list':
            case 'info':
                return $this->container->get(\App\Module\Logic\Auth\Action::class)->checkAuth($this->actionCodePrefix . 'Look', $sceneCode, $isThrow);
                break;
            case 'create':
            case 'update':
            case 'delete':
            case 'save':
            default:
                return $this->container->get(\App\Module\Logic\Auth\Action::class)->checkAuth($this->actionCodePrefix . ucfirst($funcName), $sceneCode, $isThrow);
                break;
        }
    }

    /**
     * 获取允许查看的字段
     *
     * @param array $data
     * @return array
     */
    protected function getAllowField(string $className): array
    {
        $allowField = getDao($className)->getAllColumn();
        $allowField = array_merge($allowField, ['id']);
        $allowField = array_diff($allowField, ['password']);
        return $allowField;
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
