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

// ServerDao is the data access object for table platform_server.
type ServerDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    ServerColumns    // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  *garray.StrArray // 所有字段的数组
}

// ServerColumns defines and stores column names for table platform_server.
type ServerColumns struct {
	ServerId  string // 服务器ID
	NetworkIp string // 公网IP
	LocalIp   string // 内网IP
	UpdatedAt string // 更新时间
	CreatedAt string // 创建时间
}

// serverColumns holds the columns for table platform_server.
var serverColumns = ServerColumns{
	ServerId:  "serverId",
	NetworkIp: "networkIp",
	LocalIp:   "localIp",
	UpdatedAt: "updatedAt",
	CreatedAt: "createdAt",
}

// NewServerDao creates and returns a new DAO object for table data access.
func NewServerDao() *ServerDao {
	return &ServerDao{
		group:   `default`,
		table:   `platform_server`,
		columns: serverColumns,
		primaryKey: func() string {
			return reflect.ValueOf(serverColumns).Field(0).String()
		}(),
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(serverColumns)
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
func (dao *ServerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ServerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *ServerDao) Columns() *ServerColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ServerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ServerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ServerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *ServerDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *ServerDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
