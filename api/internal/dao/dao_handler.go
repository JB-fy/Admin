package dao

import (
	"context"

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

func (daoHandlerThis *DaoHandler) Insert(data map[string]interface{}) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseInsert(data, daoHandlerThis))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Update(data map[string]interface{}) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseUpdate(data, daoHandlerThis))
	if len(daoHandlerThis.AfterUpdate) > 0 {
		daoHandlerThis.model = daoHandlerThis.model.Hook(daoHandlerThis.dao.HookUpdate(daoHandlerThis))
	}
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Delete() *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Hook(daoHandlerThis.dao.HookDelete(daoHandlerThis))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Field(field string) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseField([]string{field}, map[string]interface{}{}, daoHandlerThis))
	return daoHandlerThis
	// return daoHandlerThis.Fields([]string{field})
}

/* func (daoHandlerThis *DaoHandler) Fields(field []string) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseField(field, map[string]interface{}{}, daoHandlerThis))
	return daoHandlerThis
} */

func (daoHandlerThis *DaoHandler) Fields(field []string, fieldWithParamL ...map[string]interface{}) *DaoHandler {
	fieldWithParam := map[string]interface{}{}
	if len(fieldWithParamL) > 0 {
		fieldWithParam = fieldWithParamL[0]
	}
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseField(field, fieldWithParam, daoHandlerThis))
	if len(daoHandlerThis.AfterField) > 0 || len(daoHandlerThis.AfterFieldWithParam) > 0 {
		daoHandlerThis.model = daoHandlerThis.model.Hook(daoHandlerThis.dao.HookSelect(daoHandlerThis))
	}
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) FieldWithParam(fieldWithParam map[string]interface{}) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseField([]string{}, fieldWithParam, daoHandlerThis))
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
	idArr, _ := daoHandlerThis.GetModel(true).Array(daoHandlerThis.dao.PrimaryKey())
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
		return daoHandlerThis.GetModel(true).Group(daoHandlerThis.DbTable + `.` + daoHandlerThis.dao.PrimaryKey()).Distinct().Fields(daoHandlerThis.DbTable + `.` + daoHandlerThis.dao.PrimaryKey()).Count()
	}
	return daoHandlerThis.model.Count()
}

// 开启事务
func (daoHandlerThis *DaoHandler) Transaction(f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return daoHandlerThis.model.Transaction(daoHandlerThis.Ctx, f)
}

func (daoHandlerThis *DaoHandler) Page(page, limit int) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Page(page, limit)
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) All() (gdb.Result, error) {
	return daoHandlerThis.model.All()
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

func (daoHandlerThis *DaoHandler) Count(where ...interface{}) (int, error) {
	return daoHandlerThis.model.Count(where...)
}
