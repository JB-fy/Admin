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

// RoleRelOfOrgAdminDao is the data access object for table auth_role_rel_of_org_admin.
type RoleRelOfOrgAdminDao struct {
	table     string                   // table is the underlying table name of the DAO.
	group     string                   // group is the database configuration group name of current DAO.
	columns   RoleRelOfOrgAdminColumns // columns contains all the column names of Table for convenient usage.
	columnArr []string                 // 字段数组
	columnMap map[string]struct{}      // 字段map
}

// RoleRelOfOrgAdminColumns defines and stores column names for table auth_role_rel_of_org_admin.
type RoleRelOfOrgAdminColumns struct {
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	AdminId   string // 管理员ID
	RoleId    string // 角色ID
}

// roleRelOfOrgAdminColumns holds the columns for table auth_role_rel_of_org_admin.
var roleRelOfOrgAdminColumns = RoleRelOfOrgAdminColumns{
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	AdminId:   "admin_id",
	RoleId:    "role_id",
}

// NewRoleRelOfOrgAdminDao creates and returns a new DAO object for table data access.
func NewRoleRelOfOrgAdminDao() *RoleRelOfOrgAdminDao {
	dao := &RoleRelOfOrgAdminDao{
		group:   `default`,
		table:   `auth_role_rel_of_org_admin`,
		columns: roleRelOfOrgAdminColumns,
	}
	v := reflect.ValueOf(dao.columns)
	count := v.NumField()
	dao.columnArr = make([]string, count)
	dao.columnMap = make(map[string]struct{}, count)
	for i := 0; i < count; i++ {
		dao.columnArr[i] = v.Field(i).String()
		dao.columnMap[v.Field(i).String()] = struct{}{}
	}
	return dao
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoleRelOfOrgAdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoleRelOfOrgAdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *RoleRelOfOrgAdminDao) Columns() *RoleRelOfOrgAdminColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoleRelOfOrgAdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoleRelOfOrgAdminDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoleRelOfOrgAdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 字段数组
func (dao *RoleRelOfOrgAdminDao) ColumnArr() []string {
	return dao.columnArr
}

// 字段map
func (dao *RoleRelOfOrgAdminDao) ColumnMap() map[string]struct{} {
	return dao.columnMap
}

// 判断字段是否存在
func (dao *RoleRelOfOrgAdminDao) Contains(column string) (ok bool) {
	_, ok = dao.columnMap[column]
	return
}
