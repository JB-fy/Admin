<?php

declare(strict_types=1);

namespace app\module\db\model;

use support\Db;

/**
 * 这个类从AbstractTable中把无状态的属性拆分出来，让该类对象可以缓存在容器中，不至于每次使用table对象都要查询一次数据库（查询全部列）
 */
abstract class AbstractModel
{
    protected string $connection = '';  //默认连接
    protected string $table;   //默认表名
    protected string $tableAlias = '';   //表别名。分表情况下，表名不固定。强制使用别名用于解析数据，再生成的SQL语句，通用性好
    protected string $primaryKey = '';   //主键名

    protected array $allColumn = [];   //全部列（正常都是固定的，分库分表也都一样）

    final public function __construct()
    {
        if (empty($this->tableAlias)) {
            $this->tableAlias = $this->table;   //默认设置成默认表名
        }
        $this->allColumn = Db::connection($this->connection)->getSchemaBuilder()->getColumnListing($this->table);
    }

    public function __get(string $name)
    {
        return $this->{$name};
    }
}

/* //继承原有的模型
use support\Model;

abstract class AbstractModel extends Model
{
    protected string $connection = '';  //默认连接
    protected string $table;   //默认表名
    protected string $alias;   //表别名。分表情况下，表名不固定。强制使用别名用于解析数据，再生成的SQL语句，通用性好
    protected string $primaryKey = '';   //主键名

    protected array $column = [];   //全部列（正常都是固定的，分库分表也都一样）

    final public function __construct()
    {
        $this->column = Db::connection($this->connection)->getSchemaBuilder()->getColumnListing($this->table);
        parent::__construct();
    }
} */
