<?php

declare(strict_types=1);

namespace App\Module\Service\Platform;

use App\Module\Service\AbstractService;

class Config extends AbstractService
{
    /**
     * 获取
     *
     * @param array $where
     * @return void
     */
    public function get(array $where = [])
    {
        $config = $this->getDao()->filter($where)->getBuilder()->pluck('configValue', 'configKey')->toArray();
        throwSuccessJson(['config' => $config]);
    }

    /**
     * 保存
     *
     * @param array $data
     * @return void
     */
    public function save(array $data)
    {
        $dao = $this->getDao();
        foreach ($data as $k => $v) {
            $dao->getBuilder()->updateOrInsert(['configKey' => $k], ['configValue' => $v]);
        }
        /* Db::beginTransaction();
        try {
            $dao = $this->getDao();
            foreach ($data as $k => $v) {
                $dao->getBuilder()->updateOrInsert(['configKey' => $k], ['configValue' => $v]);
            }
            Db::commit();
        } catch (\Throwable $e) {
            Db::rollBack();
            throwFailJson();
        } */
        throwSuccessJson();
    }
}
