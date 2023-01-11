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
        $config = $this->getDao()->where($where)->getBuilder()->pluck('configValue', 'configKey')->toArray();
        if (empty($config)) {
            throwFailJson('29999999');
        }
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
        $builder = $this->getDao()->getBuilder();
        foreach ($data as $k => $v) {
            $builder->updateOrInsert(['configKey' => $k], ['configValue' => $v]);
        }
        /* Db::beginTransaction();
        try {
            $builder = $this->getDao()->getBuilder();
            foreach ($data as $k => $v) {
                $builder->updateOrInsert(['configKey' => $k], ['configValue' => $v]);
            }
            Db::commit();
        } catch (\Throwable $e) {
            Db::rollBack();
            throwFailJson();
        } */
        throwSuccessJson();
    }
}
