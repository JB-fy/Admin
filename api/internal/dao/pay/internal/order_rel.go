// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderRelDao is the data access object for the table pay_order_rel.
type OrderRelDao struct {
	table     string              // table is the underlying table name of the DAO.
	group     string              // group is the database configuration group name of the current DAO.
	columns   OrderRelColumns     // columns contains all the column names of Table for convenient usage.
	handlers  []gdb.ModelHandler  // handlers for customized model modification.
	columnArr []string            // 字段数组
	columnMap map[string]struct{} // 字段map
}

// OrderRelColumns defines and stores column names for the table pay_order_rel.
type OrderRelColumns struct {
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	OrderId        string // 订单ID
	RelOrderType   string // 关联订单类型：0默认
	RelOrderId     string // 关联订单ID
	RelOrderNo     string // 关联订单号
	RelOrderUserId string // 关联订单用户ID
	RelOrderAmount string // 关联订单实付金额
}

// orderRelColumns holds the columns for the table pay_order_rel.
var orderRelColumns = OrderRelColumns{
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	OrderId:        "order_id",
	RelOrderType:   "rel_order_type",
	RelOrderId:     "rel_order_id",
	RelOrderNo:     "rel_order_no",
	RelOrderUserId: "rel_order_user_id",
	RelOrderAmount: "rel_order_amount",
}

// NewOrderRelDao creates and returns a new DAO object for table data access.
func NewOrderRelDao(handlers ...gdb.ModelHandler) *OrderRelDao {
	dao := &OrderRelDao{
		group:    "default",
		table:    "pay_order_rel",
		columns:  orderRelColumns,
		handlers: handlers,
	}
	v := reflect.ValueOf(dao.columns)
	count := v.NumField()
	dao.columnArr = make([]string, count)
	dao.columnMap = make(map[string]struct{}, count)
	for i := range count {
		dao.columnArr[i] = v.Field(i).String()
		dao.columnMap[v.Field(i).String()] = struct{}{}
	}
	return dao
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrderRelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrderRelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *OrderRelDao) Columns() *OrderRelColumns {
	return &dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrderRelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrderRelDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OrderRelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 字段数组
func (dao *OrderRelDao) ColumnArr() []string {
	return dao.columnArr
}

// 字段map
func (dao *OrderRelDao) ColumnMap() map[string]struct{} {
	return dao.columnMap
}

// 判断字段是否存在
func (dao *OrderRelDao) Contains(column string) (ok bool) {
	_, ok = dao.columnMap[column]
	return
}
