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

// ActionRelToSceneDao is the data access object for table auth_action_rel_to_scene.
type ActionRelToSceneDao struct {
	table     string                  // table is the underlying table name of the DAO.
	group     string                  // group is the database configuration group name of current DAO.
	columns   ActionRelToSceneColumns // columns contains all the column names of Table for convenient usage.
	columnArr []string                // 字段数组
	columnMap map[string]struct{}     // 字段map
}

// ActionRelToSceneColumns defines and stores column names for table auth_action_rel_to_scene.
type ActionRelToSceneColumns struct {
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	ActionId  string // 操作ID
	SceneId   string // 场景ID
}

// actionRelToSceneColumns holds the columns for table auth_action_rel_to_scene.
var actionRelToSceneColumns = ActionRelToSceneColumns{
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	ActionId:  "action_id",
	SceneId:   "scene_id",
}

// NewActionRelToSceneDao creates and returns a new DAO object for table data access.
func NewActionRelToSceneDao() *ActionRelToSceneDao {
	dao := &ActionRelToSceneDao{
		group:   `default`,
		table:   `auth_action_rel_to_scene`,
		columns: actionRelToSceneColumns,
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
func (dao *ActionRelToSceneDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActionRelToSceneDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
// 使用较为频繁。为优化内存考虑，改成返回指针更为合适，但切忌使用过程中不可修改，否则会污染全局
func (dao *ActionRelToSceneDao) Columns() *ActionRelToSceneColumns {
	return &dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActionRelToSceneDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActionRelToSceneDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActionRelToSceneDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 字段数组
func (dao *ActionRelToSceneDao) ColumnArr() []string {
	return dao.columnArr
}

// 字段map
func (dao *ActionRelToSceneDao) ColumnMap() map[string]struct{} {
	return dao.columnMap
}

// 判断字段是否存在
func (dao *ActionRelToSceneDao) Contains(column string) (ok bool) {
	_, ok = dao.columnMap[column]
	return
}
