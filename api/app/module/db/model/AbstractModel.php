<?php

declare(strict_types=1);

namespace app\module\db\model;

use support\Model;
use support\Db;

/**
 * 这个类从AbstractDao中把无状态的属性拆分出来，让该类对象可以缓存在容器中，不至于每次使用dao对象都要查询一次数据库（查询全部列）
 */
abstract class AbstractModel extends Model
{
    protected $connection;  //默认连接
    protected $table;   //默认表名
    protected $primaryKey;   //主键名

    protected bool $isSoftDelete = false;   //是否软删除
    protected string $fieldSoftDelete = 'isDelete';   //软删除字段。当isSoftDelete为true时才有效

    public array $allColumn = [];   //全部列（正常都是固定的，分库分表也都一样）

    final public function __construct(array $attributes = [])
    {
        $this->allColumn = Db::connection($this->connection)->getSchemaBuilder()->getColumnListing($this->table);
        parent::__construct($attributes);
    }

    /* public function __get(string $name)
    {
        return $this->{$name};
    } */
}
