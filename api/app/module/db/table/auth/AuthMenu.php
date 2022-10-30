<?php

declare(strict_types=1);

namespace app\module\db\table\auth;

use app\module\db\table\AbstractTable;
use DI\Annotation\Inject;

class AuthMenu extends AbstractTable
{
    /**
     * @Inject
     * @var \app\module\db\model\auth\AuthMenu
     */
    protected $model;

    /**
     * 解析field（独有的）
     *
     * @param string $key
     * @return boolean
     */
    protected function fieldOfAlone(string $key): bool
    {
        switch ($key) {
            case 'menuTree':    //树状需要以下字段和排序方式
                $this->field['select'][] = 'menuId';
                $this->field['select'][] = 'pid';

                $this->order[] = ['method' => 'orderBy', 'param' => ['pid', 'asc']];
                $this->order[] = ['method' => 'orderBy', 'param' => ['sort', 'asc']];
                $this->order[] = ['method' => 'orderBy', 'param' => ['menuId', 'asc']];
                return true;
            case 'showMenu':    //前端显示菜单需要以下字段，且title需要转换
                $this->fieldAfter[] = 'showMenu';   //需做后续处理

                $this->field['select'][] = 'menuName';
                $this->field['select'][] = 'extendData->title AS title';
                //$this->field['select'][] = Db::raw('JSON_UNQUOTE(JSON_EXTRACT(extendData, "$.title")) AS title'); //不知道怎么直接转成对象返回
                $this->field['select'][] = 'extendData->url AS url';
                $this->field['select'][] = 'extendData->icon AS icon';
                return true;
        }
        return false;
    }

    /**
     * 获取数据库数据后，再做处理的字段
     *
     * @param object $info
     * @return object
     */
    public function handleFieldAfter(object $info): object
    {
        foreach ($this->fieldAfter as $field) {
            switch ($field) {
                case 'showMenu':
                    $info->title = $info->title ? json_decode($info->title, true) : [];
                    $info->icon = $info->icon ?? '';
                    $info->url = $info->url ?? '';
                    break;
            }
        }
        return $info;
    }
}
