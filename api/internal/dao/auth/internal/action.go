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

// ActionDao is the data access object for table auth_action.
type ActionDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    ActionColumns    // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// ActionColumns defines and stores column names for table auth_action.
type ActionColumns struct {
	ActionId   string // 权限操作ID
	ActionName string // 名称
	ActionCode string // 标识（代码中用于判断权限）
	Remark     string // 备注
	IsStop     string // 停用：0否 1是
	UpdatedAt  string // 更新时间
	CreatedAt  string // 创建时间
}

// actionColumns holds the columns for table auth_action.
var actionColumns = ActionColumns{
	ActionId:   "actionId",
	ActionName: "actionName",
	ActionCode: "actionCode",
	Remark:     "remark",
	IsStop:     "isStop",
	UpdatedAt:  "updatedAt",
	CreatedAt:  "createdAt",
}

// NewActionDao creates and returns a new DAO object for table data access.
func NewActionDao() *ActionDao {
	return &ActionDao{
		group:   `default`,
		table:   `auth_action`,
		columns: actionColumns,
		primaryKey: func() string {
			return reflect.ValueOf(actionColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(actionColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(actionColumns)
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
func (dao *ActionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActionDao) Columns() ActionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *ActionDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *ActionDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *ActionDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
