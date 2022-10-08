<?php

declare(strict_types=1);

namespace app\module\db\table;

use Illuminate\Database\Query\Builder;
use support\Db;

abstract class AbstractTable
{
    protected string $connection = '';  //分库情况下，解析后所确定的连接
    protected string $table = '';   //分表情况下，解析后所确定的表
    protected string $tableRaw = '';    //表的原生表达式。当需要强制索引等特殊情况时使用。示例：Db::raw('table AS alias FORCE INDEX (索引)')。
    protected array $field = [];   //解析后的field。格式：['select' => [字段,...], 'selectRaw' => [[字段（填selectRaw方法所需参数）],...]]。无顺序要求
    protected array $where = [];   //解析后的where。格式：[['method'=>'where', 'param'=>[参数]],...]。有顺序要求，必须与原来一致，改变顺序可能造成索引使用不同
    protected array $group = [];   //解析后的group。格式：[['method'=>'groupBy', 'param'=>[参数]],...]。有顺序要求，必须与原来一致，改变顺序会造成分组结果不同
    protected array $having = [];  //解析后的having。格式：['having'=>[[条件],...], 'havingRaw'=>[[条件],...]]。无顺序要求
    protected array $order = [];   //解析后的order。格式：[['method'=>'orderBy', 'param'=>[参数]],...]。有顺序要求，必须与原来一致，改变顺序会造成排序结果不同
    protected array $join = [];    //解析后的join。格式：[['method'=>'join', 'param'=>[参数]],...]。无顺序要求

    protected $model;   //模型
    protected $builder; //构造器

    /* protected string $modelClass = '';   //创建模型需要的类
    final public function __construct(array $tableSelectData = [], array $connectionSelectData = [])
    {
        $this->model = container($this->modelClass);
        $this->parseConnection($connectionSelectData)
            ->parseTable($tableSelectData);
    } */

    /**
     * 获取连接
     *
     * @return string
     */
    public function getConnection(): string
    {
        return empty($this->connection) ? $this->model->connection : $this->connection;
    }

    /**
     * 获取表名
     *
     * @return string
     */
    public function getTable(): string
    {
        return empty($this->table) ? $this->model->table : $this->table;
    }

    /**
     * 获取表别名
     *
     * @return string
     */
    final public function getTableAlias(): string
    {
        return $this->model->tableAlias;
    }

    /**
     * 获取主键名
     *
     * @return string
     */
    final public function getPrimaryKey(): string
    {
        return $this->model->primaryKey;
    }

    /**
     * 获取全部列
     *
     * @return array
     */
    final public function getAllColumn(): array
    {
        return $this->model->allColumn;
    }


    /**
     * 解析连接
     *
     * @param array $connectionSelectData   分库情况下，用于确定使用哪个连接
     * @return self
     */
    public function parseConnection(array $connectionSelectData = []): self
    {
        //选择逻辑
        //$this->connection = ''; //设置当前使用的连接
        return $this;
    }

    /**
     * 解析表名
     *
     * @param array $tableSelectData    分表情况下，用于确定使用哪个表
     * @return self
     */
    public function parseTable(array $tableSelectData = []): self
    {
        //选择逻辑
        //$this->table = ''; //设置当前使用的表名
        return $this;
    }

    /**
     * 解析表的原生表达式
     *
     * @param string $tableRaw  表的原生表达式。当需要强制索引等特殊情况时使用。示例：'__TABLE__ FORCE INDEX (索引)')。
     * @return self
     */
    public function parseTableRaw(string $tableRaw = ''): self
    {
        if (!empty($tableRaw)) {
            if (strpos($tableRaw, '__TABLE__') !== false) {
                $tableRaw = str_replace('__TABLE__', $this->getTable() . ' AS ' . $this->getTableAlias(), $tableRaw);
            }
            $this->tableRaw = Db::raw($tableRaw);
        }
        return $this;
    }

    /**
     * 解析field
     *
     * @param array $field  格式：['字段',...]
     * @return self
     */
    public function parseField(array $field): self
    {
        foreach ($field as $v) {
            //$this->parseJoin($v);   //父类会处理默认字段，不会联表。子类如需对某些字段特殊处理会重新定义该方法，但对默认字段的处理调用父类这个方法。
            switch ($v) {
                case 'id':
                    $this->field['select'][] = $this->getTableAlias() . '.' . $this->getPrimaryKey();
                    break;
                default:
                    $this->field['select'][] = $this->getTableAlias() . '.' . $v;
                    //$this->field['selectRaw'][] = ['IFNULL(字段名, \'\') AS ' . $v];
                    break;
            }
        }
        return $this;
    }

    /**
     * 解析where
     *
     * @param array $where  格式：['字段' => '值',...]
     * @return self
     */
    public function parseWhere(array $where): self
    {
        foreach ($where as $k => $v) {
            //$this->parseJoin($k, $v);   //父类会处理默认字段，不会联表。子类如需对某些字段特殊处理会重新定义该方法，但对默认字段的处理调用父类这个方法。
            switch ($k) {
                case 'id':
                    $this->where[] = ['method' => 'where', 'param' => [$this->getTableAlias() . '.' . $this->getPrimaryKey(), '=', $v]];
                    break;
                default:
                    $this->where[] = ['method' => 'where', 'param' => [$this->getTableAlias() . '.' . $k, '=', $v]];
                    //$this->where[] = ['method' => 'whereRaw', 'param' => ['age > :age', ['age' => $v], 'and']];
                    break;
            }
        }
        return $this;
    }

    /**
     * 解析group
     *
     * @param array $group  格式：['字段',...]
     * @return self
     */
    public function parseGroup(array $group): self
    {
        foreach ($group as $v) {
            switch ($v) {
                case 'id':
                    $this->group[] = ['method' => 'groupBy', 'param' => [$this->getTableAlias() . '.' . $this->getPrimaryKey()]];
                    break;
                default:
                    $this->group[] = ['method' => 'groupBy', 'param' => [$this->getTableAlias() . '.' . $v]];
                    //$this->group[] = ['method'=>'groupByRaw', 'param'=>[':xxxx', ['xxxx' => 'xxxx']]];
                    break;
            }
        }
        return $this;
    }

    /**
     * 解析having
     *
     * @param array $having 格式：['字段' => '值',...]
     * @return self
     */
    public function parseHaving(array $having): self
    {
        foreach ($having as $k => $v) {
            switch ($k) {
                case 'id':
                    $this->having['having'][] = [$this->getTableAlias() . '.' . $this->getPrimaryKey(), '=', $v];
                    break;
                default:
                    $this->having['having'][] = [$k, '=', $v];
                    //$this->having['havingRaw'][] = ['age > :age', ['age' => $v], 'and'];
                    break;
            }
        }
        return $this;
    }

    /**
     * 解析order
     *
     * @param array $order  格式：['字段' => 'asc或desc',...]
     * @return self
     */
    public function parseOrder(array $order): self
    {
        foreach ($order as $k => $v) {
            //$this->parseJoin($k);   //父类会处理默认字段，不会联表。子类如需对某些字段特殊处理会重新定义该方法，但对默认字段的处理调用父类这个方法。
            switch ($k) {
                case 'id':
                    $this->order[] = ['method' => 'orderBy', 'param' => [$this->getTableAlias() . '.' . $this->getPrimaryKey(), $v]];
                    break;
                default:
                    $this->order[] = ['method' => 'orderBy', 'param' => [$this->getTableAlias() . '.' . $k, $v]];
                    //$this->order[] = ['method'=>'orderByRaw', 'param'=>[':time ' . $v, ['time' => time()]]];
                    break;
            }
        }
        return $this;
    }

    /**
     * 解析join
     *
     * @param string $key   键，用于确定关联表
     * @param [type] $value 值，用于确定关联表
     * @return void
     */
    public function parseJoin(string $key, $value = null)
    {
    }

    /**
     * 判断是否联表
     *
     * @return boolean
     */
    public function isJoin(): bool
    {
        return !empty($this->join);
    }

    /**
     * 获取Db构造器
     *
     * @return Builder
     */
    final public function getBuilder(): Builder
    {
        if (empty($this->tableRaw)) {
            $this->builder = Db::table($this->getTable(), $this->getTableAlias(), $this->getConnection());
        } else {
            $this->builder = Db::table($this->tableRaw, null, $this->getConnection());
        }
        if (!empty($this->field)) {
            $this->handleField();
        }
        if (!empty($this->where)) {
            $this->handleWhere();
        }
        if (!empty($this->group)) {
            $this->handleGroup();
        }
        if (!empty($this->having)) {
            $this->handleHaving();
        }
        if (!empty($this->order)) {
            $this->handleOrder();
        }
        if (!empty($this->join)) {
            $this->handleJoin();
        }
        return $this->builder;
    }

    /**
     * 处理field
     * 
     * @return self
     */
    final protected function handleField(): self
    {
        foreach ($this->field as $k => $v) {
            switch ($k) {
                case 'select':
                    //$this->builder->{$k}(...$v);
                    $this->builder->{$k}($v);
                    break;
                case 'selectRaw':
                    foreach ($v as $v1) {
                        $this->builder->{$k}(...$v1);
                    }
                    break;
            }
        }
        return $this;
    }

    /**
     * 处理where
     *
     * @return self
     */
    final protected function handleWhere(): self
    {
        foreach ($this->where as $v) {
            $this->builder->{$v['method']}(...$v['param']);
        }
        return $this;
    }

    /**
     * 处理group
     *
     * @return self
     */
    final protected function handleGroup(): self
    {
        foreach ($this->group as $v) {
            $this->builder->{$v['method']}(...$v['param']);
        }
        return $this;
    }

    /**
     * 处理having
     *
     * @return self
     */
    final protected function handleHaving(): self
    {
        foreach ($this->having as $k => $v) {
            foreach ($v as $v1) {
                $this->builder->{$k}(...$v1);
            }
            /* switch ($k) {
                case 'having':
                    $this->builder->{$k}($v);
                    break;
                case 'havingRaw':
                    foreach ($v as $v1) {
                        $this->builder->{$k}(...$v1);
                    }
                    break;
            } */
        }
        return $this;
    }

    /**
     * 处理order
     *
     * @return self
     */
    final protected function handleOrder(): self
    {
        foreach ($this->order as $v) {
            $this->builder->{$v['method']}(...$v['param']);
        }
        return $this;
    }

    /**
     * 处理join
     *
     * @return self
     */
    final protected function handleJoin(): self
    {
        foreach ($this->join as $v) {
            $this->builder->{$v['method']}(...$v['param']);
        }
        return $this;
    }
}
