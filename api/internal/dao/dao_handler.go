package dao

/* daoHandler := dao.NewDaoHandler(ctx, &daoAuth.Scene)
count, err := daoHandler.Filter(filter).Count()
// count, err := service.AuthScene().Count(ctx, filter)
if err != nil {
	return
}
list, err := daoHandler.Field(field).Order(order).List(page, limit)
// list, err := service.AuthScene().List(ctx, filter, field, order, page, limit) */
import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
)

type DaoHandler struct {
	ctx          context.Context
	dao          DaoInterface
	model        *gdb.Model
	joinTableArr *[]string
}

func NewDaoHandler(ctx context.Context, dao DaoInterface, dbSelDataList ...map[string]interface{}) *DaoHandler {
	daoHandlerObj := DaoHandler{
		ctx:          ctx,
		dao:          dao,
		joinTableArr: &[]string{},
	}
	daoHandlerObj.model = daoHandlerObj.dao.ParseDbCtx(daoHandlerObj.ctx, dbSelDataList...)
	return &daoHandlerObj
}

// 定义接口
type DaoInterface interface {
	ParseDbGroup(ctx context.Context, dbGroupSeldata ...map[string]interface{}) string
	ParseDbTable(ctx context.Context, dbTableSelData ...map[string]interface{}) string
	ParseDbCtx(ctx context.Context, dbSelDataList ...map[string]interface{}) *gdb.Model
	ParseInsert(insert map[string]interface{}, fill ...bool) gdb.ModelHandler
	HookInsert(data map[string]interface{}) gdb.HookHandler
	ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler
	HookUpdate(data map[string]interface{}, idArr ...int) gdb.HookHandler
	HookDelete(idArr ...int) gdb.HookHandler
	ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler
	HookSelect(afterField []string) gdb.HookHandler
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

func (daoHandler *DaoHandler) Insert(data map[string]interface{}) *DaoHandler {
	daoHandler.model = daoHandler.model.Handler(daoHandler.dao.ParseInsert(data))
	return daoHandler
}

func (daoHandler *DaoHandler) Update(data map[string]interface{}) *DaoHandler {
	daoHandler.model = daoHandler.model.Handler(daoHandler.dao.ParseUpdate(data))
	return daoHandler
}

func (daoHandler *DaoHandler) HookUpdate(hookData map[string]interface{}, idArr ...int) *DaoHandler {
	daoHandler.model = daoHandler.model.Hook(daoHandler.dao.HookUpdate(hookData, idArr...))
	return daoHandler
}

func (daoHandler *DaoHandler) HookDelete(idArr ...int) *DaoHandler {
	daoHandler.model = daoHandler.model.Hook(daoHandler.dao.HookDelete(idArr...))
	return daoHandler
}

func (daoHandler *DaoHandler) Field(field []string) *DaoHandler {
	daoHandler.model = daoHandler.model.Handler(daoHandler.dao.ParseField(field, daoHandler.joinTableArr))
	return daoHandler
}

func (daoHandler *DaoHandler) Filter(filter map[string]interface{}) *DaoHandler {
	daoHandler.model = daoHandler.model.Handler(daoHandler.dao.ParseFilter(filter, daoHandler.joinTableArr))
	return daoHandler
}

func (daoHandler *DaoHandler) Group(group []string) *DaoHandler {
	daoHandler.model = daoHandler.model.Handler(daoHandler.dao.ParseGroup(group, daoHandler.joinTableArr))
	return daoHandler
}

func (daoHandler *DaoHandler) Order(order []string) *DaoHandler {
	daoHandler.model = daoHandler.model.Handler(daoHandler.dao.ParseOrder(order, daoHandler.joinTableArr))
	return daoHandler
}

func (daoHandler *DaoHandler) Join(joinCode string) *DaoHandler {
	daoHandler.model = daoHandler.model.Handler(daoHandler.dao.ParseJoin(joinCode, daoHandler.joinTableArr))
	return daoHandler
}

func (daoHandler *DaoHandler) GetModel() *gdb.Model {
	// return daoHandler.model.Clone()
	return daoHandler.model
}

// 判断是否联表
func (daoHandler *DaoHandler) IsJoin() bool {
	return len(*daoHandler.joinTableArr) > 0
}

// 联表则GroupBy主键
func (daoHandler *DaoHandler) JoinGroupByPrimaryKey() *DaoHandler {
	if daoHandler.IsJoin() {
		daoHandler.model = daoHandler.model.Group(daoHandler.dao.Table() + `.` + daoHandler.dao.PrimaryKey())
	}
	return daoHandler
}

// 总数（有联表默认group主键）
func (daoHandler *DaoHandler) Count() (count int, err error) {
	if daoHandler.IsJoin() {
		count, err = daoHandler.model.Clone().Group(daoHandler.dao.Table() + `.` + daoHandler.dao.PrimaryKey()).Distinct().Fields(daoHandler.dao.Table() + `.` + daoHandler.dao.PrimaryKey()).Count()
		return
	}
	count, err = daoHandler.model.Count()
	return
}
