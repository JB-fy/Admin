<?php

declare(strict_types=1);

namespace App\Module\Logic\Auth;

use App\Module\Logic\AbstractLogic;

class Menu extends AbstractLogic
{
    /**
     * 取出数组$list中$menuId的子孙树
     *
     * @param array $list
     * @param int $menuId
     * @return array
     */
    public function tree(array $list, int $menuId = 0): array
    {
        $tree = [];
        foreach ($list as $k => $v) {
            unset($list[$k]);
            if ($v->pid == $menuId) {
                $v->children = $this->tree($list, $v->menuId);
                $tree[] = $v;
            }
            /* if ($v['pid'] == $menuId) {
                $v['children'] = $this->tree($list, $v['menuId']);
                $tree[] = $v;
            } */
        }
        return $tree;
    }
}
