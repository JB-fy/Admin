package dao

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
)

// 定义接口
type DaoInterface interface {
	ParseDbGroup(ctx context.Context, dbGroupSelData ...map[string]interface{}) string
	ParseDbTable(ctx context.Context, dbTableSelData ...map[string]interface{}) string
	ParseDbCtx(ctx context.Context, dbSelDataList ...map[string]interface{}) *gdb.Model
	ParseInsert(insert map[string]interface{}) gdb.ModelHandler
	HookInsert(data map[string]interface{}) gdb.HookHandler
	ParseUpdate(update map[string]interface{}) gdb.ModelHandler
	HookUpdate(data map[string]interface{}, idArr ...uint) gdb.HookHandler
	HookDelete(idArr ...uint) gdb.HookHandler
	ParseField(field []string, fieldWithParam map[string]interface{}, afterField *[]string, afterFieldWithParam map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler
	HookSelect(afterField *[]string, afterFieldWithParam map[string]interface{}) gdb.HookHandler
	ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler
	ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler
	ParseOrder(order []string, joinTableArr *[]string) gdb.ModelHandler
	ParseJoin(joinCode string, joinTableArr *[]string) gdb.ModelHandler

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
	ctx context.Context
	dao DaoInterface
	// group               string //分库情况下，解析后所确定的连接
	// table               string //分表情况下，解析后所确定的表
	model               *gdb.Model
	afterField          *[]string
	afterFieldWithParam map[string]interface{}
	joinTableArr        *[]string
}

func NewDaoHandler(ctx context.Context, dao DaoInterface, dbSelDataList ...map[string]interface{}) *DaoHandler {
	daoHandlerThisObj := DaoHandler{
		ctx:                 ctx,
		dao:                 dao,
		afterField:          &[]string{},
		afterFieldWithParam: map[string]interface{}{},
		joinTableArr:        &[]string{},
	}
	// daoHandlerThisObj.group = daoHandlerThisObj.dao.ParseDbGroup(ctx, dbSelDataList[0])
	// daoHandlerThisObj.table = daoHandlerThisObj.dao.ParseDbTable(ctx, dbSelDataList[1])
	// daoHandlerThisObj.model = g.DB(daoHandlerThisObj.group).Model(daoHandlerThisObj.table).Ctx(ctx)
	daoHandlerThisObj.model = daoHandlerThisObj.dao.ParseDbCtx(daoHandlerThisObj.ctx, dbSelDataList...)
	return &daoHandlerThisObj
}

func (daoHandlerThis *DaoHandler) Insert(data map[string]interface{}) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseInsert(data))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Update(data map[string]interface{}) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseUpdate(data))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) HookUpdate(hookData map[string]interface{}, idArr ...uint) *DaoHandler {
	if len(hookData) > 0 {
		daoHandlerThis.model = daoHandlerThis.model.Hook(daoHandlerThis.dao.HookUpdate(hookData, idArr...))
	}
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) HookDelete(idArr ...uint) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Hook(daoHandlerThis.dao.HookDelete(idArr...))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Field(field []string, fieldWithParamL ...map[string]interface{}) *DaoHandler {
	fieldWithParam := map[string]interface{}{}
	if len(fieldWithParamL) > 0 {
		fieldWithParam = fieldWithParamL[0]
	}
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseField(field, fieldWithParam, daoHandlerThis.afterField, daoHandlerThis.afterFieldWithParam, daoHandlerThis.joinTableArr))
	if len(*daoHandlerThis.afterField) > 0 || len(daoHandlerThis.afterFieldWithParam) > 0 {
		daoHandlerThis.model = daoHandlerThis.model.Hook(daoHandlerThis.dao.HookSelect(daoHandlerThis.afterField, daoHandlerThis.afterFieldWithParam))
	}
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Filter(filter map[string]interface{}) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseFilter(filter, daoHandlerThis.joinTableArr))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Group(group []string) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseGroup(group, daoHandlerThis.joinTableArr))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Order(order []string) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseOrder(order, daoHandlerThis.joinTableArr))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) Join(joinCode string) *DaoHandler {
	daoHandlerThis.model = daoHandlerThis.model.Handler(daoHandlerThis.dao.ParseJoin(joinCode, daoHandlerThis.joinTableArr))
	return daoHandlerThis
}

func (daoHandlerThis *DaoHandler) GetModel(isClone ...bool) *gdb.Model {
	if len(isClone) > 0 && isClone[0] {
		return daoHandlerThis.model.Clone()
	}
	return daoHandlerThis.model
}

// 判断是否联表
func (daoHandlerThis *DaoHandler) IsJoin() bool {
	return len(*daoHandlerThis.joinTableArr) > 0
}

// 联表则GroupBy主键
func (daoHandlerThis *DaoHandler) JoinGroupByPrimaryKey() *DaoHandler {
	if daoHandlerThis.IsJoin() {
		daoHandlerThis.model = daoHandlerThis.model.Group(daoHandlerThis.dao.ParseDbTable(daoHandlerThis.ctx) + `.` + daoHandlerThis.dao.PrimaryKey())
	}
	return daoHandlerThis
}

// 总数（有联表默认group主键）
func (daoHandlerThis *DaoHandler) Count() (count int, err error) {
	if daoHandlerThis.IsJoin() {
		tableThis := daoHandlerThis.dao.ParseDbTable(daoHandlerThis.ctx)
		count, err = daoHandlerThis.GetModel(true).Group(tableThis + `.` + daoHandlerThis.dao.PrimaryKey()).Distinct().Fields(tableThis + `.` + daoHandlerThis.dao.PrimaryKey()).Count()
		return
	}
	count, err = daoHandlerThis.model.Count()
	return
}
