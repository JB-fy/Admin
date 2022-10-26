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

    protected array $fieldOfCommon = ['id', '*'];   //公共的field
    protected array $whereOfCommon = ['id', 'excId'];   //公共的where
    protected array $groupOfCommon = ['id'];   //公共的group
    protected array $havingOfCommon = ['id'];  //公共的having
    protected array $orderOfCommon = ['id'];   //公共的order

    /* final public function __construct(array $tableSelectData = [], array $connectionSelectData = [])
    {
        $this->connection($connectionSelectData)->table($tableSelectData);
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
    public function connection(array $connectionSelectData = []): self
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
    public function table(array $tableSelectData = []): self
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
    public function tableRaw(string $tableRaw = ''): self
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
     * 解析field（入口）
     *
     * @param array $field  格式：['字段',...]
     * @return self
     */
    final public function field(array $field): self
    {
        $this->fieldOfCommon = [...$this->fieldOfCommon, ...$this->getAllColumn()];
        foreach ($field as $v) {
            if (in_array($v, $this->fieldOfCommon)) {
                $this->fieldOfCommon($v);
            } else {
                $this->fieldOfAlone($v);
            }
        }
        return $this;
    }

    /**
     * 解析where（入口）
     *
     * @param array $where  格式：['字段' => '值',...]
     * @return self
     */
    final public function where(array $where): self
    {
        $this->whereOfCommon = [...$this->whereOfCommon, ...$this->getAllColumn()];
        foreach ($where as $k => $v) {
            if (in_array($k, $this->whereOfCommon)) {
                $this->whereOfCommon($k, $v);
            } else {
                $this->whereOfAlone($k, $v);
            }
        }
        return $this;
    }

    /**
     * 解析group（入口）
     *
     * @param array $group  格式：['字段',...]
     * @return self
     */
    final public function group(array $group): self
    {
        $this->groupOfCommon = [...$this->groupOfCommon, ...$this->getAllColumn()];
        foreach ($group as $v) {
            if (in_array($v, $this->groupOfCommon)) {
                $this->groupOfCommon($v);
            } else {
                $this->groupOfAlone($v);
            }
        }
        return $this;
    }

    /**
     * 解析having（入口）
     *
     * @param array $having 格式：['字段' => '值',...]
     * @return self
     */
    final public function having(array $having): self
    {
        $this->havingOfCommon = [...$this->havingOfCommon, ...$this->getAllColumn()];
        foreach ($having as $k => $v) {
            if (in_array($k, $this->havingOfCommon)) {
                $this->havingOfCommon($k, $v);
            } else {
                $this->havingOfAlone($k, $v);
            }
        }
        return $this;
    }

    /**
     * 解析order（入口）
     *
     * @param array $order  格式：['字段' => 'asc或desc',...]
     * @return self
     */
    final public function order(array $order): self
    {
        $this->orderOfCommon = [...$this->orderOfCommon, ...$this->getAllColumn()];
        foreach ($order as $k => $v) {
            if (in_array($k, $this->orderOfCommon)) {
                $this->orderOfCommon($k, $v);
            } else {
                $this->orderOfAlone($k, $v);
            }
        }
        return $this;
    }

    /**
     * 解析field（公共的）
     *
     * @param string $key
     * @return self
     */
    final protected function fieldOfCommon(string $key): self
    {
        switch ($key) {
            case '*':
                $this->field['select'][] = $key;
                break;
            case 'id':
                $this->field['select'][] = $this->getTableAlias() . '.' . $this->getPrimaryKey();
                break;
            default:
                $this->field['select'][] = $this->getTableAlias() . '.' . $key;
                break;
        }
        return $this;
    }

    /**
     * 解析where（公共的）
     *
     * @param string $key
     * @param [type] $value
     * @return self
     */
    final protected function whereOfCommon(string $key, $value): self
    {
        switch ($key) {
            case 'id':
                $this->where[] = ['method' => 'where', 'param' => [$this->getTableAlias() . '.' . $this->getPrimaryKey(), '=', $value]];
                break;
            case 'excId':
                $this->where[] = ['method' => 'where', 'param' => [$this->getTableAlias() . '.' . $this->getPrimaryKey(), '<>', $value]];
                break;
            default:
                $this->where[] = ['method' => 'where', 'param' => [$this->getTableAlias() . '.' . $key, '=', $value]];
                break;
        }
        return $this;
    }

    /**
     * 解析group（公共的）
     *
     * @param string $key
     * @return self
     */
    final protected function groupOfCommon(string $key): self
    {
        switch ($key) {
            case 'id':
                $this->group[] = ['method' => 'groupBy', 'param' => [$this->getTableAlias() . '.' . $this->getPrimaryKey()]];
                break;
            default:
                $this->group[] = ['method' => 'groupBy', 'param' => [$this->getTableAlias() . '.' . $key]];
                break;
        }
        return $this;
    }

    /**
     * 解析having（公共的）
     *
     * @param string $key
     * @param [type] $value
     * @return self
     */
    final protected function havingOfCommon(string $key, $value): self
    {
        switch ($key) {
            case 'id':
                $this->having['having'][] = [$this->getTableAlias() . '.' . $this->getPrimaryKey(), '=', $value];
                break;
            default:
                $this->having['having'][] = [$this->getTableAlias() . '.' . $key, '=', $value];
                break;
        }
        return $this;
    }

    /**
     * 解析order（公共的）
     *
     * @param string $key
     * @param [type] $value
     * @return self
     */
    final protected function orderOfCommon(string $key, $value): self
    {
        switch ($key) {
            case 'id':
                $this->order[] = ['method' => 'orderBy', 'param' => [$this->getTableAlias() . '.' . $this->getPrimaryKey(), $value]];
                break;
            default:
                $this->order[] = ['method' => 'orderBy', 'param' => [$this->getTableAlias() . '.' . $key, $value]];
                break;
        }
        return $this;
    }

    /**
     * 解析field（独有的）
     *
     * @param string $key
     * @return self
     */
    protected function fieldOfAlone(string $key): self
    {
        /* switch ($key) {
            default:
                $this->field['select'][] = $key;
                //$this->field['selectRaw'][] = ['IFNULL(字段名, \'\') AS ' . $key];
                break;
        } */
        return $this;
    }

    /**
     * 解析where（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @return self
     */
    protected function whereOfAlone(string $key, $value): self
    {
        /* switch ($key) {
            default:
                $this->where[] = ['method' => 'where', 'param' => [$key, '=', $value]];
                //$this->where[] = ['method' => 'whereRaw', 'param' => ['age > :age', ['age' => $v], 'and']];
                break;
        } */
        return $this;
    }

    /**
     * 解析group（独有的）
     *
     * @param string $key
     * @return self
     */
    protected function groupOfAlone(string $key): self
    {
        /* switch ($key) {
            default:
                $this->group[] = ['method' => 'groupBy', 'param' => [$key]];
                //$this->group[] = ['method'=>'groupByRaw', 'param'=>[':xxxx', ['xxxx' => 'xxxx']]];
                break;
        } */
        return $this;
    }

    /**
     * 解析having（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @return self
     */
    protected function havingOfAlone(string $key, $value): self
    {
        /* switch ($key) {
            default:
                $this->having['having'][] = [$key, '=', $value];
                //$this->having['havingRaw'][] = ['age > :age', ['age' => $value], 'and'];
                break;
        } */
        return $this;
    }

    /**
     * 解析order（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @return self
     */
    protected function orderOfAlone(string $key, $value): self
    {
        /* switch ($key) {
            default:
                $this->order[] = ['method' => 'orderBy', 'param' => [$key, $value]];
                //$this->order[] = ['method'=>'orderByRaw', 'param'=>[':time ' . $value, ['time' => time()]]];
                break;
        } */
        return $this;
    }

    /**
     * 解析join（独有的）
     *
     * @param string $key   键，用于确定关联表
     * @param [type] $value 值，用于确定关联表
     * @return self
     */
    protected function joinOfAlone(string $key, $value = null): self
    {
        return $this;
    }

    /**
     * 判断是否联表
     *
     * @return boolean
     */
    final public function isJoin(): bool
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
