<?php

declare(strict_types=1);

namespace App\Controller\Log;

use App\Controller\AbstractController;
use App\Module\Db\Dao\Log\Request as LogRequest;

class Request extends AbstractController
{
    /**
     * 列表
     *
     * @return void
     */
    public function list()
    {
        $sceneCode = $this->scene->getCurrentSceneCode();
        switch ($sceneCode) {
            case 'platform':
                $data = $this->validate(__FUNCTION__, $sceneCode);
                $this->checkAuth(__FUNCTION__, $sceneCode);

                /**--------参数处理 开始--------**/
                $allowField = $this->getAllowField(LogRequest::class);
                if (empty($data['field'])) {
                    $data['field'] = $allowField;
                } else {
                    $data['field'] = array_intersect($data['field'], $allowField);
                    if (empty($data['field'])) {
                        $data['field'] = $allowField;
                    }
                }
                /**--------参数处理 结束--------**/

                $this->service->listWithCount(...$data);
                break;
        }
    }
}
