<?php

declare(strict_types=1);

namespace app\module\db\dao;

use support\Db;

abstract class AbstractDao
{
    protected string|null $connection;  //分库情况下，解析后所确定的连接
    protected string $table = '';   //分表情况下，解析后所确定的表

    protected string $tableRaw = '';    //表的原生表达式。当需要强制索引等特殊情况时使用。示例：Db::raw('table AS alias FORCE INDEX (索引)')。
    protected array $insert = [];   //解析后的insert。格式：[['字段' => '值',...],...]。无顺序要求
    protected array $update = [];   //解析后的update。格式：['字段' => '值',...]。无顺序要求
    protected array $field = [];   //解析后的field。格式：['select' => [字段,...], 'selectRaw' => [[字段（填selectRaw方法所需参数）],...]]。无顺序要求
    protected array $where = [];   //解析后的where。格式：[['method'=>'where', 'param'=>[参数]],...]。有顺序要求，必须与原来一致，改变顺序可能造成索引使用不同
    protected array $group = [];   //解析后的group。格式：[['method'=>'groupBy', 'param'=>[参数]],...]。有顺序要求，必须与原来一致，改变顺序会造成分组结果不同
    protected array $having = [];  //解析后的having。格式：['having'=>[[条件],...], 'havingRaw'=>[[条件],...]]。无顺序要求
    protected array $order = [];   //解析后的order。格式：[['method'=>'orderBy', 'param'=>[参数]],...]。有顺序要求，必须与原来一致，改变顺序会造成排序结果不同
    protected array $join = [];    //解析后的join。格式：[['method'=>'join', 'param'=>[参数]],...]。无顺序要求

    protected array $fieldAfter = [];    //获取数据库数据后，再做处理的字段

    protected $model;   //模型
    protected \Illuminate\Database\Query\Builder $builder; //构造器

    /* final public function __construct(array $tableSelectData = [], array $connectionSelectData = [])
    {
        //$this->connection($connectionSelectData)->table($tableSelectData);
    } */

    /**
     * 获取连接
     *
     * @return string|null
     */
    final public function getConnection(): string|null
    {
        return empty($this->connection) ? $this->model->getConnectionName() : $this->connection;
    }

    /**
     * 获取表名
     *
     * @return string
     */
    final public function getTable(): string
    {
        return empty($this->table) ? $this->model->getTable() : $this->table;
    }

    /**
     * 获取主键名
     *
     * @return string
     */
    final public function getKey(): string
    {
        return $this->model->getKeyName();
    }

    /**
     * 获取全部列
     *
     * @return array
     */
    final public function getAllColumn(): array
    {
        return $this->model->getAllColumn();
    }

    /*----------------解析 开始----------------*/
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
            $tableRaw = str_replace('__TABLE__', $this->getTable(), $tableRaw);
            $this->tableRaw = Db::raw($tableRaw);
        }
        return $this;
    }

    /**
     * 解析insert（入口）
     *
     * @param array $insert 格式：['字段' => '值',...] 或 [['字段' => '值',...],...]
     * @return self
     */
    final public function insert(array $insert): self
    {
        if (isset($insert[0]) && is_array($insert[0])) {
            foreach ($insert as $k => $v) {
                foreach ($v as $k1 => $v1) {
                    if (!$this->insertOfAlone($k1, $v1, $k)) {
                        $this->insertOfCommon($k1, $v1, $k);
                    }
                }
            }
        } else {
            foreach ($insert as $k => $v) {
                if (!$this->insertOfAlone($k, $v)) {
                    $this->insertOfCommon($k, $v);
                }
            }
        }
        return $this;
    }

    /**
     * 解析update（入口）
     *
     * @param array $update 格式：['字段' => '值',...]
     * @return self
     */
    final public function update(array $update): self
    {
        foreach ($update as $k => $v) {
            if (!$this->updateOfAlone($k, $v)) {
                $this->updateOfCommon($k, $v);
            }
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
        foreach ($field as $v) {
            if (!$this->fieldOfAlone($v)) {
                $this->fieldOfCommon($v);
            }
        }
        return $this;
    }

    /**
     * 解析where（入口）
     *
     * @param array $where  格式：['字段' => '值', ['字段'，'运算符', '值', 'and|or'],...]
     * @return self
     */
    final public function where(array $where): self
    {
        foreach ($where as $k => $v) {
            if (is_numeric($k) && is_array($v)) {
                if (!$this->whereOfAlone(...$v)) {
                    $this->whereOfCommon(...$v);
                }
            } else {
                if (!$this->whereOfAlone($k, null, $v)) {
                    $this->whereOfCommon($k, null, $v);
                }
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
        foreach ($group as $v) {
            if (!$this->groupOfAlone($v)) {
                $this->groupOfCommon($v);
            }
        }
        return $this;
    }

    /**
     * 解析having（入口）
     *
     * @param array $having 格式：['字段' => '值', ['字段'，'运算符', '值', 'and|or'],...]
     * @return self
     */
    final public function having(array $having): self
    {
        foreach ($having as $k => $v) {
            if (is_numeric($k) && is_array($v)) {
                if (!$this->havingOfAlone(...$v)) {
                    $this->havingOfCommon(...$v);
                }
            } else {
                if (!$this->havingOfAlone($k, null, $v)) {
                    $this->havingOfCommon($k, null, $v);
                }
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
        foreach ($order as $k => $v) {
            if (!$this->orderOfAlone($k, $v)) {
                $this->orderOfCommon($k, $v);
            }
        }
        return $this;
    }

    /**
     * 解析insert（公共的）
     *
     * @param string $key
     * @param [type] $value
     * @param integer $index
     * @return boolean
     */
    final protected function insertOfCommon(string $key, $value, int $index = 0): bool
    {
        switch ($key) {
            case 'id':
                $this->insert[$index][$this->getKey()] = $value;
                return true;
            default:
                //数据库不存在的字段过滤掉
                if (in_array($key, $this->getAllColumn())) {
                    //$this->insert[$index][$this->getTable() . '.' . $key] = $value;
                    $this->insert[$index][$key] = $value;
                    return true;
                }
        }
        return false;
    }

    /**
     * 解析update（公共的）
     *
     * @param string $key
     * @param [type] $value
     * @return boolean
     */
    final protected function updateOfCommon(string $key, $value): bool
    {
        switch ($key) {
            case 'id':
                $this->update[$this->getKey()] = $value;
                return true;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    $this->update[$this->getTable() . '.' . $key] = $value;
                } else {
                    $this->update[$key] = $value;
                }
                return true;
        }
        return false;
    }

    /**
     * 解析field（公共的）
     *
     * @param string $key
     * @return boolean
     */
    final protected function fieldOfCommon(string $key): bool
    {
        switch ($key) {
            case '*':
                $this->field['select'][] = $key;
                return true;
            case 'id':
                $this->field['select'][] = $this->getTable() . '.' . $this->getKey();
                return true;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    $this->field['select'][] = $this->getTable() . '.' . $key;
                } else {
                    $this->field['select'][] = $key;
                }
                return true;
        }
        return false;
    }

    /**
     * 解析where（公共的）
     *
     * @param string $key
     * @param string|null $operator
     * @param [type] $value
     * @param string|null $boolean
     * @return boolean
     */
    final protected function whereOfCommon(string $key, string $operator = null, $value, string $boolean = null): bool
    {
        switch ($key) {
            case 'id':
                $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.' . $this->getKey(), $operator ?? '=', $value, $boolean ?? 'and']];
                return true;
            case 'excId':
                $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.' . $this->getKey(), $operator ?? '<>', $value, $boolean ?? 'and']];
                return true;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    $this->where[] = ['method' => 'where', 'param' => [$this->getTable() . '.' . $key, $operator ?? '=', $value, $boolean ?? 'and']];
                } else {
                    $this->where[] = ['method' => 'where', 'param' => [$key, $operator ?? '=', $value, $boolean ?? 'and']];
                }
                return true;
        }
        return false;
    }

    /**
     * 解析group（公共的）
     *
     * @param string $key
     * @return boolean
     */
    final protected function groupOfCommon(string $key): bool
    {
        switch ($key) {
            case 'id':
                $this->group[] = ['method' => 'groupBy', 'param' => [$this->getTable() . '.' . $this->getKey()]];
                return true;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    $this->group[] = ['method' => 'groupBy', 'param' => [$this->getTable() . '.' . $key]];
                } else {
                    $this->group[] = ['method' => 'groupBy', 'param' => [$key]];
                }
                return true;
        }
        return false;
    }

    /**
     * 解析having（公共的）
     *
     * @param string $key
     * @param [type] $value
     * @return boolean
     */
    final protected function havingOfCommon(string $key, string $operator = null, $value, string $boolean = null): bool
    {
        switch ($key) {
            case 'id':
                $this->having['having'][] = [$this->getTable() . '.' . $this->getKey(), $operator ?? '=', $value, $boolean ?? 'and'];
                return true;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    $this->having['having'][] = [$this->getTable() . '.' . $key, $operator ?? '=', $value, $boolean ?? 'and'];
                } else {
                    $this->having['having'][] = [$key, $operator ?? '=', $value, $boolean ?? 'and'];
                }
                return true;
        }
        return false;
    }

    /**
     * 解析order（公共的）
     *
     * @param string $key
     * @param [type] $value
     * @return boolean
     */
    final protected function orderOfCommon(string $key, $value): bool
    {
        switch ($key) {
            case 'id':
                $this->order[] = ['method' => 'orderBy', 'param' => [$this->getTable() . '.' . $this->getKey(), $value]];
                return true;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    $this->order[] = ['method' => 'orderBy', 'param' => [$this->getTable() . '.' . $key, $value]];
                } else {
                    $this->order[] = ['method' => 'orderBy', 'param' => [$key, $value]];
                }
                return true;
        }
        return false;
    }

    /**
     * 解析insert（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @param integer $index
     * @return boolean
     */
    protected function insertOfAlone(string $key, $value, int $index = 0): bool
    {
        /* switch ($key) {
            case 'xxxx':
                $this->insert[$index][$key] = $value;
                return true;
        } */
        return false;
    }

    /**
     * 解析update（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @return boolean
     */
    final protected function updateOfAlone(string $key, $value = null): bool
    {
        /* switch ($key) {
            case 'xxxx':
                $this->update[$key] = $value;
                return true;
        } */
        return false;
    }

    /**
     * 解析field（独有的）
     *
     * @param string $key
     * @return self
     */
    protected function fieldOfAlone(string $key): bool
    {
        /* switch ($key) {
            case 'xxxx':
                $this->field['select'][] = $key;
                //$this->field['select'][] = Db::raw('JSON_UNQUOTE(JSON_EXTRACT(extendData, "$.' . $key . '")) AS ' . $key);    //不能防sql注入
                //$this->field['selectRaw'][] = ['IFNULL(字段名, \'\') AS ?', [$key]];  //能防sql注入
                return true;
        } */
        return false;
    }

    /**
     * 解析where（独有的）
     *
     * @param string $key
     * @param string|null $operator
     * @param [type] $value
     * @param string|null $boolean
     * @return boolean
     */
    protected function whereOfAlone(string $key, string $operator = null, $value, string $boolean = null): bool
    {
        /* switch ($key) {
            case 'xxxx':
                $this->where[] = ['method' => 'where', 'param' => [$key, $operator ?? '=', $value, $boolean ?? 'and']];
                //$this->where[] = ['method' => 'whereRaw', 'param' => [':key > :value', ['key' => $key, 'value' => $value], $boolean ?? 'and']];
                return true;
        } */
        return false;
    }

    /**
     * 解析group（独有的）
     *
     * @param string $key
     * @return boolean
     */
    protected function groupOfAlone(string $key): bool
    {
        /* switch ($key) {
            case 'xxxx':
                $this->group[] = ['method' => 'groupBy', 'param' => [$key]];
                //$this->group[] = ['method'=>'groupByRaw', 'param'=>[':key', ['key' => $key]]];
                return true;
        } */
        return false;
    }

    /**
     * 解析having（独有的）
     *
     * @param string $key
     * @param string|null $operator
     * @param [type] $value
     * @param string|null $boolean
     * @return boolean
     */
    protected function havingOfAlone(string $key, string $operator = null, $value, string $boolean = null): bool
    {
        /* switch ($key) {
            case 'xxxx':
                $this->having['having'][] = [$key, '=', $value];
                //$this->having['havingRaw'][] = [':key > :value', ['key' => $key, 'value' => $value], 'and'];
                return true;
        } */
        return false;
    }

    /**
     * 解析order（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @return boolean
     */
    protected function orderOfAlone(string $key, $value = null): bool
    {
        /* switch ($key) {
            case 'xxxx':
                $this->order[] = ['method' => 'orderBy', 'param' => [$key, $value]];
                //$this->order[] = ['method'=>'orderByRaw', 'param'=>[':key ' . $value, ['key' => $key]]];
                return true;
        } */
        return false;
    }

    /**
     * 解析join（独有的）
     *
     * @param string $key   键，用于确定关联表
     * @param [type] $value 值，用于确定关联表
     * @return boolean
     */
    protected function joinOfAlone(string $key, $value = null): bool
    {
        return false;
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
     * 重置解析数据，用于复用对象，为了不影响下一条sql生成执行（可选，依情况使用）
     *
     * @return self
     */
    final public function resetParseData(): self
    {
        $this->tableRaw = '';
        $this->insert = [];
        $this->update = [];
        $this->field = [];
        $this->where = [];
        $this->group = [];
        $this->having = [];
        $this->order = [];
        $this->join = [];

        $this->fieldAfter = [];
        return $this;
    }

    /**
     * 重置分表分库数据，用于复用对象，为了不影响下一条sql生成执行（可选，依情况使用）
     *
     * @return self
     */
    final public function resetSelectData(): self
    {
        $this->connection = '';
        $this->table = '';
        return $this;
    }
    /*----------------解析 结束----------------*/

    /*----------------解析后处理 开始----------------*/
    /**
     * 获取Db构造器
     *
     * @return \Illuminate\Database\Query\Builder
     */
    final public function getBuilder(): \Illuminate\Database\Query\Builder
    {
        if (empty($this->tableRaw)) {
            $this->builder = Db::connection($this->getConnection())->table($this->getTable());
        } else {
            $this->builder = Db::connection($this->getConnection())->table($this->tableRaw);
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
                    //$this->builder->{$k}($v);
                    $this->builder->{$k}(array_unique($v));
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
    /*----------------解析后处理 结束----------------*/

    /*----------------封装部分方法方便使用 开始----------------*/
    /**
     * 保存插入
     *
     * @param boolean $isGetId
     * @return boolean|integer
     */
    final public function saveInsert(bool $isGetId = true): bool|int
    {
        $this->getBuilder();
        if ((isset($this->insert[0]) && is_array($this->insert[0])) || !$isGetId) {
            return $this->builder->insert($this->insert);
        }
        return $this->builder->insertGetId($this->insert);
    }

    /**
     * 保存更新
     *
     * @param integer $offset
     * @param integer $limit
     * @return integer
     */
    final public function saveUpdate(int $offset = 0, int $limit = 0): int
    {
        $this->getBuilder();
        $this->handleLimit($offset, $limit);
        return $this->builder->update($this->update);
    }

    /**
     * 获取信息
     *
     * @param boolean $isUseWriter
     * @return object
     */
    final public function getInfo(bool $isUseWriter = false): object
    {
        $this->getBuilder();
        if ($isUseWriter) {
            $this->builder->useWritePdo();
        }
        $info = $this->builder->first();
        return $this->handleFieldAfter($info);
    }

    /**
     * 获取列表
     *
     * @param integer $offset
     * @param integer $limit
     * @param boolean $isUseWriter  读写分离时，是否使用写库读（因读写分离有延迟，有些时候需要使用写库读取）
     * @return array
     */
    final public function getList(int $offset = 0, int $limit = 0, bool $isUseWriter = false): array
    {
        $this->getBuilder();
        if ($isUseWriter) {
            $this->builder->useWritePdo();
        }
        $this->handleLimit($offset, $limit);
        $list = $this->builder->get()->toArray();
        if (!empty($this->fieldAfter)) {
            foreach ($list as &$v) {
                $v = $this->handleFieldAfter($v);
            }
        }
        return $list;
    }

    /**
     * 删除
     *
     * @return integer
     */
    final public function delete(): int
    {
        $this->getBuilder();
        return $this->builder->delete();
    }

    /**
     * 获取数据库数据后，再做处理的字段
     *
     * @param object $info
     * @return object
     */
    protected function handleFieldAfter(object $info): object
    {
        /* foreach ($this->fieldAfter as $field) {
            switch ($field) {
                case 'xxxx':
                    $info->xxxx = 'xxxx';
                    break;
            }
        } */
        return $info;
    }

    /**
     * 处理limit（limit不算常用，故不做解析，也不做解析后处理，即不在getBuilder方法中处理）
     *
     * @return self
     */
    final protected function handleLimit(int $offset, int $limit): self
    {
        if ($limit > 0) {
            $this->builder->offset($offset)->limit($limit);
        } elseif ($offset > 0) {    //当offset>0，limit==0时表示取剩下全部数据。需要limit足够大，这里写99999999，这样还不够的话，服务器也抗不住了
            $this->builder->offset($offset)->limit(99999999);
        }
        return $this;
    }
    /*----------------封装部分方法方便使用 结束----------------*/
}
