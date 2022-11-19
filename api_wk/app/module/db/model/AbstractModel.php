<?php

declare(strict_types=1);

namespace app\module\db\model;

use support\Db;
// use support\Model;

/**
 * 这个类从AbstractDao中把无状态的属性拆分出来，让该类对象可以缓存在容器中，不至于每次使用dao对象都要查询一次数据库（查询全部列）
 */
abstract class AbstractModel/*  extends Model */
{
    protected $connection;  //默认连接
    protected $table;   //默认表名
    protected $primaryKey;   //主键名

    protected array $allColumn = [];   //全部列（正常都是固定的，分库分表也都一样）

    final public function __construct(array $attributes = [])
    {
        $this->allColumn = Db::connection($this->connection)->getSchemaBuilder()->getColumnListing($this->table);
        //parent::__construct($attributes);
    }

    //模型不用的话，即不继承框架自带模型，需要定义以下几个方法。否则需要修改AbstractDao类中的部分方法
    final public function getConnectionName(): string|null
    {
        return $this->connection;
    }

    final public function getTable(): string
    {
        return $this->table;
    }

    final public function getKeyName(): string
    {
        return $this->primaryKey;
    }

    final public function getAllColumn(): array
    {
        return $this->allColumn;
    }

    /* public function __get(string $name)
    {
        return $this->{$name};
    } */
}
