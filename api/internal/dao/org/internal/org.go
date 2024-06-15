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

// OrgDao is the data access object for table org.
type OrgDao struct {
	table     string           // table is the underlying table name of the DAO.
	group     string           // group is the database configuration group name of current DAO.
	columns   OrgColumns       // columns contains all the column names of Table for convenient usage.
	columnArr *garray.StrArray // 所有字段的数组
}

// OrgColumns defines and stores column names for table org.
type OrgColumns struct {
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	IsStop    string // 停用：0否 1是
	OrgId     string // 机构ID
	OrgName   string // 机构名称
}

// orgColumns holds the columns for table org.
var orgColumns = OrgColumns{
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsStop:    "is_stop",
	OrgId:     "org_id",
	OrgName:   "org_name",
}

// NewOrgDao creates and returns a new DAO object for table data access.
func NewOrgDao() *OrgDao {
	return &OrgDao{
		group:   `default`,
		table:   `org`,
		columns: orgColumns,
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(orgColumns)
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
func (dao *OrgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *OrgDao) Columns() *OrgColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 所有字段的数组
func (dao *OrgDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
