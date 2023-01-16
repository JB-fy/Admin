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
    //注意：如果表带有唯一索引（含主键），则用于建立分区的字段必须和唯一索引字段组成联合索引
    //查看分区
    SELECT PARTITION_NAME, PARTITION_EXPRESSION, PARTITION_DESCRIPTION, TABLE_ROWS
    FROM INFORMATION_SCHEMA.PARTITIONS
    WHERE TABLE_SCHEMA = SCHEMA() AND TABLE_NAME = '表名';
    //修改表为分区表
    ALTER TABLE `表名` PARTITION BY RANGE (TO_DAYS( 分区字段 ))
    (PARTITION `20220115` VALUES LESS THAN (TO_DAYS('2022-01-16')) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
    PARTITION `max` VALUES LESS THAN (MAXVALUE) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 );
    //新增分区
    ALTER TABLE `表名` ADD PARTITION
    (PARTITION `20220115` VALUES LESS THAN (TO_DAYS('2022-01-16')) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
    PARTITION `20220116` VALUES LESS THAN (TO_DAYS('2022-01-17')) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 );
    //删除分区
    ALTER TABLE `表名` DROP PARTITION 20220115,20220116;
    */
    /**
     * 请求日志表分区
     *
     * @return void
     */
    public function partition()
    {
        try {
            $dao = getDao(Request::class);
            $connection = $dao->getConnection();
            $table = $dao->getTable();
            $db = Db::connection($connection);
            $partitionField = 'createTime'; //分区字段，即根据该字段做分区
            //分区间隔时间，分区数量设置建议：两者总时长要比定时器间隔多几天时间，方便分区失败时，有时间让技术人工处理
            $partitionTime = 1 * 24 * 60 * 60;  //间隔多长时间创建一个分区
            $partitionNumber = 7;  //当前时间后面，需要新增的分区数量

            /**--------查询分区（不是分区表或无分区，查询结果都会有一项，且第一项内PARTITION_NAME值为null） 开始--------**/
            $partitionSelSql = 'SELECT PARTITION_NAME FROM INFORMATION_SCHEMA.PARTITIONS WHERE TABLE_SCHEMA = SCHEMA() AND TABLE_NAME = \'' . $table . '\'';
            $partitionResult = $db->select($partitionSelSql, [], false);
            $partitionList = [];
            if ($partitionResult[0]->PARTITION_NAME !== null) {
                foreach ($partitionResult as $v) {
                    $partitionList[] = $v->PARTITION_NAME;
                }
            }
            /**--------查询分区（不是分区表或无分区，查询结果都会有一项，且第一项内PARTITION_NAME值为null） 结束--------**/

            /**--------无分区则建立当前分区 开始--------**/
            if (empty($partitionList)) {
                $ltTime = time() + $partitionTime;
                $ltDate = date('Y-m-d', $ltTime);
                $partitionName = date('Ymd', $ltTime - 24 * 60 * 60); //该分区的最大日期作为分区名称
                $partitionCreateSql = 'PARTITION `' . $partitionName . '` VALUES LESS THAN (TO_DAYS( \'' . $ltDate . '\' )) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0';
                $partitionCreateSql = 'ALTER TABLE `' . $table . '` PARTITION BY RANGE (TO_DAYS( ' . $partitionField . ' )) ( ' . $partitionCreateSql . ' )';
                $db->update($partitionCreateSql);
                $maxPartitionName = $partitionName;
            } else {
                $maxPartitionName = $partitionList[count($partitionList) - 1];
            }
            /**--------无分区则建立当前分区 结束--------**/

            /**--------检测需要创建的分区是否存在，没有则新增分区 开始--------**/
            $partitionAddSqlList = [];
            for ($i = 1; $i <= $partitionNumber; $i++) {
                //时间超过最大的分区后才开始新增分区，且以最大分区的时间开始往后计算
                if (date('Ymd', time() + ($i + 1) * $partitionTime - 24 * 60 * 60) > $maxPartitionName) {
                    $ltTime = strtotime($maxPartitionName) + ($i + 1) * $partitionTime;
                    $ltDate = date('Y-m-d', $ltTime);
                    $partitionName = date('Ymd', $ltTime - 24 * 60 * 60); //该分区的最大日期作为分区名称
                    $partitionAddSqlList[] = 'PARTITION `' . $partitionName . '` VALUES LESS THAN (TO_DAYS( \'' . $ltDate . '\' )) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0';
                }
            }
            if (!empty($partitionAddSqlList)) {
                $partitionAddSql = implode(',', $partitionAddSqlList);
                $partitionAddSql = 'ALTER TABLE `' . $table . '` ADD PARTITION ( ' . $partitionAddSql . ' )';
                $db->update($partitionAddSql);
            }
            /**--------检测需要创建的分区是否存在，没有则新增分区 结束--------**/
        } catch (\Throwable $th) {
            //出错时，做记录通知后台管理人员，让其联系技术人工处理
            var_dump($th->getMessage());
        }
    }
}
