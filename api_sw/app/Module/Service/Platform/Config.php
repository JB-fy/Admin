<?php

declare(strict_types=1);

namespace App\Module\Service\Platform;

use App\Module\Service\AbstractService;

class Config extends AbstractService
{
    /**
     * 获取
     *
     * @param array $filter
     * @return void
     */
    public function get(array $filter = [])
    {
        $config = $this->getDao()->parseFilter($filter)->getBuilder()->pluck('configValue', 'configKey')->toArray();
        foreach ($config as $k => $v) {
            switch ($k) {
                case 'hotSearch':
                    $config[$k] = json_decode($v, true);
                    break;
            }
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
        foreach ($data as $k => $v) {
            switch ($k) {
                case 'hotSearch':
                    $this->getDao()->getBuilder()->updateOrInsert(['configKey' => $k], ['configValue' => json_encode($v, JSON_UNESCAPED_UNICODE)]);
                    break;
                default:
                    $this->getDao()->getBuilder()->updateOrInsert(['configKey' => $k], ['configValue' => $v]);
                    break;
            }
        }
        /* Db::beginTransaction();
        try {
            foreach ($data as $k => $v) {
                switch ($k) {
                    case 'hotSearch':
                        $this->getDao()->getBuilder()->updateOrInsert(['configKey' => $k], ['configValue' => json_encode($v, JSON_UNESCAPED_UNICODE)]);
                        break;
                    default:
                        $this->getDao()->getBuilder()->updateOrInsert(['configKey' => $k], ['configValue' => $v]);
                        break;
                }
            }
            Db::commit();
        } catch (\Throwable $e) {
            Db::rollBack();
            throwFailJson();
        } */
        throwSuccessJson();
    }
}
