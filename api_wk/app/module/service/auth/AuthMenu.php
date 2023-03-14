<?php

declare(strict_types=1);

namespace app\module\service\auth;

use app\module\logic\auth\AuthMenu as LogicAuthMenu;
use app\module\service\AbstractService;

class AuthMenu extends AbstractService
{
    protected $dao = \app\module\db\dao\auth\AuthMenu::class;

    /**
     * 获取树状权限菜单
     *
     * @param array $field
     * @param array $where
     * @return void
     */
    public function tree(array $field = [], array $where = [])
    {
        $daoAuthMenu = container($this->dao, true);
        $list = $daoAuthMenu->field($field)->where($where)->getList();

        $tree = container(LogicAuthMenu::class)->tree($list);
        throwSuccessJson(['tree' => $tree]);
    }
}
