<?php

declare(strict_types=1);

namespace app\module\service;
//use DI\Annotation\Inject;

abstract class AbstractService
{
    // /**
    //  * @Inject
    //  * @var \app\module\db\table\auth\AuthMenu
    //  */
    // protected $table;
    //protected $table = \app\module\db\table\auth\AuthMenu::class;
    protected $table = '';   //table类的路径，调用地方实例化对象。因table类带有状态，无法直接使用容器注释功能做依赖注入。

    /**
     * 列表
     * 
     * @param array $field
     * @param array $where
     * @param array $order
     * @param integer $offset
     * @param integer $limit
     * @return void
     */
    public function list(array $field = [], array $where = [], array $order = [], int $offset = 1, int $limit = 10)
    {
        $countAfter = ($offset == 0 && $limit == 0);  //用于判断是否先获取$list，再通过count($list)计算$count
        $table = container($this->table, true);
        $table->where($where);
        if (!$countAfter) {
            if ($table->isJoin()) {
                $count = $table->getBuilder()->distinct()->count($table->getTable() . '.' . $table->getKey());
            } else {
                $count = $table->getBuilder()->count();
            }
        }

        $list = [];
        if ($countAfter || $count > $offset) {
            $table->field($field)->order($order);
            if ($table->isJoin()) {
                $table->group(['id']);
            }
            $list = $table->getList($offset, $limit);
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
        $table = container($this->table, true);
        $id = $table->insert($data)->saveInsert();
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
        $table = container($this->table, true);
        $result = $table->where(['id' => $id])->update($data)->saveUpdate();
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
        $table = container($this->table, true);
        $result = $table->where([['id', 'in', $idArr]])->delete();
        if (empty($result)) {
            throwFailJson('999999');
        }
        throwSuccessJson();
    }
}
