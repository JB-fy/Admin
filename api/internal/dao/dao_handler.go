package dao

/* query := utils.NewQuery(ctx, &daoAuth.Scene)
count, err := query.Filter(filter).Count()
// count, err := service.AuthScene().Count(ctx, filter)
if err != nil {
	return
}
list, err := query.Field(field).Order(order).List(page, limit)
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

func NewDaoHandler(ctx context.Context, dao DaoInterface) *DaoHandler {
	daoHandlerObj := DaoHandler{
		Ctx:          ctx,
		Dao:          dao,
		JoinTableArr: &[]string{},
	}
	daoHandlerObj.Model = daoHandlerObj.Dao.ParseDbCtx(ctx)
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

func (daoHandler *DaoHandler) HandleFilter(filter map[string]interface{}) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Handler(daoHandler.Dao.ParseFilter(filter, daoHandler.JoinTableArr))
	return daoHandler
}

func (daoHandler *DaoHandler) HandleField(field []string) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Handler(daoHandler.Dao.ParseField(field, daoHandler.JoinTableArr))
	return daoHandler
}

func (daoHandler *DaoHandler) HandleOrder(order []string) *DaoHandler {
	daoHandler.Model = daoHandler.Model.Handler(daoHandler.Dao.ParseOrder(order, daoHandler.JoinTableArr))
	return daoHandler
}

func (daoHandler *DaoHandler) Count() (count int, err error) {
	// daoHandlerTmp := *(daoHandler.Model)
	if len(*daoHandler.JoinTableArr) > 0 {
		daoHandler.Model = daoHandler.Model.Group(daoHandler.Dao.Table() + `.` + daoHandler.Dao.PrimaryKey()).Distinct().Fields(daoHandler.Dao.Table() + `.` + daoHandler.Dao.PrimaryKey())
	}
	count, err = daoHandler.Model.Count()
	// daoHandler.Model = &daoHandlerTmp	//如果连续调用Count和List方法时,Count有联表,List方法调用时会受影响(多出一个Distinct主键字段)
	return
}

func (daoHandler *DaoHandler) List(page int, limit int) (list gdb.Result, err error) {
	if len(*daoHandler.JoinTableArr) > 0 {
		daoHandler.Model = daoHandler.Model.Group(daoHandler.Dao.Table() + `.` + daoHandler.Dao.PrimaryKey())
	}
	if limit > 0 {
		daoHandler.Model = daoHandler.Model.Offset((page - 1) * limit).Limit(limit)
	}
	list, err = daoHandler.Model.All()
	return
}
