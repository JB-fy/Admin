// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleRelToAction is the golang structure of table auth_role_rel_to_action for DAO operations like Where/Data.
type RoleRelToAction struct {
	g.Meta    `orm:"table:auth_role_rel_to_action, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	RoleId    any         // 角色ID
	ActionId  any         // 操作ID
}
