// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoleRelToActionDao is the data access object for table auth_role_rel_to_action.
type RoleRelToActionDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns RoleRelToActionColumns // columns contains all the column names of Table for convenient usage.
}

// RoleRelToActionColumns defines and stores column names for table auth_role_rel_to_action.
type RoleRelToActionColumns struct {
	RoleId    string // 权限角色ID
	ActionId  string // 权限操作ID
	UpdatedAt string // 更新时间
	CreatedAt string // 创建时间
}

// roleRelToActionColumns holds the columns for table auth_role_rel_to_action.
var roleRelToActionColumns = RoleRelToActionColumns{
	RoleId:    "roleId",
	ActionId:  "actionId",
	UpdatedAt: "updatedAt",
	CreatedAt: "createdAt",
}

// NewRoleRelToActionDao creates and returns a new DAO object for table data access.
func NewRoleRelToActionDao() *RoleRelToActionDao {
	return &RoleRelToActionDao{
		group:   "default",
		table:   "auth_role_rel_to_action",
		columns: roleRelToActionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoleRelToActionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoleRelToActionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoleRelToActionDao) Columns() RoleRelToActionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoleRelToActionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoleRelToActionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoleRelToActionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
