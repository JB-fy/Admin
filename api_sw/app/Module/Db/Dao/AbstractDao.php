<?php

declare(strict_types=1);

namespace App\Module\Db\Dao;

use Hyperf\DbConnection\Db;

/**
 * 需使用框架命令快速生成Dao时
 *      继承extends \Hyperf\DbConnection\Model\Model
 *      设置$connection属性为default
 *      注释掉getConnection方法
 *      注释掉update方法
 */
abstract class AbstractDao/*  extends \Hyperf\DbConnection\Model\Model */
{
    protected ?string $connection/*  = 'default' */;  //分库情况下，解析后所确定的连接
    protected ?string $table;   //分表情况下，解析后所确定的表

    protected array $insertData = [];   //解析后的insert。格式：[['字段' => '值',...],...]。无顺序要求
    protected array $updateData = [];   //解析后的update。格式：['字段' => '值',...]。无顺序要求

    protected array $joinCode = [];    //已联表标识。格式：['joinCode',...]

    protected array $afterField = [];    //获取数据后，再处理的字段

    protected array $jsonField = [];    //json类型字段。这些字段创建|更新时，需要特殊处理

    //#[Inject(value: \App\Module\Db\Model\Platform\Admin::class)]
    protected \App\Module\Db\Model\AbstractModel $model;   //模型
    protected \Hyperf\Database\Query\Builder $builder; //构造器

    public function __construct()
    {
        //子类未定义$model时会自动设置。注意：Dao类目录和Model目录的对应关系
        if (empty($this->model)) {
            $modelClassName = str_replace('\\Dao\\', '\\Model\\', get_class($this));
            $this->model = getModel($modelClassName);
        }
        $this->initBuilder();
    }

    /**
     * 获取连接
     *
     * @return string
     */
    final public function getConnection(): string
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
     * 初始化Db构造器
     *
     * @param string $tableRaw  表的原生表达式。当需要强制索引等特殊情况时使用。示例：'__TABLE__ FORCE INDEX (索引)')。
     * @return self
     */
    final public function initBuilder(string $tableRaw = ''): self
    {
        if (empty($tableRaw)) {
            $this->builder = Db::connection($this->getConnection())->table($this->getTable());
        } else {
            $tableRaw = str_replace('__TABLE__', $this->getTable(), $tableRaw);
            $this->builder = Db::connection($this->getConnection())->table(Db::raw($tableRaw));
        }
        return $this;
    }

    /**
     * 获取Db构造器
     *
     * @return \Hyperf\Database\Query\Builder
     */
    final public function getBuilder(bool $isClone = false): \Hyperf\Database\Query\Builder
    {
        if ($isClone) {
            return $this->builder->clone();
        }
        return $this->builder;
    }

    /**
     * 解析insert（入口）
     *
     * @param array $insert 格式：['字段' => '值',...] 或 [['字段' => '值',...],...]
     * @return self
     */
    final public function parseInsert(array $insert): self
    {
        if (isset($insert[0]) && is_array($insert[0])) {
            foreach ($insert as $k => $v) {
                foreach ($v as $k1 => $v1) {
                    $this->parseInsertOne($k1, $v1, $k);
                }
            }
        } else {
            foreach ($insert as $k => $v) {
                $this->parseInsertOne($k, $v);
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
    final public function parseUpdate(array $update): self
    {
        foreach ($update as $k => $v) {
            $this->parseUpdateOne($k, $v);
        }
        return $this;
    }

    /**
     * 解析field（入口）
     *
     * @param array $field  格式：['字段', '字段' => '值',...]
     * @return self
     */
    final public function parseField(array $field): self
    {
        foreach ($field as $k => $v) {
            if (is_numeric($k)) {
                $this->parseFieldOne($v);
            } else {
                $this->parseFieldOne($k, $v);
            }
        }
        return $this;
    }

    /**
     * 获取数据库数据后，再处理的字段（入口）
     *
     * @param object $info
     * @return void
     */
    final protected function afterField(object &$info)
    {
        //isset($info->{$this->getKey()}) ? $info->id = $info->{$this->getKey()} : null;  //设置id字段
        foreach ($this->afterField as $field) {
            $this->parseAfterField($field, $info);
        }
    }

    /**
     * 解析filter（入口）
     *
     * @param array $filter  格式：['字段' => '值', ['字段'，'运算符', '值', 'and|or'],...]
     * @return self
     */
    final public function parseFilter(array $filter): self
    {
        foreach ($filter as $k => $v) {
            if (is_numeric($k) && is_array($v)) {
                $this->parseFilterOne(...$v);
            } else {
                $this->parseFilterOne($k, null, $v);
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
    final public function parseGroup(array $group): self
    {
        foreach ($group as $v) {
            $this->parseGroupOne($v);
        }
        return $this;
    }

    /**
     * 解析order（入口）
     *
     * @param array $order  格式：['字段' => 'asc或desc',...]
     * @return self
     */
    final public function parseOrder(array $order): self
    {
        foreach ($order as $k => $v) {
            $this->parseOrderOne($k, $v);
        }
        return $this;
    }

    /**
     * 解析join（入口）
     *
     * @param string $joinAlias
     * @return self
     */
    final public function parseJoin(string $joinAlias): self
    {
        if (in_array($joinAlias, $this->joinCode)) {
            return $this;
        }
        $this->joinCode[] = $joinAlias;
        $this->parseJoinOne($joinAlias);
        return $this;
    }

    /**
     * 解析insert（单个）
     *
     * @param string $key
     * @param [type] $value
     * @param integer $index
     * @return void
     */
    protected function parseInsertOne(string $key, $value, int $index = 0): void
    {
        switch ($key) {
            case 'id':
                $this->insertData[$index][$this->getKey()] = $value;
                break;
            case 'password':
                if (in_array('salt', $this->getAllColumn())) {
                    $salt = randStr(8);
                    if (strlen($value) != 32) {
                        $value = md5($value);
                    }
                    $this->insertData[$index]['salt'] = $salt;
                    $this->insertData[$index][$key] = md5($value . $salt);
                } else {
                    $this->insertData[$index][$key] = $value;
                }
                break;
            case 'salt':    //password字段处理过程自动生成
                break;
            default:
                //数据库不存在的字段过滤掉
                if (in_array($key, $this->getAllColumn())) {
                    if (in_array($key, $this->jsonField)) {
                        if ($value === '' || $value === null) {
                            $this->insertData[$index][$key] =  null;
                        } else {
                            $this->insertData[$index][$key] =  is_array($value) ? json_encode($value, JSON_UNESCAPED_UNICODE) : $value;
                        }
                    } else {
                        $this->insertData[$index][$key] = $value;
                    }
                }
        }
    }

    /**
     * 解析update（单个）
     *
     * @param string $key
     * @param [type] $value
     * @return void
     */
    protected function parseUpdateOne(string $key, $value): void
    {
        switch ($key) {
            case 'id':
                $this->updateData[$this->getTable() . '.' . $this->getKey()] = $value;
                break;
            case 'password':
                if (in_array('salt', $this->getAllColumn())) {
                    $salt = randStr(8);
                    if (strlen($value) != 32) {
                        $value = md5($value);
                    }
                    $this->updateData[$this->getTable() . '.salt'] = $salt;
                    $this->updateData[$this->getTable() . '.' . $key] = md5($value . $salt);
                } else {
                    $this->updateData[$this->getTable() . '.' . $key] = $value;
                }
                break;
            case 'salt':    //password字段处理过程自动生成
                break;
            default:
                /* //暂时不考虑其它复杂字段。复杂字段建议直接写入parseUpdateOfAlone方法
                list($realKey) = explode('->', $key);   //json情况
                list($realKey) = explode(' AS ', $realKey); //别名情况
                list($realKey) = explode(' as ', $realKey); //别名情况
                $realKey = trim($realKey);  //去除两边空白
                //数据库不存在的字段过滤掉
                if (in_array($realKey, $this->getAllColumn())) { */
                //数据库不存在的字段过滤掉
                if (in_array($key, $this->getAllColumn())) {
                    if (in_array($key, $this->jsonField)) {
                        if ($value === '' || $value === null) {
                            $this->updateData[$this->getTable() . '.' . $key] = null;
                        } else {
                            $this->updateData[$this->getTable() . '.' . $key] = is_array($value) ? json_encode($value, JSON_UNESCAPED_UNICODE) : $value;
                        }
                    } else {
                        $this->updateData[$this->getTable() . '.' . $key] = $value;
                    }
                }
        }
    }

    /**
     * 解析field（单个）
     *
     * @param string $key
     * @param [type] $value
     * @return void
     */
    protected function parseFieldOne(string $key, $value = null): void
    {
        switch ($key) {
                /* case '*':
                $this->builder->addSelect($key);
                break; */
            case 'id':
                $this->builder->addSelect($this->getTable() . '.' . $this->getKey() . ' AS ' . $key);
                break;
            case 'label':
                $nameField = str_replace('Id', 'Name', $this->getKey());
                if (in_array($nameField, $this->getAllColumn())) {
                    $this->builder->addSelect($this->getTable() . '.' . $nameField . ' AS ' . $key);
                }
                break;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    $this->builder->addSelect($this->getTable() . '.' . $key);
                } else {
                    $this->builder->addSelect($key);
                }
                break;
        }
    }

    /**
     * 获取数据后，再处理的字段（单个）
     *
     * @param string $key
     * @param object $info
     * @return void
     */
    protected function parseAfterField(string $key, object &$info): void
    {
        switch ($key) {
            case 'id':
                $info->{$key} = $info->{$this->getKey()};
                break;
            default:
                getLogger('daoAfterField')->warning('未处理字段：' . $key);
        }
    }

    /**
     * 解析filter（单个）
     *
     * @param string $key
     * @param string|null $operator
     * @param [type] $value
     * @param string|null $boolean
     * @return void
     */
    protected function parseFilterOne(string $key, string $operator = null, $value, string $boolean = null): void
    {
        switch ($key) {
            case 'id':
            case 'idArr':
                if (is_array($value)) {
                    if (count($value) === 1) {
                        $this->builder->where($this->getTable() . '.' . $this->getKey(), $operator ?? '=', array_shift($value), $boolean ?? 'and');
                    } else {
                        $this->builder->whereIn($this->getTable() . '.' . $this->getKey(), $value, $boolean ?? 'and');
                    }
                } else {
                    $this->builder->where($this->getTable() . '.' . $this->getKey(), $operator ?? '=', $value, $boolean ?? 'and');
                }
                break;
            case 'excId':
            case 'excIdArr':
                if (is_array($value)) {
                    if (count($value) === 1) {
                        $this->builder->where($this->getTable() . '.' . $this->getKey(), $operator ?? '<>', array_shift($value), $boolean ?? 'and');
                    } else {
                        $this->builder->whereNotIn($this->getTable() . '.' . $this->getKey(), $value, $boolean ?? 'and');
                    }
                } else {
                    $this->builder->where($this->getTable() . '.' . $this->getKey(), $operator ?? '<>', $value, $boolean ?? 'and');
                }
                break;
            case 'timeRangeStart':
                $this->builder->where($this->getTable() . '.createdAt', $operator ?? '>=', date('Y-m-d H:i:s', strtotime($value)), $boolean ?? 'and');
                break;
            case 'timeRangeEnd':
                $this->builder->where($this->getTable() . '.createdAt', $operator ?? '<=', date('Y-m-d H:i:s', strtotime($value)), $boolean ?? 'and');
                break;
            case 'label':
                $nameField = str_replace('Id', 'Name', $this->getKey());
                if (in_array($nameField, $this->getAllColumn())) {
                    $this->builder->where($this->getTable() . '.' . $nameField, $operator ?? 'Like', '%' . $value . '%', $boolean ?? 'and');
                }
                break;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    //id类型字段和部分字段，可通过传递数组做查询
                    if (strtolower(substr($key, -2)) === 'id' || in_array($key, ['configKey'])) {
                        if (is_array($value)) {
                            if (count($value) === 1) {
                                $this->builder->where($this->getTable() . '.' . $key, $operator ?? '=', array_shift($value), $boolean ?? 'and');
                            } else {
                                $this->builder->whereIn($this->getTable() . '.' . $key, $value, $boolean ?? 'and');
                            }
                        } else {
                            $this->builder->where($this->getTable() . '.' . $key, $operator ?? '=', $value, $boolean ?? 'and');
                        }
                    } else if (strtolower(substr($key, -4)) === 'name') {
                        $this->builder->where($this->getTable() . '.' . $key, $operator ?? 'like', '%' . $value . '%', $boolean ?? 'and');
                    } else {
                        $this->builder->where($this->getTable() . '.' . $key, $operator ?? '=', $value, $boolean ?? 'and');
                    }
                } else {
                    $this->builder->where($key, $operator ?? '=', $value, $boolean ?? 'and');
                }
                break;
        }
    }

    /**
     * 解析group（单个）
     *
     * @param string $key
     * @return void
     */
    protected function parseGroupOne(string $key): void
    {
        switch ($key) {
            case 'id':
                $this->builder->groupBy($this->getTable() . '.' . $this->getKey());
                break;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    $this->builder->groupBy($this->getTable() . '.' . $key);
                } else {
                    $this->builder->groupBy($key);
                }
                break;
        }
    }

    /**
     * 解析order（单个）
     *
     * @param string $key
     * @param [type] $value
     * @return void
     */
    protected function parseOrderOne(string $key, $value = null): void
    {
        switch ($key) {
            case 'id':
                $this->builder->orderBy($this->getTable() . '.' . $this->getKey(), $value);
                break;
            case 'sort':
                $this->builder->orderBy($this->getTable() . '.' . $key, $value);
                $this->builder->orderBy($this->getTable() . '.' . $this->getKey(), 'desc');
                break;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    $this->builder->orderBy($this->getTable() . '.' . $key, $value);
                } else {
                    $this->builder->orderBy($key, $value);
                }
                break;
        }
    }

    /**
     * 解析join（单个）
     *
     * @param string $joinAlias
     * @return void
     */
    protected function parseJoinOne(string $joinAlias): void
    {
        switch ($joinAlias) {
                /* case 'xxxxTable':
                $this->builder->leftJoin($joinAlias, $joinAlias . '.xxxxId', '=', $this->getTable() . '.' . $this->getKey());
                // $this->builder->leftJoin(getDao(Xxxx::class)->getTable() . ' AS ' . $joinAlias, $joinAlias . '.xxxxId', '=', $this->getTable() . '.' . $this->getKey());
                break; */
            default:
                $this->builder->leftJoin($joinAlias, $joinAlias . '.' . $this->getKey(), '=', $this->getTable() . '.' . $this->getKey());
        }
    }

    /**
     * 判断是否联表
     *
     * @return boolean
     */
    final public function isJoin(): bool
    {
        return !empty($this->joinCode);
    }
    /*----------------解析 结束----------------*/

    /*----------------封装部分方法方便使用 开始----------------*/
    /**
     * 插入
     *
     * @param boolean $isGetId
     * @return boolean|integer
     */
    final public function insert(bool $isGetId = true): bool|int
    {
        $this->getBuilder();
        if (count($this->insertData) > 1 || !$isGetId) {
            return $this->builder->insert($this->insertData);
        }
        return $this->builder->insertGetId($this->insertData[0]);
    }

    /**
     * 更新
     *
     * @param integer $offset
     * @param integer $limit
     * @return integer
     */
    final public function update(int $offset = 0, int $limit = 0): int
    {
        $this->getBuilder();
        $this->handleLimit($offset, $limit);
        return $this->builder->update($this->updateData);
    }

    /**
     * 删除
     *
     * @return integer
     */
    final public function delete(int $offset = 0, int $limit = 0): int
    {
        $this->getBuilder();
        $this->handleLimit($offset, $limit);
        return $this->builder->delete();
    }

    /**
     * 获取信息
     *
     * @param boolean $isUseWriter
     * @return object|null
     */
    final public function info(bool $isUseWriter = false): object|null
    {
        $this->getBuilder();
        if ($isUseWriter) {
            $this->builder->useWritePdo();
        }
        $info = $this->builder->first();
        if (empty($info)) {
            return $info;
        }
        $this->afterField($info);
        return $info;
    }

    /**
     * 获取列表
     *
     * @param integer $offset
     * @param integer $limit
     * @param boolean $isUseWriter  读写分离时，是否使用写库读（因读写分离有延迟，有些时候需要使用写库读取）
     * @return array
     */
    final public function list(int $offset = 0, int $limit = 0, bool $isUseWriter = false): array
    {
        $this->getBuilder();
        if ($isUseWriter) {
            $this->builder->useWritePdo();
        }
        $this->handleLimit($offset, $limit);
        $list = $this->builder->get()->toArray();
        if (!empty($this->afterField)) {
            $wg = new \Hyperf\Utils\WaitGroup();
            foreach ($list as &$v) {
                $wg->add(1);
                co(function () use ($wg, $v) {
                    // \Swoole\Coroutine::sleep(3);
                    $this->afterField($v);
                    $wg->done();
                });
            }
            $wg->wait();
        }
        return $list;
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
