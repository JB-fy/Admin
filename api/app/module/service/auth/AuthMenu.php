<?php

declare(strict_types=1);

namespace app\module\service\auth;

use app\module\db\table\auth\AuthMenu as TableAuthMenu;
use app\module\logic\auth\AuthMenu as LogicAuthMenu;
use app\module\service\AbstractService;

class AuthMenu extends AbstractService
{
    /**
     * 获取列表
     */
    public function list(array $field = [], array $where = [], array $order = [], int $offset = 1, int $limit = 10)
    {
        $countAfter = ($offset == 0 && $limit == 0);  //用于判断是否先获取$list，再通过count($list)计算$count

        $tableAuthMenu = container(TableAuthMenu::class, true);
        $tableAuthMenu->where($where);
        if (!$countAfter) {
            if ($tableAuthMenu->isJoin()) {
                $count = $tableAuthMenu->getBuilder()->distinct()->count($tableAuthMenu->getTable() . '.' . $tableAuthMenu->getPrimaryKey());
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

    /**
     * 获取树状权限菜单
     *
     * @param array $field
     * @param array $where
     * @return void
     */
    public function tree(array $field = [], array $where = [])
    {
        $tableAuthMenu = container(TableAuthMenu::class, true);
        $list = $tableAuthMenu->field($field)->where($where)->getList();

        $tree = container(LogicAuthMenu::class)->tree($list);
        throwSuccessJson(['tree' => $tree]);
    }
}
