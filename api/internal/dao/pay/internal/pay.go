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

// PayDao is the data access object for table pay.
type PayDao struct {
	table     string           // table is the underlying table name of the DAO.
	group     string           // group is the database configuration group name of current DAO.
	columns   PayColumns       // columns contains all the column names of Table for convenient usage.
	columnArr *garray.StrArray // 所有字段的数组
}

// PayColumns defines and stores column names for table pay.
type PayColumns struct {
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	IsStop      string // 停用：0否 1是
	PayId       string // 支付ID
	PayName     string // 名称
	PayIcon     string // 图标
	PayType     string // 类型：0支付宝 1微信
	PayConfig   string // 配置。根据pay_type类型设置
	PayRate     string // 费率
	TotalAmount string // 总额
	Balance     string // 余额
	Sort        string // 排序值。从大到小排序
}

// payColumns holds the columns for table pay.
var payColumns = PayColumns{
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	IsStop:      "is_stop",
	PayId:       "pay_id",
	PayName:     "pay_name",
	PayIcon:     "pay_icon",
	PayType:     "pay_type",
	PayConfig:   "pay_config",
	PayRate:     "pay_rate",
	TotalAmount: "total_amount",
	Balance:     "balance",
	Sort:        "sort",
}

// NewPayDao creates and returns a new DAO object for table data access.
func NewPayDao() *PayDao {
	return &PayDao{
		group:   `default`,
		table:   `pay`,
		columns: payColumns,
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(payColumns)
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
func (dao *PayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PayDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *PayDao) Columns() *PayColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PayDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PayDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 所有字段的数组
func (dao *PayDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
