<?php

declare(strict_types=1);

namespace App\Crontab;

use App\Module\Db\Dao\Log\Request;
use Hyperf\DbConnection\Db;
/* use Hyperf\Di\Annotation\Inject;
use Psr\Container\ContainerInterface;

class LogRequest
{
    #[Inject]
    protected ContainerInterface $container; */

class LogRequest
{
    /* 
    //查看分区
    SELECT PARTITION_NAME, PARTITION_EXPRESSION, PARTITION_DESCRIPTION, TABLE_ROWS
    FROM INFORMATION_SCHEMA.PARTITIONS
    WHERE TABLE_SCHEMA = SCHEMA() AND TABLE_NAME = '表名';
    //修改表为分区表
    ALTER TABLE `表名` PARTITION BY RANGE (TO_DAYS( addTime ))
    (PARTITION `p20190204` VALUES LESS THAN (TO_DAYS('2019-02-05')) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
    PARTITION `pMax` VALUES LESS THAN (MAXVALUE) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 );
    //新增分区
    ALTER TABLE `表名` ADD PARTITION
    (PARTITION `p20190107` VALUES LESS THAN (TO_DAYS('2019-01-08')) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
    PARTITION `p20190108` VALUES LESS THAN (TO_DAYS('2019-01-09')) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 );
    //删除分区
    ALTER TABLE `表名` DROP PARTITION p20190107,p20190108;
    */
    /**
     * 请求日志表每周新增分区（无分区时建立分区）
     */
    public function partition()
    {
        $dao = getDao(Request::class);
        $connection = $dao->getConnection();
        $table = $dao->getTable();
        $db = Db::connection($connection);
        /**--------查询分区（不是分区表或无分区，查询结果都会有一项，且第一项内PARTITION_NAME值为null） 开始--------**/
        $partitionSelSql = 'SELECT PARTITION_NAME FROM INFORMATION_SCHEMA.PARTITIONS WHERE TABLE_SCHEMA = SCHEMA() AND TABLE_NAME = \'' . $table . '\'';
        $partitionResult = $db->select($partitionSelSql, [], false);
        $partitionList = [];
        if ($partitionResult[0]['PARTITION_NAME'] !== null) {
            foreach ($partitionResult as $v) {
                $partitionList[] = $v['PARTITION_NAME'];
            }
        }
        /**--------查询分区（不是分区表或无分区，查询结果都会有一项，且第一项内PARTITION_NAME值为null） 结束--------**/

        /**--------无分区则建立当天的分区 开始--------**/
        if (empty($partitionList)) {
            $partitionCreateSql = 'PARTITION `p' . date('Ymd') . '` VALUES LESS THAN (TO_DAYS( \'' . date('Y-m-d', time() + 24 * 60 * 60) . '\' )) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0';
            $partitionCreateSql = 'ALTER TABLE `' . $table . '` PARTITION BY RANGE (TO_DAYS( addTime )) ( ' . $partitionCreateSql . ' )';
            $db->update($partitionCreateSql);
        }
        /**--------无分区则建立当天的分区 结束--------**/

        /**--------检测后面7天的分区是否存在，没有则新增分区 开始--------**/
        $partitionAddSql = '';
        for ($i = 1; $i < 7; $i++) {
            if (!in_array('p' . date('Ymd', time() + $i * 24 * 60 * 60), $partitionList)) {
                $partitionAddSql .= 'PARTITION `p' . date('Ymd', time() + $i * 24 * 60 * 60) . '` VALUES LESS THAN (TO_DAYS( \'' . date('Y-m-d', time() + ($i + 1) * 24 * 60 * 60) . '\' )) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,';
            }
        }
        if (!in_array('p' . date('Ymd', time() + $i * 24 * 60 * 60), $partitionList)) {
            $partitionAddSql .= 'PARTITION `p' . date('Ymd', time() + $i * 24 * 60 * 60) . '` VALUES LESS THAN (TO_DAYS( \'' . date('Y-m-d', time() + ($i + 1) * 24 * 60 * 60) . '\' )) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0';
        }
        if ($partitionAddSql) {
            $partitionAddSql = 'ALTER TABLE `' . $table . '` ADD PARTITION ( ' . $partitionAddSql . ' )';
            $db->update($partitionAddSql);
        }
        /**--------检测后面7天的分区是否存在，没有则新增分区 结束--------**/
    }
}
