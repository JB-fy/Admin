<?php

declare(strict_types=1);

namespace App\Module\Db\Model;

abstract class AbstractModel extends \Hyperf\DbConnection\Model\Model
{
    protected ?string $connection = 'default';  //默认连接
    protected ?string $table;   //默认表名
    protected string $primaryKey = 'id';   //主键名
    //protected $keyType = 'int';   //主键不是整数时，设置为string
    //public $incrementing = true;  //主键非递增或非数字时，设置为false
    public bool $timestamps = false; //默认true，即默认表中存在created_at和updated_at两个列，且会自动管理

    final public function getAllColumn(): array
    {
        //return \Hyperf\DbConnection\Db::connection($this->connection)->getSchemaBuilder()->getColumnListing($this->table);
        return $this->fillable; //必须禁止修改自动生成的模型的fillable属性
    }
}
