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

// MenuDao is the data access object for the table auth_menu.
type MenuDao struct {
	table     string              // table is the underlying table name of the DAO.
	group     string              // group is the database configuration group name of the current DAO.
	columns   MenuColumns         // columns contains all the column names of Table for convenient usage.
	handlers  []gdb.ModelHandler  // handlers for customized model modification.
	columnArr []string            // 字段数组
	columnMap map[string]struct{} // 字段map
}

// MenuColumns defines and stores column names for the table auth_menu.
type MenuColumns struct {
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	IsStop    string // 停用：0否 1是
	MenuId    string // 菜单ID
	MenuName  string // 名称
	SceneId   string // 场景ID
	Pid       string // 父ID
	IsLeaf    string // 叶子：0否 1是
	Level     string // 层级
	IdPath    string // ID路径
	NamePath  string // 名称路径
	MenuIcon  string // 图标
	MenuUrl   string // 链接
	ExtraData string // 额外数据。JSON格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}
	Sort      string // 排序值。从大到小排序
}

// menuColumns holds the columns for the table auth_menu.
var menuColumns = MenuColumns{
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsStop:    "is_stop",
	MenuId:    "menu_id",
	MenuName:  "menu_name",
	SceneId:   "scene_id",
	Pid:       "pid",
	IsLeaf:    "is_leaf",
	Level:     "level",
	IdPath:    "id_path",
	NamePath:  "name_path",
	MenuIcon:  "menu_icon",
	MenuUrl:   "menu_url",
	ExtraData: "extra_data",
	Sort:      "sort",
}

// NewMenuDao creates and returns a new DAO object for table data access.
func NewMenuDao(handlers ...gdb.ModelHandler) *MenuDao {
	dao := &MenuDao{
		group:    "default",
		table:    "auth_menu",
		columns:  menuColumns,
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
func (dao *MenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *MenuDao) Columns() *MenuColumns {
	return &dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MenuDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 字段数组
func (dao *MenuDao) ColumnArr() []string {
	return dao.columnArr
}

// 字段map
func (dao *MenuDao) ColumnMap() map[string]struct{} {
	return dao.columnMap
}

// 判断字段是否存在
func (dao *MenuDao) Contains(column string) (ok bool) {
	_, ok = dao.columnMap[column]
	return
}
