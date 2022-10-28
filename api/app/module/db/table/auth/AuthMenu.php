<?php

declare(strict_types=1);

namespace app\module\db\table\auth;

use app\module\db\table\AbstractTable;
use DI\Annotation\Inject;
use support\Db;

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
            case 'showMenu':
                $this->field['select'][] = 'extendData';
                /* $this->field['select'][] = 'extendData->title AS title';
                //$this->field['select'][] = Db::raw('JSON_UNQUOTE(JSON_EXTRACT(extendData, "$.title")) AS title'); //不知道怎么直接转成对象返回
                $this->field['select'][] = 'extendData->url AS url';
                $this->field['select'][] = 'extendData->icon AS icon'; */
                return true;
        }
        return false;
    }
}
