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
    protected $table = '';   //table类的路径，调用地方实例化对象。因table类带有状态，无法直接使用容器注释功能作依赖注入。

    /**
     * 获取列表（通用，如果某些类需要特殊处理，可重新定义）
     */
    public function list(array $field = [], array $where = [], array $order = [], int $offset = 1, int $limit = 10)
    {
        $countAfter = ($offset == 0 && $limit == 0);  //用于判断是否先获取$list，再通过count($list)计算$count
        $tableAuthMenu = container($this->table, true);
        $tableAuthMenu->where($where);
        if (!$countAfter) {
            if ($tableAuthMenu->isJoin()) {
                $count = $tableAuthMenu->getBuilder()->distinct()->count($tableAuthMenu->getTable() . '.' . $tableAuthMenu->getKey());
            } else {
                $count = $tableAuthMenu->getBuilder()->count();
            }
        }

        $list = [];
        if ($countAfter || $count > $offset) {
            $tableAuthMenu->field($field)->order($order);
            if ($tableAuthMenu->isJoin()) {
                $tableAuthMenu->group(['id']);
            }
            $list = $tableAuthMenu->getList($offset, $limit);
            if ($countAfter) {
                $count = count($list);
            }
        }

        throwSuccessJson(['count' => $count, 'list' => $list]);
    }
}
