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
     * 获取验证场景
     * 
     * @param string $funcName
     * @param string $sceneCode
     * @return string
     */
    final protected function getValidateSceneName(string $funcName, string $sceneCode = ''): string
    {
        //$funcName === 'tree' ? $funcName = 'list' : null;
        $sceneName = ($sceneCode === '' || $sceneCode == 'platformAdmin') ? $funcName : $funcName . 'Of' . ucfirst($sceneCode);
        return $sceneName;
    }

    /**
     * 参数验证并处理
     * 
     * @param string $funcName
     * @param string $sceneCode 需要区分场景时使用
     * @return array
     */
    final protected function validate(string $funcName, string $sceneCode = ''): array
    {
        $data = $this->request->all();
        switch ($funcName) {
            case 'tree':
            case 'list':
                if (!empty($data)) {
                    //$data =  $this->container->get(\App\Module\Validation\CommonList::class)->make($data)->validated();  //不存在的字段不验证。相当于加sometimes规则
                    $data =  $this->container->get(\App\Module\Validation\CommonList::class)->make($data)->validate();
                    !isset($data['page']) ?: $data['page'] = (int)$data['page'];
                    !isset($data['limit']) ?: $data['limit'] = (int)$data['limit'];

                    if (!empty($data['where'])) {
                        $sceneName = $this->getValidateSceneName($funcName, $sceneCode);
                        $data['where'] = $this->validation->make($data['where'], $sceneName)->validate();
                    }
                }
                break;
            case 'update':
                $sceneName = $this->getValidateSceneName($funcName, $sceneCode);
                $data = $this->validation->make($data, $sceneName)->validate();
                if (count($data) < 2) { //更新除了id还必须有其他参数，所以至少需要两个参数
                    throwFailJson('89999999');
                }
                break;
            case 'info':
            case 'create':
            case 'delete':
            case 'get':
            case 'save':
            default:
                $sceneName = $this->getValidateSceneName($funcName, $sceneCode);
                $data = $this->validation->make($data, $sceneName)->validate();
                break;
        }
        return $data;
    }

    /**
     * 判断操作权限
     * 
     * @param string $funcName
     * @param string $sceneCode
     * @param boolean $isThrow
     * @return boolean
     */
    final protected function checkAuth(string $funcName, string $sceneCode, bool $isThrow = true): bool
    {
        switch ($funcName) {
            case 'list':
            case 'info':
            case 'tree':
            case 'get':
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
     * @param string $className
     * @return array
     */
    protected function getAllowField(string $className): array
    {
        $allowField = getDao($className)->getAllColumn();
        $allowField = array_merge($allowField, ['id']);
        $allowField = array_diff($allowField, ['password']);
        return $allowField;
    }
}
