// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// internalRoleRelToMenuDao is internal type for wrapping internal DAO implements.
type internalRoleRelToMenuDao = *internal.RoleRelToMenuDao

// roleRelToMenuDao is the data access object for table auth_role_rel_to_menu.
// You can define custom methods on it to extend its functionality as you wish.
type roleRelToMenuDao struct {
	internalRoleRelToMenuDao
}

var (
	// RoleRelToMenu is globally public accessible object for table auth_role_rel_to_menu operations.
	RoleRelToMenu = roleRelToMenuDao{
		internal.NewRoleRelToMenuDao(),
	}
)

// 解析分库
func (daoThis *roleRelToMenuDao) ParseDbGroup(dbGroupSeldata map[string]interface{}) string {
	group := daoThis.Group()
	/* if len(dbGroupSeldata) > 0 { //分库逻辑
	} */
	return group
}

// 解析分表
func (daoThis *roleRelToMenuDao) ParseDbTable(dbTableSelData map[string]interface{}) string {
	table := daoThis.Table()
	/* if len(dbTableSelData) > 0 { //分表逻辑
	} */
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *roleRelToMenuDao) ParseDbCtx(ctx context.Context, dbSelDataList ...map[string]interface{}) *gdb.Model {
	switch len(dbSelDataList) {
	case 1:
		return g.DB(daoThis.ParseDbGroup(dbSelDataList[0])).Model(daoThis.Table()).Safe().Ctx(ctx)
	case 2:
		return g.DB(daoThis.ParseDbGroup(dbSelDataList[0])).Model(daoThis.ParseDbTable(dbSelDataList[1])).Safe().Ctx(ctx)
	default:
		return daoThis.Ctx(ctx)
	}
}

// Fill with you ideas below.
