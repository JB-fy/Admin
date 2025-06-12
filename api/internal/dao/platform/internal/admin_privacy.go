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

// AdminPrivacyDao is the data access object for the table platform_admin_privacy.
type AdminPrivacyDao struct {
	table     string              // table is the underlying table name of the DAO.
	group     string              // group is the database configuration group name of the current DAO.
	columns   AdminPrivacyColumns // columns contains all the column names of Table for convenient usage.
	handlers  []gdb.ModelHandler  // handlers for customized model modification.
	columnArr []string            // 字段数组
	columnMap map[string]struct{} // 字段map
}

// AdminPrivacyColumns defines and stores column names for the table platform_admin_privacy.
type AdminPrivacyColumns struct {
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	AdminId   string // 管理员ID
	Password  string // 密码。md5保存
	Salt      string // 密码盐
}

// adminPrivacyColumns holds the columns for the table platform_admin_privacy.
var adminPrivacyColumns = AdminPrivacyColumns{
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	AdminId:   "admin_id",
	Password:  "password",
	Salt:      "salt",
}

// NewAdminPrivacyDao creates and returns a new DAO object for table data access.
func NewAdminPrivacyDao(handlers ...gdb.ModelHandler) *AdminPrivacyDao {
	dao := &AdminPrivacyDao{
		group:    "default",
		table:    "platform_admin_privacy",
		columns:  adminPrivacyColumns,
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
func (dao *AdminPrivacyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminPrivacyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *AdminPrivacyDao) Columns() *AdminPrivacyColumns {
	return &dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminPrivacyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminPrivacyDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminPrivacyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 字段数组
func (dao *AdminPrivacyDao) ColumnArr() []string {
	return dao.columnArr
}

// 字段map
func (dao *AdminPrivacyDao) ColumnMap() map[string]struct{} {
	return dao.columnMap
}

// 判断字段是否存在
func (dao *AdminPrivacyDao) Contains(column string) (ok bool) {
	_, ok = dao.columnMap[column]
	return
}
