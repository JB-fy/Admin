// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
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
	table      string      // table is the underlying table name of the DAO.
	group      string      // group is the database configuration group name of current DAO.
	columns    MenuColumns // columns contains all the column names of Table for convenient usage.
	primaryKey string
	columnArr  []string
	columnArrG *garray.StrArray
}

// MenuColumns defines and stores column names for table auth_menu.
type MenuColumns struct {
	MenuId     string // 权限菜单ID
	SceneId    string // 权限场景ID（只能是auth_scene表中sceneType为0的菜单类型场景）
	Pid        string // 父ID
	MenuName   string // 名称
	MenuIcon   string // 图标
	MenuUrl    string // 链接
	Level      string // 层级
	PidPath    string // 层级路径
	ExtraData  string // 额外数据。（json格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}）
	Sort       string // 排序值（从小到大排序，默认50，范围0-100）
	IsStop     string // 是否停用：0否 1是
	UpdateTime string // 更新时间
	CreateTime string // 创建时间
}

// menuColumns holds the columns for table auth_menu.
var menuColumns = MenuColumns{
	MenuId:     "menuId",
	SceneId:    "sceneId",
	Pid:        "pid",
	MenuName:   "menuName",
	MenuIcon:   "menuIcon",
	MenuUrl:    "menuUrl",
	Level:      "level",
	PidPath:    "pidPath",
	ExtraData:  "extraData",
	Sort:       "sort",
	IsStop:     "isStop",
	UpdateTime: "updateTime",
	CreateTime: "createTime",
}

// NewMenuDao creates and returns a new DAO object for table data access.
func NewMenuDao() *MenuDao {
	return &MenuDao{
		group:   "default",
		table:   "auth_menu",
		columns: menuColumns,
		primaryKey: func() string {
			return reflect.ValueOf(menuColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(menuColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
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
func (dao *MenuDao) Columns() MenuColumns {
	return dao.columns
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
func (dao *MenuDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（返回的格式更方便使用）
func (dao *MenuDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
