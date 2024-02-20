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

// RoleRelOfPlatformAdminDao is the data access object for table auth_role_rel_of_platform_admin.
type RoleRelOfPlatformAdminDao struct {
	table      string                        // table is the underlying table name of the DAO.
	group      string                        // group is the database configuration group name of current DAO.
	columns    RoleRelOfPlatformAdminColumns // columns contains all the column names of Table for convenient usage.
	primaryKey string                        // 主键ID
	columnArr  *garray.StrArray              // 所有字段的数组
}

// RoleRelOfPlatformAdminColumns defines and stores column names for table auth_role_rel_of_platform_admin.
type RoleRelOfPlatformAdminColumns struct {
	RoleId    string // 角色ID
	AdminId   string // 管理员ID
	UpdatedAt string // 更新时间
	CreatedAt string // 创建时间
}

// roleRelOfPlatformAdminColumns holds the columns for table auth_role_rel_of_platform_admin.
var roleRelOfPlatformAdminColumns = RoleRelOfPlatformAdminColumns{
	RoleId:    "roleId",
	AdminId:   "adminId",
	UpdatedAt: "updatedAt",
	CreatedAt: "createdAt",
}

// NewRoleRelOfPlatformAdminDao creates and returns a new DAO object for table data access.
func NewRoleRelOfPlatformAdminDao() *RoleRelOfPlatformAdminDao {
	return &RoleRelOfPlatformAdminDao{
		group:   `default`,
		table:   `auth_role_rel_of_platform_admin`,
		columns: roleRelOfPlatformAdminColumns,
		primaryKey: func() string {
			return reflect.ValueOf(roleRelOfPlatformAdminColumns).Field(0).String()
		}(),
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(roleRelOfPlatformAdminColumns)
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
func (dao *RoleRelOfPlatformAdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoleRelOfPlatformAdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *RoleRelOfPlatformAdminDao) Columns() *RoleRelOfPlatformAdminColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoleRelOfPlatformAdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoleRelOfPlatformAdminDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoleRelOfPlatformAdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *RoleRelOfPlatformAdminDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *RoleRelOfPlatformAdminDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
