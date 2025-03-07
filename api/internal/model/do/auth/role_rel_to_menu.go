// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleRelToMenu is the golang structure of table auth_role_rel_to_menu for DAO operations like Where/Data.
type RoleRelToMenu struct {
	g.Meta    `orm:"table:auth_role_rel_to_menu, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	RoleId    interface{} // 角色ID
	MenuId    interface{} // 菜单ID
}
