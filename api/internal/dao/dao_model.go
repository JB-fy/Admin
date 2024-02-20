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
	ParseInsert(insert map[string]interface{}, daoModel *DaoModel) gdb.ModelHandler
	HookInsert(daoModel *DaoModel) gdb.HookHandler
	ParseUpdate(update map[string]interface{}, daoModel *DaoModel) gdb.ModelHandler
	HookUpdate(daoModel *DaoModel) gdb.HookHandler
	HookDelete(daoModel *DaoModel) gdb.HookHandler
	ParseField(field []string, fieldWithParam map[string]interface{}, daoModel *DaoModel) gdb.ModelHandler
	HookSelect(daoModel *DaoModel) gdb.HookHandler
	ParseFilter(filter map[string]interface{}, daoModel *DaoModel) gdb.ModelHandler
	ParseGroup(group []string, daoModel *DaoModel) gdb.ModelHandler
	ParseOrder(order []string, daoModel *DaoModel) gdb.ModelHandler
	ParseJoin(joinTable string, daoModel *DaoModel) gdb.ModelHandler

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

type DaoModel struct {
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

func NewDaoModel(ctx context.Context, dao DaoInterface, dbOpt ...map[string]interface{}) *DaoModel {
	daoModelObj := DaoModel{
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
		daoModelObj.DbGroup = daoModelObj.dao.ParseDbGroup(ctx, dbOpt[0])
		daoModelObj.DbTable = daoModelObj.dao.ParseDbTable(ctx)
	case 2:
		daoModelObj.DbGroup = daoModelObj.dao.ParseDbGroup(ctx, dbOpt[0])
		daoModelObj.DbTable = daoModelObj.dao.ParseDbTable(ctx, dbOpt[1])
	default:
		daoModelObj.DbGroup = daoModelObj.dao.ParseDbGroup(ctx)
		daoModelObj.DbTable = daoModelObj.dao.ParseDbTable(ctx)
	}
	daoModelObj.model = daoModelObj.NewModel()
	return &daoModelObj
}

/*--------业务可能用到的方法 开始--------*/
// 生成模型
func (daoModelThis *DaoModel) NewModel() *gdb.Model {
	return g.DB(daoModelThis.DbGroup).Model(daoModelThis.DbTable). /* Safe(). */ Ctx(daoModelThis.Ctx)
}

// 返回当前模型的副本（当外部还需要做特殊处理时使用）
func (daoModelThis *DaoModel) CloneModel() *gdb.Model {
	return daoModelThis.model.Clone()
}

// 返回当前模型（当外部还需要做特殊处理时使用）
func (daoModelThis *DaoModel) GetModel() *gdb.Model {
	return daoModelThis.model
}

// 一般在更新|删除操作需要做后置处理时使用，注意：必须在filter条件都设置完成后使用
func (daoModelThis *DaoModel) SetIdArr() *DaoModel {
	idArr, _ := daoModelThis.CloneModel().Array(daoModelThis.dao.PrimaryKey())
	daoModelThis.IdArr = gconv.SliceUint(idArr)
	return daoModelThis
}

// 判断是否联表
func (daoModelThis *DaoModel) IsJoin() bool {
	return len(daoModelThis.JoinTableArr) > 0
}

// 联表时，GroupBy主键
func (daoModelThis *DaoModel) GroupPriOnJoin() *DaoModel {
	if daoModelThis.IsJoin() {
		daoModelThis.model = daoModelThis.model.Group(daoModelThis.DbTable + `.` + daoModelThis.dao.PrimaryKey())
	}
	return daoModelThis
}

// 列表（联表时，GroupBy主键）
func (daoModelThis *DaoModel) ListPri() (gdb.Result, error) {
	return daoModelThis.GroupPriOnJoin().All()
}

// 总数（联表时，主键去重）
func (daoModelThis *DaoModel) CountPri() (int, error) {
	if daoModelThis.IsJoin() {
		return daoModelThis.CloneModel().Group(daoModelThis.DbTable + `.` + daoModelThis.dao.PrimaryKey()).Distinct().Fields(daoModelThis.DbTable + `.` + daoModelThis.dao.PrimaryKey()).Count()
	}
	return daoModelThis.Count()
}

// 详情（联表时，GroupBy主键）
func (daoModelThis *DaoModel) InfoPri() (gdb.Record, error) {
	return daoModelThis.GroupPriOnJoin().One()
}

/*--------业务可能用到的方法 结束--------*/

/*--------简化对dao方法的调用 开始--------*/
func (daoModelThis *DaoModel) HookInsert(data map[string]interface{}) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseInsert(data, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) HookUpdate(data map[string]interface{}) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseUpdate(data, daoModelThis))
	if len(daoModelThis.AfterUpdate) > 0 {
		daoModelThis.Hook(daoModelThis.dao.HookUpdate(daoModelThis))
	}
	return daoModelThis
}

func (daoModelThis *DaoModel) HookDelete() *DaoModel {
	daoModelThis.Hook(daoModelThis.dao.HookDelete(daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) Field(field string) *DaoModel {
	return daoModelThis.Fields([]string{field})
}

func (daoModelThis *DaoModel) Fields(field []string) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseField(field, map[string]interface{}{}, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) FieldWithParam(fieldWithParam map[string]interface{}) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseField([]string{}, fieldWithParam, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) HookSelect() *DaoModel {
	if len(daoModelThis.AfterField) > 0 || len(daoModelThis.AfterFieldWithParam) > 0 {
		daoModelThis.Hook(daoModelThis.dao.HookSelect(daoModelThis))
	}
	return daoModelThis
}

func (daoModelThis *DaoModel) Filter(key string, val interface{}) *DaoModel {
	return daoModelThis.Filters(map[string]interface{}{key: val})
}

func (daoModelThis *DaoModel) Filters(filter map[string]interface{}) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseFilter(filter, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) Group(group string) *DaoModel {
	return daoModelThis.Groups([]string{group})
}

func (daoModelThis *DaoModel) Groups(group []string) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseGroup(group, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) Order(order string) *DaoModel {
	return daoModelThis.Orders([]string{order})
}

func (daoModelThis *DaoModel) Orders(order []string) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseOrder(order, daoModelThis))
	return daoModelThis
}

func (daoModelThis *DaoModel) Join(joinTable string) *DaoModel {
	daoModelThis.Handler(daoModelThis.dao.ParseJoin(joinTable, daoModelThis))
	return daoModelThis
}

/*--------简化对dao方法的调用 结束--------*/

/*--------简化对model方法的调用，并封装部分常用方法 开始--------*/
func (daoModelThis *DaoModel) Transaction(f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return daoModelThis.model.Transaction(daoModelThis.Ctx, f)
}

func (daoModelThis *DaoModel) TX(tx gdb.TX) *DaoModel {
	daoModelThis.model = daoModelThis.model.TX(tx)
	return daoModelThis
}

func (daoModelThis *DaoModel) LockShared() *DaoModel {
	daoModelThis.model = daoModelThis.model.LockShared()
	return daoModelThis
}

func (daoModelThis *DaoModel) LockUpdate() *DaoModel {
	daoModelThis.model = daoModelThis.model.LockUpdate()
	return daoModelThis
}

func (daoModelThis *DaoModel) Master() *DaoModel {
	daoModelThis.model = daoModelThis.model.Master()
	return daoModelThis
}

func (daoModelThis *DaoModel) Slave() *DaoModel {
	daoModelThis.model = daoModelThis.model.Slave()
	return daoModelThis
}

func (daoModelThis *DaoModel) Schema(schema string) *DaoModel {
	daoModelThis.model = daoModelThis.model.Schema(schema)
	return daoModelThis
}

func (daoModelThis *DaoModel) Partition(partitions ...string) *DaoModel {
	daoModelThis.model = daoModelThis.model.Partition(partitions...)
	return daoModelThis
}

func (daoModelThis *DaoModel) Union(unions ...*gdb.Model) *DaoModel {
	daoModelThis.model = daoModelThis.model.Union(unions...)
	return daoModelThis
}

func (daoModelThis *DaoModel) UnionAll(unions ...*gdb.Model) *DaoModel {
	daoModelThis.model = daoModelThis.model.UnionAll(unions...)
	return daoModelThis
}

func (daoModelThis *DaoModel) Unscoped() *DaoModel {
	daoModelThis.model = daoModelThis.model.Unscoped()
	return daoModelThis
}

func (daoModelThis *DaoModel) Handler(handlers ...gdb.ModelHandler) *DaoModel {
	daoModelThis.model = daoModelThis.model.Handler(handlers...)
	return daoModelThis
}

func (daoModelThis *DaoModel) Hook(hook gdb.HookHandler) *DaoModel {
	daoModelThis.model = daoModelThis.model.Hook(hook)
	return daoModelThis
}

func (daoModelThis *DaoModel) Cache(option gdb.CacheOption) *DaoModel {
	daoModelThis.model = daoModelThis.model.Cache(option)
	return daoModelThis
}

func (daoModelThis *DaoModel) Data(data ...interface{}) *DaoModel {
	daoModelThis.model = daoModelThis.model.Data(data...)
	return daoModelThis
}

func (daoModelThis *DaoModel) Distinct() *DaoModel {
	daoModelThis.model = daoModelThis.model.Distinct()
	return daoModelThis
}

func (daoModelThis *DaoModel) Page(page, limit int) *DaoModel {
	if limit == 0 {
		return daoModelThis
	}
	daoModelThis.model = daoModelThis.model.Page(page, limit)
	return daoModelThis
}

func (daoModelThis *DaoModel) Offset(offset int) *DaoModel {
	daoModelThis.model = daoModelThis.model.Offset(offset)
	return daoModelThis
}

func (daoModelThis *DaoModel) Limit(limit ...int) *DaoModel {
	daoModelThis.model = daoModelThis.model.Limit(limit...)
	return daoModelThis
}

func (daoModelThis *DaoModel) Save(data ...interface{}) (result sql.Result, err error) {
	return daoModelThis.model.Save(data...)
}

func (daoModelThis *DaoModel) Replace(data ...interface{}) (result sql.Result, err error) {
	return daoModelThis.model.Replace(data...)
}

func (daoModelThis *DaoModel) Insert(data ...interface{}) (result sql.Result, err error) {
	return daoModelThis.model.Insert(data...)
}

func (daoModelThis *DaoModel) InsertAndGetId(data ...interface{}) (lastInsertId int64, err error) {
	return daoModelThis.model.InsertAndGetId(data...)
}

func (daoModelThis *DaoModel) InsertIgnore(data ...interface{}) (result sql.Result, err error) {
	return daoModelThis.model.InsertIgnore(data...)
}

func (daoModelThis *DaoModel) Update(dataAndWhere ...interface{}) (result sql.Result, err error) {
	return daoModelThis.model.Update(dataAndWhere...)
}

func (daoModelThis *DaoModel) UpdateAndGetAffected(dataAndWhere ...interface{}) (affected int64, err error) {
	return daoModelThis.model.UpdateAndGetAffected(dataAndWhere...)
}

func (daoModelThis *DaoModel) Delete(where ...interface{}) (result sql.Result, err error) {
	return daoModelThis.model.Delete(where...)
}

// 封装常用方法
func (daoModelThis *DaoModel) DeleteAndGetAffected(where ...interface{}) (affected int64, err error) {
	result, err := daoModelThis.Delete(where...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (daoModelThis *DaoModel) Chunk(size int, handler gdb.ChunkHandler) {
	daoModelThis.model.Chunk(size, handler)
}

func (daoModelThis *DaoModel) Scan(pointer interface{}, where ...interface{}) error {
	return daoModelThis.model.Scan(pointer, where...)
}

func (daoModelThis *DaoModel) ScanAndCount(pointer interface{}, totalCount *int, useFieldForCount bool) (err error) {
	return daoModelThis.model.ScanAndCount(pointer, totalCount, useFieldForCount)
}

func (daoModelThis *DaoModel) ScanList(structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields ...string) (err error) {
	return daoModelThis.model.ScanList(structSlicePointer, bindToAttrName, relationAttrNameAndFields...)
}

func (daoModelThis *DaoModel) All() (gdb.Result, error) {
	return daoModelThis.model.All()
}

func (daoModelThis *DaoModel) AllAndCount(useFieldForCount bool) (result gdb.Result, totalCount int, err error) {
	return daoModelThis.model.AllAndCount(useFieldForCount)
}

func (daoModelThis *DaoModel) One(where ...interface{}) (gdb.Record, error) {
	return daoModelThis.model.One(where...)
}

func (daoModelThis *DaoModel) Array(fieldsAndWhere ...interface{}) ([]gdb.Value, error) {
	return daoModelThis.model.Array(fieldsAndWhere...)
}

func (daoModelThis *DaoModel) Value(fieldsAndWhere ...interface{}) (gdb.Value, error) {
	return daoModelThis.model.Value(fieldsAndWhere...)
}

// 封装常用方法
func (daoModelThis *DaoModel) Pluck(field string, key string) (gdb.Record, error) {
	list, err := daoModelThis.Fields([]string{field, key}).HookSelect().All()
	if err != nil {
		return nil, err
	}
	if list.IsEmpty() {
		return nil, nil
	}
	result := gdb.Record{}
	for _, v := range list {
		result[v[key].String()] = v[field]
	}
	return result, nil
}

// 封装常用方法
func (daoModelThis *DaoModel) Plucks(field []string, key string) (map[string]gdb.Record, error) {
	list, err := daoModelThis.Fields(append(field, key)).HookSelect().All()
	if err != nil {
		return nil, err
	}
	if list.IsEmpty() {
		return nil, nil
	}
	result := map[string]gdb.Record{}
	for _, v := range list {
		result[v[key].String()] = v
	}
	return result, nil
}

func (daoModelThis *DaoModel) HasField(field string) (bool, error) {
	return daoModelThis.model.HasField(field)
}

func (daoModelThis *DaoModel) Count(where ...interface{}) (int, error) {
	return daoModelThis.model.Count(where...)
}

func (daoModelThis *DaoModel) CountColumn(column string) (int, error) {
	return daoModelThis.model.CountColumn(column)
}

func (daoModelThis *DaoModel) Sum(column string) (float64, error) {
	return daoModelThis.model.Sum(column)
}

func (daoModelThis *DaoModel) Avg(column string) (float64, error) {
	return daoModelThis.model.Avg(column)
}

func (daoModelThis *DaoModel) Min(column string) (float64, error) {
	return daoModelThis.model.Min(column)
}

func (daoModelThis *DaoModel) Max(column string) (float64, error) {
	return daoModelThis.model.Max(column)
}

func (daoModelThis *DaoModel) Increment(column string, amount interface{}) (sql.Result, error) {
	return daoModelThis.model.Increment(column, amount)
}

func (daoModelThis *DaoModel) Decrement(column string, amount interface{}) (sql.Result, error) {
	return daoModelThis.model.Decrement(column, amount)
}

/*--------简化对gdb.Model方法的调用，并封装部分常用方法 结束--------*/
