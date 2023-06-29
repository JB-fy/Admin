// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoleRelToMenuDao is the data access object for table auth_role_rel_to_menu.
type RoleRelToMenuDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns RoleRelToMenuColumns // columns contains all the column names of Table for convenient usage.
}

// RoleRelToMenuColumns defines and stores column names for table auth_role_rel_to_menu.
type RoleRelToMenuColumns struct {
	RoleId    string // 权限角色ID
	MenuId    string // 权限菜单ID
	UpdatedAt string // 更新时间
	CreatedAt string // 创建时间
}

// roleRelToMenuColumns holds the columns for table auth_role_rel_to_menu.
var roleRelToMenuColumns = RoleRelToMenuColumns{
	RoleId:    "roleId",
	MenuId:    "menuId",
	UpdatedAt: "updatedAt",
	CreatedAt: "createdAt",
}

// NewRoleRelToMenuDao creates and returns a new DAO object for table data access.
func NewRoleRelToMenuDao() *RoleRelToMenuDao {
	return &RoleRelToMenuDao{
		group:   "default",
		table:   "auth_role_rel_to_menu",
		columns: roleRelToMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoleRelToMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoleRelToMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoleRelToMenuDao) Columns() RoleRelToMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoleRelToMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoleRelToMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoleRelToMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
