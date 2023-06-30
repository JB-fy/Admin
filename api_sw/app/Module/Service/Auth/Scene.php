<?php

declare(strict_types=1);

namespace App\Module\Service\Auth;

use App\Module\Db\Dao\Auth\ActionRelToScene;
use App\Module\Db\Dao\Auth\Menu;
use App\Module\Db\Dao\Auth\Role;
use App\Module\Service\AbstractService;

class Scene extends AbstractService
{
    /**
     * 删除
     *
     * @param array $filter
     * @return void
     */
    public function delete(array $filter)
    {
        $idArr = $this->getIdArr($filter);
        $result = $this->getDao()->parseFilter($filter)->delete();
        if (empty($result)) {
            throwFailJson();
        }
        getDao(Menu::class)->parseFilter(['roleId' => $idArr])->delete();
        getDao(ActionRelToScene::class)->parseFilter(['roleId' => $idArr])->delete();
        getDao(Role::class)->parseFilter(['roleId' => $idArr])->delete();
        throwSuccessJson();
    }
}
