<?php

declare(strict_types=1);

namespace app\controller;

use app\module\db\table\system\SystemAdmin;
use support\Db;
use support\Request;

class Test extends AbstractController
{
    public function index(Request $request)
    {
        //\support\Container::get(SystemAdmin::class)->getTable();
        //Db::table(Db::raw('system_admin AS alias FORCE INDEX (account)'))->dump();
        //Db::table('system_admin', 'alias')->dump();
        
        /* $data = [
            'where' => ['id' => 1],
            'field' => ['*', 'roleName'],
            'order' => ['id' => 'asc']
        ];
        $tableSystemAdmin = container(SystemAdmin::class, true);
        $tableSystemAdmin->parseWhere($data['where']);
        if ($tableSystemAdmin->isJoin()) {
            $count = $tableSystemAdmin->getBuilder()->distinct()->count($tableSystemAdmin->getTableAlias() . '.' . $tableSystemAdmin->getPrimaryKey());
        } else {
            $count = $tableSystemAdmin->getBuilder()->count();
        }
        $list = [];
        if ($count > 0) {
            $tableSystemAdmin->parseField($data['field'])->parseOrder($data['order']);
            if ($tableSystemAdmin->isJoin()) {
                $tableSystemAdmin->parseGroup(['id']);
            }
            $list = $tableSystemAdmin->getBuilder()->offset(0)->limit(10)->get()->toArray();
        }
        throwSuccessJson(['count' => $count, 'list' => $list]); */

        /* $count = Db::select('select count(*) from `system_admin` as `system_admin` where `system_admin`.`adminId` = 1');
        $list = Db::select('select `system_admin`.*, `auth_role`.`roleName` from `system_admin` as `system_admin` left join `auth_role_rel_of_system_admin` as `auth_role_rel_of_system_admin` on `auth_role_rel_of_system_admin`.`adminId` = `system_admin`.`adminId` left join `auth_role` as `auth_role` on `auth_role`.`roleId` = `auth_role_rel_of_system_admin`.`roleId` where `system_admin`.`adminId` = 1 group by `system_admin`.`adminId` order by `system_admin`.`adminId` asc limit 10 offset 0');
        throwSuccessJson(['count' => $count, 'list' => $list]); */

        /*--------压测 开始--------*/
        //sleep(2); //这个方式压测，workman在这里会阻塞，不能处理新的请求。而swoole有协程切换，不会阻塞，所以这个方式压测swoole会更快
        /* for ($i = 0; $i < 100; $i++) {
            //Db::table('system_admin')->get()->toArray();
            container(SystemAdmin::class)->getBuilder()->get()->toArray();
        } */
        /*--------压测 结束--------*/

        /*--------测试状态污染 开始--------*/
        /* if ($request->input('test') == 1) {
            throwSuccessJson(['count' => '1', 'list' => []]);
        } else {
            throwSuccessJson(['count' => '2', 'list' => []]);
        } */
        /*--------测试状态污染 结束--------*/

        /*--------容器生成的对象代码提示测试 开始--------*/
        /* (new SystemAdmin())->getTable();
        container(SystemAdmin::class, true)->getTable();
        \support\Container::get(SystemAdmin::class)->getTable();
        \support\Container::get(SystemAdmin::class)->getTable();
        (new \DI\Container())->get(SystemAdmin::class)->getTable(); */
        /*--------容器生成的对象代码提示测试 结束--------*/
        return response('hello webman');
    }
}
