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
        $afterCount = ($offset == 0 && $limit == 0);  //用于判断是否先获取$list，再通过count($list)计算$count

        $tableAuthMenu = container(TableAuthMenu::class, true);
        $tableAuthMenu->where($where);
        if (!$afterCount) {
            if ($tableAuthMenu->isJoin()) {
                $count = $tableAuthMenu->getBuilder()->distinct()->count($tableAuthMenu->getTableAlias() . '.' . $tableAuthMenu->getPrimaryKey());
            } else {
                $count = $tableAuthMenu->getBuilder()->count();
            }
        }

        $list = [];
        if ($afterCount || $count > $offset) {
            $tableAuthMenu->field($field)->order($order);
            if ($tableAuthMenu->isJoin()) {
                $tableAuthMenu->group(['id']);
            }
            if ($afterCount) {
                $list = $tableAuthMenu->getBuilder()->get()->toArray();
                $count = count($list);
            } else {
                $list = $tableAuthMenu->getBuilder()->offset($offset)->limit($limit)->get()->toArray();
            }
        }

        $fieldAfter = ['showMenu']; //需得到列表后才能处理的字段
        $fieldAfter = array_intersect($fieldAfter, $field);
        if (!empty($list) && !empty($fieldAfter)) {
            foreach ($fieldAfter as $filedStr) {
                switch ($filedStr) {
                    case 'showMenu':
                        foreach ($list as &$v) {
                            $tmp = json_decode($v->extendData, true);
                            $v->title = $tmp['title'] ?? [];
                            $v->icon = $tmp['icon'] ?? '';
                            $v->url = $tmp['url'] ?? '';
                        }
                        break;
                }
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
        /* if (!empty($field)) {
            //不为空时,补充一些必须的字段,否则生成树状菜单会报错
            $field = array_unique(array_merge($field, ['menuId', 'pid']));
        }*/
        $order = [
            'pid' => 'asc',
            'sort' => 'asc',
            'menuId' => 'asc'
        ];
        try {
            $this->list($field, $where, $order, 0, 0);
        } catch (\app\exception\Json $e) {
            $list = json_decode($e->getMessage(), true)['data']['list'] ?? [];
        }

        $tree = [];
        if (!empty($list)) {
            $tree = container(LogicAuthMenu::class)->tree($list);
        }
        throwSuccessJson(['tree' => $tree]);
    }
}
