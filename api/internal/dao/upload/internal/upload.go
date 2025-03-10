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

// UploadDao is the data access object for table upload.
type UploadDao struct {
	table     string              // table is the underlying table name of the DAO.
	group     string              // group is the database configuration group name of current DAO.
	columns   UploadColumns       // columns contains all the column names of Table for convenient usage.
	columnArr []string            // 字段数组
	columnMap map[string]struct{} // 字段map
}

// UploadColumns defines and stores column names for table upload.
type UploadColumns struct {
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
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
	UploadId:     "upload_id",
	UploadType:   "upload_type",
	UploadConfig: "upload_config",
	IsDefault:    "is_default",
	Remark:       "remark",
}

// NewUploadDao creates and returns a new DAO object for table data access.
func NewUploadDao() *UploadDao {
	dao := &UploadDao{
		group:   `default`,
		table:   `upload`,
		columns: uploadColumns,
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

// 字段数组
func (dao *UploadDao) ColumnArr() []string {
	return append([]string{}, dao.columnArr...) 
}

// 字段map
func (dao *UploadDao) ColumnMap() map[string]struct{} {
	return dao.columnMap
}

// 判断字段是否存在
func (dao *UploadDao) Contains(column string) (ok bool) {
	_, ok = dao.columnMap[column]
	return
}
