<?php

declare(strict_types=1);

namespace App\Module\Service;

use Hyperf\Di\Annotation\Inject;
use Psr\Container\ContainerInterface;

abstract class AbstractService
{
    #[Inject]
    protected ContainerInterface $container;

    //protected string $daoClassName = \App\Module\Db\Dao\Auth\Scene::class;
    protected string $daoClassName;   //dao类的路径，调用地方实例化对象。因dao类带有状态，使用依赖注入会污染进程环境

    public function __construct()
    {
        //子类未定义$daoClassName时会自动设置。注意：Dao类目录和Service目录的对应关系
        if (empty($this->daoClassName)) {
            $this->daoClassName = str_replace('\\Service\\', '\\Db\\Dao\\', get_class($this));
        }
    }

    /**
     * 获取当前Dao实例
     *
     * @return \App\Module\Db\Dao\AbstractDao
     */
    final protected function getDao(): \App\Module\Db\Dao\AbstractDao
    {
        return getDao($this->daoClassName);
    }

    /**
     * 列表（通用。需要特殊处理的，子类重新定义即可）
     * 
     * @param array $field
     * @param array $where
     * @param array $order
     * @param integer $page
     * @param integer $limit
     * @return void
     */
    public function list(array $field = [], array $where = [], array $order = [], int $page = 1, int $limit = 10)
    {
        $dao = $this->getDao();
        $dao->where($where);
        $offset = ($page - 1) * $limit;

        empty($order) ? $order = ['id' => 'DESC'] : null;
        $dao->field($field)->order($order);
        if ($dao->isJoin()) {
            $dao->group(['id']);
        }
        $list = $dao->getList($offset, $limit);
        throwSuccessJson(['list' => $list]);
    }
    /* //重新定义示例
    public function list(array $field = [], array $where = [], array $order = [], int $page = 1, int $limit = 10)
    {
        try {
            parent::list(...func_get_args());
        } catch (\App\Exception\Json $th) {
            $responseData = $th->getResponseData();
            //数据处理
            $th->setApiData($responseData['data']);
            throw $th;
        }
    } */

    /**
     * 列表，带总数（通用。需要特殊处理的，子类重新定义即可）
     * 
     * @param array $field
     * @param array $where
     * @param array $order
     * @param integer $page
     * @param integer $limit
     * @return void
     */
    public function listWithCount(array $field = [], array $where = [], array $order = [], int $page = 1, int $limit = 10)
    {
        $offset = ($page - 1) * $limit;
        $dao = $this->getDao();
        $dao->where($where);
        if ($offset == 0 && $limit == 0) {  //是否先获取$list，再通过count($list)计算$count
            empty($order) ? $order = ['id' => 'DESC'] : null;
            $dao->field($field)->order($order);
            if ($dao->isJoin()) {
                $dao->group(['id']);
            }
            $list = $dao->getList($offset, $limit);
            $count = count($list);
        } else {
            if ($dao->isJoin()) {
                $count = $dao->getBuilder()->distinct()->count($dao->getTable() . '.' . $dao->getKey());
            } else {
                $count = $dao->getBuilder()->count();
            }

            $list = [];
            if ($count > $offset) {
                empty($order) ? $order = ['id' => 'DESC'] : null;
                $dao->field($field)->order($order);
                if ($dao->isJoin()) {
                    $dao->group(['id']);
                }
                $list = $dao->getList($offset, $limit);
            }
        }
        throwSuccessJson(['count' => $count, 'list' => $list]);
    }

    /**
     * 详情（通用。需要特殊处理的，子类重新定义即可）
     *
     * @param array $where
     * @param array $field
     * @return void
     */
    public function info(array $where, array $field = [])
    {
        $info = $this->getDao()->field($field)->where($where)->getInfo();
        if (empty($info)) {
            throwFailJson(29999999);
        }
        throwSuccessJson(['info' => $info]);
    }

    /**
     * 创建（通用。需要特殊处理的，子类重新定义即可）
     *
     * @param array $data
     * @return void
     */
    public function create(array $data)
    {
        $id = $this->getDao()->insert($data)->saveInsert();
        /* //重复索引错误已在\App\Exception\Handler\AppExceptionHandler处理
        try {
            $id = $this->getDao()->insert($data)->saveInsert();
        } catch (\Hyperf\Database\Exception\QueryException $th) {
            if (preg_match('/^SQLSTATE.*1062 Duplicate.*\.([^\']*)\'/', $th->getMessage(), $matches) === 1) {
                $nameKey = 'validation.attributes.' . $matches[1];
                $name =  trans($nameKey);
                if ($name === $nameKey) {
                    throwFailJson(29991062);
                } else {
                    throwFailJson(29991063, trans('code.29991063', ['name' => $name]));
                }
            }
        } */
        if (empty($id)) {
            throwFailJson();
        }
        throwSuccessJson();
    }

    /**
     * 更新（通用。需要特殊处理的，子类重新定义即可）
     *
     * @param array $data
     * @param array $where
     * @return void
     */
    public function update(array $data, array $where)
    {
        $result = $this->getDao()->where($where)->update($data)->saveUpdate();
        if (empty($result)) {
            throwFailJson();
        }
        throwSuccessJson();
    }

    /**
     * 删除（通用。需要特殊处理的，子类重新定义即可）
     *
     * @param array $where
     * @return void
     */
    public function delete(array $where)
    {
        $result = $this->getDao()->where($where)->delete();
        if (empty($result)) {
            throwFailJson();
        }
        throwSuccessJson();
    }

    /**
     * 获取更新|删除的id数组
     *
     * @param array $where
     * @return array
     */
    final protected function getIdArr(array $where): array
    {
        $dao = $this->getDao();
        return $dao->where($where)->getBuilder()->pluck($dao->getKey())->toArray();
        if (isset($where['id']) && count($where) == 1) {
            return is_array($where['id']) ? $where['id'] : [$where['id']];
        }
        if (isset($where['idArr']) && count($where) == 1) {
            return is_array($where['idArr']) ? $where['idArr'] : [$where['idArr']];
        }
    }
}
