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

// MenuDao is the data access object for table auth_menu.
type MenuDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    MenuColumns      // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  *garray.StrArray // 所有字段的数组
}

// MenuColumns defines and stores column names for table auth_menu.
type MenuColumns struct {
	MenuId    string // 菜单ID
	MenuName  string // 名称
	SceneId   string // 场景ID
	Pid       string // 父ID
	Level     string // 层级
	IdPath    string // 层级路径
	MenuIcon  string // 图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}
	MenuUrl   string // 链接
	ExtraData string // 额外数据。JSON格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}
	Sort      string // 排序值。从小到大排序，默认50，范围0-100
	IsStop    string // 停用：0否 1是
	UpdatedAt string // 更新时间
	CreatedAt string // 创建时间
}

// menuColumns holds the columns for table auth_menu.
var menuColumns = MenuColumns{
	MenuId:    "menu_id",
	MenuName:  "menu_name",
	SceneId:   "scene_id",
	Pid:       "pid",
	Level:     "level",
	IdPath:    "id_path",
	MenuIcon:  "menu_icon",
	MenuUrl:   "menu_url",
	ExtraData: "extra_data",
	Sort:      "sort",
	IsStop:    "is_stop",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

// NewMenuDao creates and returns a new DAO object for table data access.
func NewMenuDao() *MenuDao {
	return &MenuDao{
		group:   `default`,
		table:   `auth_menu`,
		columns: menuColumns,
		primaryKey: func() string {
			return reflect.ValueOf(menuColumns).Field(0).String()
		}(),
		columnArr: func() *garray.StrArray {
			v := reflect.ValueOf(menuColumns)
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
func (dao *MenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *MenuDao) Columns() *MenuColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *MenuDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *MenuDao) ColumnArr() *garray.StrArray {
	return dao.columnArr
}
