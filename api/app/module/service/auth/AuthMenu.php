<?php

declare(strict_types=1);

namespace app\module\service\auth;

use app\module\logic\auth\AuthMenu as LogicAuthMenu;
use app\module\service\AbstractService;

class AuthMenu extends AbstractService
{
    protected $table = \app\module\db\table\auth\AuthMenu::class;

    /**
     * 获取树状权限菜单
     *
     * @param array $field
     * @param array $where
     * @return void
     */
    public function tree(array $field = [], array $where = [])
    {
        $tableAuthMenu = container($this->table, true);
        $list = $tableAuthMenu->field($field)->where($where)->getList();

        $tree = container(LogicAuthMenu::class)->tree($list);
        throwSuccessJson(['tree' => $tree]);
    }
}
