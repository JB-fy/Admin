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

// AppDao is the data access object for the table app.
type AppDao struct {
	table     string              // table is the underlying table name of the DAO.
	group     string              // group is the database configuration group name of the current DAO.
	columns   AppColumns          // columns contains all the column names of Table for convenient usage.
	handlers  []gdb.ModelHandler  // handlers for customized model modification.
	columnArr []string            // 字段数组
	columnMap map[string]struct{} // 字段map
}

// AppColumns defines and stores column names for the table app.
type AppColumns struct {
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	IsStop      string // 停用：0否 1是
	AppId       string // APPID
	NameType    string // 名称：0APP。有两种以上APP时自行扩展
	AppType     string // 类型：0安卓 1苹果 2PC
	PackageName string // 包名
	PackageFile string // 安装包
	VerNo       string // 版本号
	VerName     string // 版本名称
	VerIntro    string // 版本介绍
	ExtraConfig string // 额外配置
	Remark      string // 备注
	IsForcePrev string // 强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关
}

// appColumns holds the columns for the table app.
var appColumns = AppColumns{
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	IsStop:      "is_stop",
	AppId:       "app_id",
	NameType:    "name_type",
	AppType:     "app_type",
	PackageName: "package_name",
	PackageFile: "package_file",
	VerNo:       "ver_no",
	VerName:     "ver_name",
	VerIntro:    "ver_intro",
	ExtraConfig: "extra_config",
	Remark:      "remark",
	IsForcePrev: "is_force_prev",
}

// NewAppDao creates and returns a new DAO object for table data access.
func NewAppDao(handlers ...gdb.ModelHandler) *AppDao {
	dao := &AppDao{
		group:    "default",
		table:    "app",
		columns:  appColumns,
		handlers: handlers,
	}
	v := reflect.ValueOf(dao.columns)
	count := v.NumField()
	dao.columnArr = make([]string, count)
	dao.columnMap = make(map[string]struct{}, count)
	for i := range count {
		dao.columnArr[i] = v.Field(i).String()
		dao.columnMap[v.Field(i).String()] = struct{}{}
	}
	return dao
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *AppDao) Columns() *AppColumns {
	return &dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AppDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 字段数组
func (dao *AppDao) ColumnArr() []string {
	return dao.columnArr
}

// 字段map
func (dao *AppDao) ColumnMap() map[string]struct{} {
	return dao.columnMap
}

// 判断字段是否存在
func (dao *AppDao) Contains(column string) (ok bool) {
	_, ok = dao.columnMap[column]
	return
}
