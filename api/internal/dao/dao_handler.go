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
	Ctx          context.Context
	Dao          DaoInterface
	Model        *gdb.Model
	JoinTableArr *[]string
}

func NewDaoHandler(ctx context.Context, dao DaoInterface, dbSelDataList ...map[string]interface{}) *DaoHandler {
	daoHandlerObj := DaoHandler{
		Ctx:          ctx,
		Dao:          dao,
		JoinTableArr: &[]string{},
	}
	daoHandlerObj.Model = daoHandlerObj.Dao.ParseDbCtx(daoHandlerObj.Ctx, dbSelDataList...)
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

/* func (daoHandler *DaoHandler) GetModel(joinCode string) *gdb.Model {
	return daoHandler.Model
} */

func (daoHandler *DaoHandler) HandleInsert(data map[string]interface{}) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Handler(daoHandler.Dao.ParseInsert(data))
	return daoHandler
}

func (daoHandler *DaoHandler) HandleUpdate(data map[string]interface{}) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Handler(daoHandler.Dao.ParseUpdate(data))
	return daoHandler
}

func (daoHandler *DaoHandler) HandleHookUpdate(hookData map[string]interface{}, idArr ...int) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Hook(daoHandler.Dao.HookUpdate(hookData, idArr...))
	return daoHandler
}

func (daoHandler *DaoHandler) HandleHookDelete(idArr ...int) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Hook(daoHandler.Dao.HookDelete(idArr...))
	return daoHandler
}

func (daoHandler *DaoHandler) HandleField(field []string) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Handler(daoHandler.Dao.ParseField(field, daoHandler.JoinTableArr))
	return daoHandler
}

func (daoHandler *DaoHandler) HandleFilter(filter map[string]interface{}) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Handler(daoHandler.Dao.ParseFilter(filter, daoHandler.JoinTableArr))
	return daoHandler
}

func (daoHandler *DaoHandler) HandleGroup(group []string) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Handler(daoHandler.Dao.ParseGroup(group, daoHandler.JoinTableArr))
	return daoHandler
}

func (daoHandler *DaoHandler) HandleOrder(order []string) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Handler(daoHandler.Dao.ParseOrder(order, daoHandler.JoinTableArr))
	return daoHandler
}

func (daoHandler *DaoHandler) HandleJoin(joinCode string) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Handler(daoHandler.Dao.ParseJoin(joinCode, daoHandler.JoinTableArr))
	return daoHandler
}

// 分页处理
func (daoHandler *DaoHandler) Page(page int, limit int) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Page(page, limit)
	return daoHandler
}

// 总数（有联表默认group主键）
func (daoHandler *DaoHandler) Count() (count int, err error) {
	daoHandlerTmp := *(daoHandler.Model)
	if len(*daoHandler.JoinTableArr) > 0 {
		daoHandler.Model = daoHandler.Model.Group(daoHandler.Dao.Table() + `.` + daoHandler.Dao.PrimaryKey()).Distinct().Fields(daoHandler.Dao.Table() + `.` + daoHandler.Dao.PrimaryKey())
	}
	count, err = daoHandler.Model.Count()
	daoHandler.Model = &daoHandlerTmp //执行完还原model，以便继续调用List方法（连续调用Count和List方法时，Count有联表，List方法调用时会受影响，多出一个Distinct主键字段）
	return
}

// 列表（有联表默认group主键的）
func (daoHandler *DaoHandler) List() (list gdb.Result, err error) {
	if len(*daoHandler.JoinTableArr) > 0 {
		daoHandler.Model = daoHandler.Model.Group(daoHandler.Dao.Table() + `.` + daoHandler.Dao.PrimaryKey())
	}
	list, err = daoHandler.Model.All()
	return
}
