package utils

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

type Query struct {
	Ctx          context.Context
	Dao          DaoInterface
	Model        *gdb.Model
	JoinTableArr *[]string
}

func NewQuery(ctx context.Context, dao DaoInterface) *Query {
	queryObj := Query{
		Ctx:          ctx,
		Dao:          dao,
		JoinTableArr: &[]string{},
	}
	queryObj.Model = queryObj.Dao.ParseDbCtx(ctx)
	return &queryObj
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

// func (logicThis *Query) List(ctx context.Context, daoThis Dao, filter map[string]interface{}, field []string, order []string) (list gdb.Result, err error) {
func (query *Query) Filter(filter map[string]interface{}) *Query {
	query.Model = query.Model.Handler(query.Dao.ParseFilter(filter, query.JoinTableArr))
	return query
}

func (query *Query) Field(field []string) *Query {
	query.Model = query.Model.Handler(query.Dao.ParseField(field, query.JoinTableArr))
	return query
}

func (query *Query) Order(order []string) *Query {
	query.Model = query.Model.Handler(query.Dao.ParseOrder(order, query.JoinTableArr))
	return query
}

func (query *Query) Count() (count int, err error) {
	// queryTmp := *(query.Model)
	if len(*query.JoinTableArr) == 0 {
		query.Model = query.Model.Group(query.Dao.Table() + `.` + query.Dao.PrimaryKey()).Distinct().Fields(query.Dao.Table() + `.` + query.Dao.PrimaryKey())
	}
	count, err = query.Model.Count()
	// query.Model = &queryTmp
	return
}

func (query *Query) List(page int, limit int) (list gdb.Result, err error) {
	if len(*query.JoinTableArr) > 0 {
		query.Model = query.Model.Group(query.Dao.Table() + `.` + query.Dao.PrimaryKey())
	}
	if limit > 0 {
		query.Model = query.Model.Offset((page - 1) * limit).Limit(limit)
	}
	list, err = query.Model.All()
	return
}
