<?php

declare(strict_types=1);

namespace App\Module\Service;

use Hyperf\Di\Annotation\Inject;
use Psr\Container\ContainerInterface;

abstract class AbstractService
{
    #[Inject]
    protected ContainerInterface $container;

    //protected $daoClassName = \app\module\db\dao\auth\AuthMenu::class;
    protected $daoClassName;   //dao类的路径，调用地方实例化对象。因dao类带有状态，使用依赖注入会污染进程环境

    /**
     * 列表（通用，需要特殊处理的覆盖重新定义）
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
        empty($order) ? $order = ['id' => 'desc'] : null;
        $offset = ($page - 1) * $limit;

        $countAfter = ($offset == 0 && $limit == 0);  //用于判断是否先获取$list，再通过count($list)计算$count
        $dao = getDao($this->daoClassName);
        $dao->where($where);
        if (!$countAfter) {
            if ($dao->isJoin()) {
                $count = $dao->getBuilder()->distinct()->count($dao->getTable() . '.' . $dao->getKey());
            } else {
                $count = $dao->getBuilder()->count();
            }
        }

        $list = [];
        if ($countAfter || $count > $offset) {
            $dao->field($field)->order($order);
            if ($dao->isJoin()) {
                $dao->group(['id']);
            }
            $list = $dao->getList($offset, $limit);
            if ($countAfter) {
                $count = count($list);
            }
        }

        throwSuccessJson(['count' => $count, 'list' => $list]);
    }

    /**
     * 创建
     *
     * @param array $data
     * @return void
     */
    public function create(array $data)
    {
        $dao = getDao($this->daoClassName);
        $id = $dao->insert($data)->saveInsert();
        if (empty($id)) {
            throwFailJson('999999');
        }
        throwSuccessJson();
    }

    /**
     * 修改
     *
     * @param array $data
     * @param integer $id
     * @return void
     */
    public function update(array $data, int $id)
    {
        $dao = getDao($this->daoClassName);
        $result = $dao->where(['id' => $id])->update($data)->saveUpdate();
        if (empty($result)) {
            throwFailJson('999999');
        }
        throwSuccessJson();
    }

    /**
     * 删除
     *
     * @param array $idArr
     * @return void
     */
    public function delete(array $idArr)
    {
        $dao = getDao($this->daoClassName);
        $result = $dao->where([['id', 'in', $idArr]])->delete();
        if (empty($result)) {
            throwFailJson('999999');
        }
        throwSuccessJson();
    }
}
