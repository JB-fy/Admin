<?php

declare(strict_types=1);

namespace App\Module\Db\Model;

use Hyperf\DbConnection\Db;

/**
 * 这个类从AbstractDao中把无状态的属性拆分出来，让该类对象可以缓存在容器中，不至于每次使用dao对象都要查询一次数据库（查询全部列）
 * 
 * 强烈建议：不继承\Hyperf\DbConnection\Model\Model类。
 *      如果继承，防止改变其状态，否则会对app\Module\Db\Dao的使用造成影响。
 * 使用建议：
 *      禁止在任何地方使用app\Module\Db\Model类，防止改变其状态
 *      只使用app\Module\Db\Dao或Hyperf\DbConnection\Db做数据库处理
 * 
 * 需使用框架命令快速生成模型时，临时改为继承模型
 */
abstract class AbstractModel/*  extends \Hyperf\DbConnection\Model\Model */
{
    protected ?string $connection = 'default';  //默认连接
    protected ?string $table;   //默认表名
    protected string $primaryKey = 'id';   //主键名
    //protected $keyType = 'int';   //主键不是整数时，设置为string
    //public $incrementing = true;  //主键非递增或非数字时，设置为false
    public bool $timestamps = false; //默认true，即默认表中存在created_at和updated_at两个列，且会自动管理

    protected array $allColumn = [];   //全部列（正常都是固定的，分库分表也都一样）

    final public function __construct(array $attributes = [])
    {
        $this->allColumn = Db::connection($this->connection)->getSchemaBuilder()->getColumnListing($this->table);
        //parent::__construct($attributes);
    }

    final public function getAllColumn(): array
    {
        return $this->allColumn;
    }

    //不继承\Hyperf\DbConnection\Model\Model类，需要定义以下几个方法。否则需要修改AbstractDao类中的部分方法
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
}
