<?php

declare(strict_types=1);

namespace app\module\service\auth;

use app\module\db\table\auth\AuthMenu as TableAuthMenu;
use app\module\logic\auth\AuthMenu as LogicAuthMenu;
use app\module\service\AbstractService;

class AuthMenu extends AbstractService
{
    /**
     * 获取树状权限菜单
     *
     * @param array $field
     * @param array $where
     * @return void
     */
    public function getTree(array $field = [], array $where = [])
    {
        /* if (!empty($field)) {
            //不为空时,补充一些必须的字段,否则生成树状菜单会报错
            $field = array_unique(array_merge($field, ['menuId', 'pid']));
        }*/
        $order = [
            'pid' => 'asc',
            'sort' => 'asc',
            'menuId' => 'asc'
        ];
        $tableAuthMenu = container(TableAuthMenu::class, true);
        $tableAuthMenu->where($where);
        $tableAuthMenu->field($field)->order($order);
        if ($tableAuthMenu->isJoin()) {
            $tableAuthMenu->group(['id']);
        }
        $list = $tableAuthMenu->getBuilder()->get()->toArray();

        $tree = container(LogicAuthMenu::class)->getTree($list);

        throwSuccessJson(['tree' => $tree]);
    }
}
