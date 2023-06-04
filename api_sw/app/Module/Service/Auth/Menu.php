<?php

declare(strict_types=1);

namespace App\Module\Service\Auth;

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
            $pInfo = $this->getDao()->parseField(['pidPath', 'level'])->parseFilter(['id' => $data['pid'], 'sceneId' => $data['sceneId']])->info();
            if (empty($pInfo)) {
                throwFailJson(29999998);
            }
        }
        $id = $this->getDao()->parseInsert($data)->insert();
        if (empty($id)) {
            throwFailJson();
        }
        if (!empty($data['pid'])) {
            $update['pidPath'] = $pInfo->pidPath . '-' . $id;
            $update['level'] = $pInfo->level + 1;
        } else {
            $update['pidPath'] = '0-' . $id;
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
        if (isset($data['pid'])) {
            $oldInfo = $this->getDao()->parseFilter($filter)->info();
            if ($data['pid'] == $oldInfo->menuId) { //父级不能是自身
                throwFailJson(29999997);
            }
            if ($data['pid'] == $oldInfo->pid) {
                unset($data['pid']);    //未修改则删除，更新后就不用处理$data['pid']
            } else {
                if ($data['pid'] > 0) {
                    $pInfo = $this->getDao()->parseField(['pidPath', 'level'])->parseFilter(['id' => $data['pid'], 'sceneId' => $data['sceneId'] ?? $oldInfo->sceneId])->info();
                    if (empty($pInfo)) {
                        throwFailJson(29999998);
                    }
                    if (in_array($oldInfo->menuId, explode('-',  $pInfo->pidPath))) {   //父级不能是自身的子孙级
                        throwFailJson(29999996);
                    }
                    $data['pidPath'] =  $pInfo->pidPath . '-' . $oldInfo->menuId;
                    $data['level'] = $pInfo->level + 1;
                } else {
                    $data['pidPath'] = '0-' . $oldInfo->menuId;
                    $data['level'] = 1;
                }
            }
        }
        $result = $this->getDao()->parseFilter($filter)->parseUpdate($data)->update();
        if (empty($result)) {
            throwFailJson();
        }
        //修改pid时，更新所有子孙级的pidPath和level
        if (isset($data['pid'])) {
            $this->getDao()->parseFilter([['pidPath', 'like', $oldInfo->pidPath . '%']])
                ->parseUpdate([
                    'pidPathOfChild' => [
                        'newVal' => $data['pidPath'],
                        'oldVal' => $oldInfo->pidPath
                    ],
                    'levelOfChild' => [
                        'newVal' => $data['level'],
                        'oldVal' => $oldInfo->level
                    ]
                ])
                ->update();
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
            throwFailJson(29999995);
        }
        $result = $this->getDao()->parseFilter($filter)->delete();
        if (empty($result)) {
            throwFailJson();
        }
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
