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

// UploadDao is the data access object for table upload.
type UploadDao struct {
	table     string           // table is the underlying table name of the DAO.
	group     string           // group is the database configuration group name of current DAO.
	columns   UploadColumns    // columns contains all the column names of Table for convenient usage.
	columnArr *garray.StrArray // 所有字段的数组
}

// UploadColumns defines and stores column names for table upload.
type UploadColumns struct {
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	IsStop       string // 停用：0否 1是
	UploadId     string // 上传ID
	UploadType   string // 类型：0本地 1阿里云OSS
	UploadConfig string // 配置。根据upload_type类型设置
	IsDefault    string // 默认：0否 1是
	Remark       string // 备注
}

// uploadColumns holds the columns for table upload.
var uploadColumns = UploadColumns{
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	IsStop:       "is_stop",
	UploadId:     "upload_id",
	UploadType:   "upload_type",
	UploadConfig: "upload_config",
	IsDefault:    "is_default",
	Remark:       "remark",
}

// NewUploadDao creates and returns a new DAO object for table data access.
func NewUploadDao() *UploadDao {
	return &UploadDao{
		group:   `default`,
		table:   `upload`,
		columns: uploadColumns,
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(uploadColumns)
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
func (dao *UploadDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UploadDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *UploadDao) Columns() *UploadColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UploadDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UploadDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UploadDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 所有字段的数组
func (dao *UploadDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
