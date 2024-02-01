<?php

declare(strict_types=1);

namespace App\Module\Service\Auth;

use App\Module\Db\Dao\Auth\RoleRelToMenu;
use App\Module\Logic\Auth\Menu as LogicAuthMenu;
use App\Module\Service\AbstractService;

class Menu extends AbstractService
{
    /**
     * 创建
     *
     * @param array $data
     * @return void
     */
    public function create(array $data)
    {
        if (!empty($data['pid'])) {
            $pInfo = $this->getDao()->parseField(['idPath', 'level'])->parseFilter(['id' => $data['pid'], 'sceneId' => $data['sceneId']])->info();
            if (empty($pInfo)) {
                throwFailJson(29999997);
            }
        }
        $id = $this->getDao()->parseInsert($data)->insert();
        if (empty($id)) {
            throwFailJson();
        }
        if (!empty($data['pid'])) {
            $update['idPath'] = $pInfo->idPath . '-' . $id;
            $update['level'] = $pInfo->level + 1;
        } else {
            $update['idPath'] = '0-' . $id;
            $update['level'] = 1;
        }
        $this->getDao()->parseFilter(['id' => $id])->parseUpdate($update)->update();
        throwSuccessJson();
    }

    /**
     * 更新
     *
     * @param array $data
     * @param array $filter
     * @return void
     */
    public function update(array $data, array $filter)
    {
        if (isset($data['pid'])) { //存在pid则只能一个个循环更新
            $idArr = $this->getIdArr($filter);
            foreach ($idArr as $id) {
                $filterOne = ['id' => $id];
                $oldInfo = $this->getDao()->parseFilter($filterOne)->info();
                if ($data['pid'] == $oldInfo->menuId) { //父级不能是自身
                    throwFailJson(29999996);
                }
                if ($data['pid'] != $oldInfo->pid) {
                    if ($data['pid'] > 0) {
                        $pInfo = $this->getDao()->parseField(['idPath', 'level'])->parseFilter(['id' => $data['pid'], 'sceneId' => $data['sceneId'] ?? $oldInfo->sceneId])->info();
                        if (empty($pInfo)) {
                            throwFailJson(29999997);
                        }
                        if (in_array($oldInfo->menuId, explode('-',  $pInfo->idPath))) {   //父级不能是自身的子孙级
                            throwFailJson(29999995);
                        }
                        $data['idPath'] =  $pInfo->idPath . '-' . $oldInfo->menuId;
                        $data['level'] = $pInfo->level + 1;
                    } else {
                        $data['idPath'] = '0-' . $oldInfo->menuId;
                        $data['level'] = 1;
                    }
                    //修改pid时，更新所有子孙级的idPath和level
                    $this->getDao()->parseFilter([['idPath', 'like', $oldInfo->idPath . '%']])
                        ->parseUpdate([
                            'idPathOfChild' => [
                                'newVal' => $data['idPath'],
                                'oldVal' => $oldInfo->idPath
                            ],
                            'levelOfChild' => [
                                'newVal' => $data['level'],
                                'oldVal' => $oldInfo->level
                            ]
                        ])
                        ->update();
                }
                $this->getDao()->parseFilter($filterOne)->parseUpdate($data)->update();
            }
        } else {
            $result = $this->getDao()->parseFilter($filter)->parseUpdate($data)->update();
            if (empty($result)) {
                throwFailJson();
            }
        }
        throwSuccessJson();
    }

    /**
     * 删除
     *
     * @param array $filter
     * @return void
     */
    public function delete(array $filter)
    {
        $idArr = $this->getIdArr($filter);
        if ($this->getDao()->parseFilter(['pid' => $idArr])->getBuilder()->exists()) {
            throwFailJson(29999994);
        }
        $result = $this->getDao()->parseFilter($filter)->delete();
        if (empty($result)) {
            throwFailJson();
        }
        getDao(RoleRelToMenu::class)->parseFilter(['menuId' => $idArr])->delete();
        throwSuccessJson();
    }

    /**
     * 获取树状权限菜单
     *
     * @param array $filter
     * @param array $field
     * @return void
     */
    public function tree(array $filter = [], array $field = [])
    {
        $list = $this->getDao()->parseField($field)->parseFilter($filter)->list();

        $tree = $this->container->get(LogicAuthMenu::class)->tree($list);
        throwSuccessJson(['tree' => $tree]);
    }
}
