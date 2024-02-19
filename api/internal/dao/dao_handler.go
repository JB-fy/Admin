package dao

import (
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// 定义接口
type DaoInterface interface {
	ParseDbGroup(ctx context.Context, dbGroupOpt ...map[string]interface{}) string
	ParseDbTable(ctx context.Context, dbTableOpt ...map[string]interface{}) string
	ParseDbCtx(ctx context.Context, dbOpt ...map[string]interface{}) *gdb.Model
	ParseInsert(insert map[string]interface{}, daoHandler *DaoHandler) gdb.ModelHandler
	HookInsert(daoHandler *DaoHandler) gdb.HookHandler
	ParseUpdate(update map[string]interface{}, daoHandler *DaoHandler) gdb.ModelHandler
	HookUpdate(daoHandler *DaoHandler) gdb.HookHandler
	HookDelete(daoHandler *DaoHandler) gdb.HookHandler
	ParseField(field []string, fieldWithParam map[string]interface{}, daoHandler *DaoHandler) gdb.ModelHandler
	HookSelect(daoHandler *DaoHandler) gdb.HookHandler
	ParseFilter(filter map[string]interface{}, daoHandler *DaoHandler) gdb.ModelHandler
	ParseGroup(group []string, daoHandler *DaoHandler) gdb.ModelHandler
	ParseOrder(order []string, daoHandler *DaoHandler) gdb.ModelHandler
	ParseJoin(joinTable string, daoHandler *DaoHandler) gdb.ModelHandler

	DB() gdb.DB
	Table() string
	// Columns() SceneColumns
	// Columns() struct{}
	// Columns() interface{}
	Group() string
	Ctx(ctx context.Context) *gdb.Model
	Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error)

	PrimaryKey() string
	ColumnArr() []string
	ColumnArrG() *garray.StrArray
}

type DaoHandler struct {
	Ctx                 context.Context
	dao                 DaoInterface
	model               *gdb.Model
	DbGroup             string //分库情况下，解析后所确定的库
	DbTable             string //分表情况下，解析后所确定的表
	IdArr               []uint
	AfterInsert         map[string]interface{}
	AfterUpdate         map[string]interface{}
	AfterField          []string
	AfterFieldWithParam map[string]interface{}
	JoinTableArr        []string
}

func NewDaoHandler(ctx context.Context, dao DaoInterface, dbOpt ...map[string]interface{}) *DaoHandler {
	daoHandlerObj := DaoHandler{
		Ctx:                 ctx,
		dao:                 dao,
		IdArr:               []uint{},
		AfterInsert:         map[string]interface{}{},
		AfterUpdate:         map[string]interface{}{},
		AfterField:          []string{},
		AfterFieldWithParam: map[string]interface{}{},
		JoinTableArr:        []string{},
	}
	switch len(dbOpt) {
	case 1:
		daoHandlerObj.DbGroup = daoHandlerObj.dao.ParseDbGroup(ctx, dbOpt[0])
		daoHandlerObj.DbTable = daoHandlerObj.dao.ParseDbTable(ctx)
	case 2:
		daoHandlerObj.DbGroup = daoHandlerObj.dao.ParseDbGroup(ctx, dbOpt[0])
		daoHandlerObj.DbTable = daoHandlerObj.dao.ParseDbTable(ctx, dbOpt[1])
	default:
		daoHandlerObj.DbGroup = daoHandlerObj.dao.ParseDbGroup(ctx)
		daoHandlerObj.DbTable = daoHandlerObj.dao.ParseDbTable(ctx)
	}
	daoHandlerObj.model = daoHandlerObj.NewModel()
	return &daoHandlerObj
}

func (daoHandlerThis *DaoHandler) HookInsert(data map[string]interface{}) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseInsert(data, daoHandlerThis))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) HookUpdate(data map[string]interface{}) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseUpdate(data, daoHandlerThis))
	if len(daoHandlerThis.AfterUpdate) > 0 {
		daoHandlerThis.model = daoHandlerThis.model.Hook(daoHandlerThis.dao.HookUpdate(daoHandlerThis))
	}
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) HookDelete() *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Hook(daoHandlerThis.dao.HookDelete(daoHandlerThis))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Field(field string) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseField([]string{field}, map[string]interface{}{}, daoHandlerThis))
	return daoHandlerThis
	// return daoHandlerThis.Fields([]string{field})
}

func (daoHandlerThis *DaoHandler) Fields(field []string) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseField(field, map[string]interface{}{}, daoHandlerThis))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) FieldWithParam(fieldWithParam map[string]interface{}) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseField([]string{}, fieldWithParam, daoHandlerThis))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) HookSelect() *DaoHandler {
	if len(daoHandlerThis.AfterField) > 0 || len(daoHandlerThis.AfterFieldWithParam) > 0 {
		daoHandlerThis.model = daoHandlerThis.model.Hook(daoHandlerThis.dao.HookSelect(daoHandlerThis))
	}
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Filter(key string, val interface{}) *DaoHandler {
	return daoHandlerThis.Filters(map[string]interface{}{key: val})
}

func (daoHandlerThis *DaoHandler) Filters(filter map[string]interface{}) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseFilter(filter, daoHandlerThis))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Group(group string) *DaoHandler {
	return daoHandlerThis.Groups([]string{group})
}

func (daoHandlerThis *DaoHandler) Groups(group []string) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseGroup(group, daoHandlerThis))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Order(order string) *DaoHandler {
	return daoHandlerThis.Orders([]string{order})
}

func (daoHandlerThis *DaoHandler) Orders(order []string) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseOrder(order, daoHandlerThis))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Join(joinTable string) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseJoin(joinTable, daoHandlerThis))
	return daoHandlerThis
}

// 生成模型
func (daoHandlerThis *DaoHandler) NewModel() *gdb.Model {
	return g.DB(daoHandlerThis.DbGroup).Model(daoHandlerThis.DbTable). /* Safe(). */ Ctx(daoHandlerThis.Ctx)
}

// 一般在更新|删除操作需要做后置处理时使用，注意：必须在filter条件都设置完成后使用
func (daoHandlerThis *DaoHandler) SetIdArr() *DaoHandler {
	idArr, _ := daoHandlerThis.model.Clone().Array(daoHandlerThis.dao.PrimaryKey())
	daoHandlerThis.IdArr = gconv.SliceUint(idArr)
	return daoHandlerThis
}

// 返回当前模型（当外部还需要做特殊处理时使用）
func (daoHandlerThis *DaoHandler) GetModel(isClone ...bool) *gdb.Model {
	if len(isClone) > 0 && isClone[0] {
		return daoHandlerThis.model.Clone()
	}
	return daoHandlerThis.model
}

// 判断是否联表
func (daoHandlerThis *DaoHandler) isJoin() bool {
	return len(daoHandlerThis.JoinTableArr) > 0
}

// 联表则GroupBy主键
func (daoHandlerThis *DaoHandler) JoinGroupByPrimaryKey() *DaoHandler {
	if daoHandlerThis.isJoin() {
		daoHandlerThis.model = daoHandlerThis.model.Group(daoHandlerThis.DbTable + `.` + daoHandlerThis.dao.PrimaryKey())
	}
	return daoHandlerThis
}

// 列表（有联表默认group主键）
func (daoHandlerThis *DaoHandler) ListOfApi() (gdb.Result, error) {
	return daoHandlerThis.JoinGroupByPrimaryKey().All()
}

// 总数（有联表默认group主键）
func (daoHandlerThis *DaoHandler) CountOfApi() (int, error) {
	if daoHandlerThis.isJoin() {
		return daoHandlerThis.model.Clone().Group(daoHandlerThis.DbTable + `.` + daoHandlerThis.dao.PrimaryKey()).Distinct().Fields(daoHandlerThis.DbTable + `.` + daoHandlerThis.dao.PrimaryKey()).Count()
	}
	return daoHandlerThis.model.Count()
}

// 详情（有联表默认group主键）
func (daoHandlerThis *DaoHandler) OneOfApi() (gdb.Record, error) {
	return daoHandlerThis.JoinGroupByPrimaryKey().One()
}

/*--------复制原模型方法并封装一些常用方法 开始--------*/
// 开启事务
func (daoHandlerThis *DaoHandler) Transaction(f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return daoHandlerThis.model.Transaction(daoHandlerThis.Ctx, f)
}

func (daoHandlerThis *DaoHandler) TX(tx gdb.TX) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.TX(tx)
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) LockShared() *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.LockShared()
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) LockUpdate() *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.LockUpdate()
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Master() *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Master()
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Slave() *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Slave()
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Schema(schema string) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Schema(schema)
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Partition(partitions ...string) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Partition(partitions...)
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Union(unions ...*gdb.Model) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Union(unions...)
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) UnionAll(unions ...*gdb.Model) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.UnionAll(unions...)
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Unscoped() *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Unscoped()
	return daoHandlerThis
}

/* func (daoHandlerThis *DaoHandler) Data(data ...interface{}) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Data(data...)
	return daoHandlerThis
} */

func (daoHandlerThis *DaoHandler) Distinct() *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Distinct()
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Page(page, limit int) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Page(page, limit)
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Offset(offset int) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Offset(offset)
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Limit(limit ...int) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Limit(limit...)
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Save(data ...interface{}) (result sql.Result, err error) {
	return daoHandlerThis.model.Save(data...)
}

func (daoHandlerThis *DaoHandler) Insert(data ...interface{}) (result sql.Result, err error) {
	return daoHandlerThis.model.Insert(data...)
}

func (daoHandlerThis *DaoHandler) InsertAndGetId(data ...interface{}) (lastInsertId int64, err error) {
	return daoHandlerThis.model.InsertAndGetId(data...)
}

func (daoHandlerThis *DaoHandler) InsertIgnore(data ...interface{}) (result sql.Result, err error) {
	return daoHandlerThis.model.InsertIgnore(data...)
}

func (daoHandlerThis *DaoHandler) Update(dataAndWhere ...interface{}) (result sql.Result, err error) {
	return daoHandlerThis.model.Update(dataAndWhere...)
}

func (daoHandlerThis *DaoHandler) UpdateAndGetAffected(dataAndWhere ...interface{}) (affected int64, err error) {
	return daoHandlerThis.model.UpdateAndGetAffected(dataAndWhere...)
}

func (daoHandlerThis *DaoHandler) Delete(where ...interface{}) (result sql.Result, err error) {
	return daoHandlerThis.model.Delete(where...)
}

func (daoHandlerThis *DaoHandler) DeleteAndGetAffected(where ...interface{}) (affected int64, err error) {
	result, err := daoHandlerThis.model.Delete(where...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (daoHandlerThis *DaoHandler) All() (gdb.Result, error) {
	return daoHandlerThis.model.All()
}

func (daoHandlerThis *DaoHandler) AllAndCount(useFieldForCount bool) (result gdb.Result, totalCount int, err error) {
	return daoHandlerThis.model.AllAndCount(useFieldForCount)
}

func (daoHandlerThis *DaoHandler) One(where ...interface{}) (gdb.Record, error) {
	return daoHandlerThis.model.One(where...)
}

func (daoHandlerThis *DaoHandler) Array(fieldsAndWhere ...interface{}) ([]gdb.Value, error) {
	return daoHandlerThis.model.Array(fieldsAndWhere...)
}

func (daoHandlerThis *DaoHandler) Value(fieldsAndWhere ...interface{}) (gdb.Value, error) {
	return daoHandlerThis.model.Value(fieldsAndWhere...)
}

func (daoHandlerThis *DaoHandler) Pluck(key string, val string) (gdb.Record, error) {
	list, err := daoHandlerThis.model.Fields(key, val).All()
	if err != nil {
		return nil, err
	}
	if list.IsEmpty() {
		return nil, nil
	}
	data := gdb.Record{}
	for _, v := range list {
		data[v[key].String()] = v[val]
	}
	return data, nil
}

func (daoHandlerThis *DaoHandler) Count(where ...interface{}) (int, error) {
	return daoHandlerThis.model.Count(where...)
}

func (daoHandlerThis *DaoHandler) CountColumn(column string) (int, error) {
	return daoHandlerThis.model.CountColumn(column)
}

func (daoHandlerThis *DaoHandler) Sum(column string) (float64, error) {
	return daoHandlerThis.model.Sum(column)
}

func (daoHandlerThis *DaoHandler) Avg(column string) (float64, error) {
	return daoHandlerThis.model.Avg(column)
}

func (daoHandlerThis *DaoHandler) Min(column string) (float64, error) {
	return daoHandlerThis.model.Min(column)
}

func (daoHandlerThis *DaoHandler) Max(column string) (float64, error) {
	return daoHandlerThis.model.Max(column)
}

func (daoHandlerThis *DaoHandler) Increment(column string, amount interface{}) (sql.Result, error) {
	return daoHandlerThis.model.Increment(column, amount)
}

func (daoHandlerThis *DaoHandler) Decrement(column string, amount interface{}) (sql.Result, error) {
	return daoHandlerThis.model.Decrement(column, amount)
}

/*--------复制原模型方法并封装一些常用方法 结束--------*/
