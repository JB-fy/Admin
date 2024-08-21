// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderDao is the data access object for table pay_order.
type OrderDao struct {
	table     string           // table is the underlying table name of the DAO.
	group     string           // group is the database configuration group name of current DAO.
	columns   OrderColumns     // columns contains all the column names of Table for convenient usage.
	columnArr *garray.StrArray // 所有字段的数组
}

// OrderColumns defines and stores column names for table pay_order.
type OrderColumns struct {
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	OrderId      string // 订单ID
	OrderNo      string // 订单号
	PayId        string // 支付ID
	ChannelId    string // 通道ID
	PayType      string // 类型：0支付宝 1微信
	PayPrice     string // 实付金额
	PayStatus    string // 状态：0未付款 1已付款
	PayTime      string // 支付时间
	PayRate      string // 费率
	ThirdOrderNo string // 第三方订单号
}

// orderColumns holds the columns for table pay_order.
var orderColumns = OrderColumns{
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	OrderId:      "order_id",
	OrderNo:      "order_no",
	PayId:        "pay_id",
	ChannelId:    "channel_id",
	PayType:      "pay_type",
	PayPrice:     "pay_price",
	PayStatus:    "pay_status",
	PayTime:      "pay_time",
	PayRate:      "pay_rate",
	ThirdOrderNo: "third_order_no",
}

// NewOrderDao creates and returns a new DAO object for table data access.
func NewOrderDao() *OrderDao {
	return &OrderDao{
		group:   `default`,
		table:   `pay_order`,
		columns: orderColumns,
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(orderColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return garray.NewStrArrayFrom(column)
		}(),
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *OrderDao) Columns() *OrderColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 所有字段的数组
func (dao *OrderDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
