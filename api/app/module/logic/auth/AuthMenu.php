<?php

declare(strict_types=1);

namespace app\module\logic\auth;

use app\module\logic\AbstractLogic;

class AuthMenu extends AbstractLogic
{
    /**
     * 取出数组$list中$menuId的子孙树
     *
     * @param array $list
     * @param int $menuId
     * @return array
     */
    public function getTree(array $list, int $menuId = 0): array
    {
        $tree = [];
        foreach ($list as $k => $v) {
            unset($list[$k]);
            if ($v->pid == $menuId) {
                $v->children = $this->getTree($list, $v->menuId);
                $tree[] = $v;
            }
        }
        return $tree;
    }
}
