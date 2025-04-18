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

// OrderDao is the data access object for the table pay_order.
type OrderDao struct {
	table     string              // table is the underlying table name of the DAO.
	group     string              // group is the database configuration group name of the current DAO.
	columns   OrderColumns        // columns contains all the column names of Table for convenient usage.
	handlers  []gdb.ModelHandler  // handlers for customized model modification.
	columnArr []string            // 字段数组
	columnMap map[string]struct{} // 字段map
}

// OrderColumns defines and stores column names for the table pay_order.
type OrderColumns struct {
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	OrderId        string // 订单ID
	OrderNo        string // 订单号
	RelOrderType   string // 关联订单类型：0默认
	RelOrderUserId string // 关联订单用户ID
	PayId          string // 支付ID
	ChannelId      string // 通道ID
	PayType        string // 类型：0支付宝 1微信
	Amount         string // 实付金额
	PayStatus      string // 状态：0未付款 1已付款
	PayTime        string // 支付时间
	PayRate        string // 费率
	ThirdOrderNo   string // 第三方订单号
}

// orderColumns holds the columns for the table pay_order.
var orderColumns = OrderColumns{
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	OrderId:        "order_id",
	OrderNo:        "order_no",
	RelOrderType:   "rel_order_type",
	RelOrderUserId: "rel_order_user_id",
	PayId:          "pay_id",
	ChannelId:      "channel_id",
	PayType:        "pay_type",
	Amount:         "amount",
	PayStatus:      "pay_status",
	PayTime:        "pay_time",
	PayRate:        "pay_rate",
	ThirdOrderNo:   "third_order_no",
}

// NewOrderDao creates and returns a new DAO object for table data access.
func NewOrderDao(handlers ...gdb.ModelHandler) *OrderDao {
	dao := &OrderDao{
		group:    "default",
		table:    "pay_order",
		columns:  orderColumns,
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
func (dao *OrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *OrderDao) Columns() *OrderColumns {
	return &dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 字段数组
func (dao *OrderDao) ColumnArr() []string {
	return dao.columnArr
}

// 字段map
func (dao *OrderDao) ColumnMap() map[string]struct{} {
	return dao.columnMap
}

// 判断字段是否存在
func (dao *OrderDao) Contains(column string) (ok bool) {
	_, ok = dao.columnMap[column]
	return
}
