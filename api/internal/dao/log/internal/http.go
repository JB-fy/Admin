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

// HttpDao is the data access object for table log_http.
type HttpDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    HttpColumns      // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// HttpColumns defines and stores column names for table log_http.
type HttpColumns struct {
	HttpId    string // Http日志ID
	Url       string // 地址
	Header    string // 请求头
	ReqData   string // 请求数据
	ResData   string // 响应数据
	RunTime   string // 运行时间（单位：毫秒）
	UpdatedAt string // 更新时间
	CreatedAt string // 创建时间
}

// httpColumns holds the columns for table log_http.
var httpColumns = HttpColumns{
	HttpId:    "httpId",
	Url:       "url",
	Header:    "header",
	ReqData:   "reqData",
	ResData:   "resData",
	RunTime:   "runTime",
	UpdatedAt: "updatedAt",
	CreatedAt: "createdAt",
}

// NewHttpDao creates and returns a new DAO object for table data access.
func NewHttpDao() *HttpDao {
	return &HttpDao{
		group:   `default`,
		table:   `log_http`,
		columns: httpColumns,
		primaryKey: func() string {
			return reflect.ValueOf(httpColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(httpColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(httpColumns)
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
func (dao *HttpDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HttpDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HttpDao) Columns() HttpColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HttpDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HttpDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HttpDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *HttpDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *HttpDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *HttpDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
