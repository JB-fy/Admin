<?php

declare(strict_types=1);

namespace App\Module\Service\Auth;

use App\Module\Logic\Auth\Menu as LogicAuthMenu;
use App\Module\Service\AbstractService;

class Menu extends AbstractService
{
    /**
     * 获取树状权限菜单
     *
     * @param array $field
     * @param array $where
     * @return void
     */
    public function tree(array $field = [], array $where = [])
    {
        $dao = getDao($this->daoClassName);
        $list = $dao->field($field)->where($where)->getList();

        $tree = $this->container->get(LogicAuthMenu::class)->tree($list);
        throwSuccessJson(['tree' => $tree]);
    }
}
