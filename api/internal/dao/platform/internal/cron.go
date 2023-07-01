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

// CronDao is the data access object for table platform_cron.
type CronDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    CronColumns      // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// CronColumns defines and stores column names for table platform_cron.
type CronColumns struct {
	CronId      string // 定时器ID
	CronName    string // 名称
	CronCode    string // 标识
	CronPattern string // 表达式
	Remark      string // 备注
	IsStop      string // 是否停用：0否 1是
	UpdatedAt   string // 更新时间
	CreatedAt   string // 创建时间
}

// cronColumns holds the columns for table platform_cron.
var cronColumns = CronColumns{
	CronId:      "cronId",
	CronName:    "cronName",
	CronCode:    "cronCode",
	CronPattern: "cronPattern",
	Remark:      "remark",
	IsStop:      "isStop",
	UpdatedAt:   "updatedAt",
	CreatedAt:   "createdAt",
}

// NewCronDao creates and returns a new DAO object for table data access.
func NewCronDao() *CronDao {
	return &CronDao{
		group:   `default`,
		table:   `platform_cron`,
		columns: cronColumns,
		primaryKey: func() string {
			return reflect.ValueOf(cronColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(cronColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(cronColumns)
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
func (dao *CronDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CronDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CronDao) Columns() CronColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CronDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CronDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CronDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *CronDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *CronDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *CronDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
