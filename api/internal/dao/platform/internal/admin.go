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

// AdminDao is the data access object for table platform_admin.
type AdminDao struct {
	table     string           // table is the underlying table name of the DAO.
	group     string           // group is the database configuration group name of current DAO.
	columns   AdminColumns     // columns contains all the column names of Table for convenient usage.
	columnArr *garray.StrArray // 所有字段的数组
}

// AdminColumns defines and stores column names for table platform_admin.
type AdminColumns struct {
	AdminId   string // 管理员ID
	Phone     string // 手机
	Account   string // 账号
	Password  string // 密码。md5保存
	Salt      string // 密码盐
	Nickname  string // 昵称
	Avatar    string // 头像
	IsStop    string // 停用：0否 1是
	UpdatedAt string // 更新时间
	CreatedAt string // 创建时间
}

// adminColumns holds the columns for table platform_admin.
var adminColumns = AdminColumns{
	AdminId:   "admin_id",
	Phone:     "phone",
	Account:   "account",
	Password:  "password",
	Salt:      "salt",
	Nickname:  "nickname",
	Avatar:    "avatar",
	IsStop:    "is_stop",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

// NewAdminDao creates and returns a new DAO object for table data access.
func NewAdminDao() *AdminDao {
	return &AdminDao{
		group:   `default`,
		table:   `platform_admin`,
		columns: adminColumns,
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(adminColumns)
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
func (dao *AdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *AdminDao) Columns() *AdminColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 所有字段的数组
func (dao *AdminDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
