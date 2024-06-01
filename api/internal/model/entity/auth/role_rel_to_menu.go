// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleRelToMenu is the golang structure for table role_rel_to_menu.
type RoleRelToMenu struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	RoleId    uint        `json:"roleId"    orm:"role_id"    ` // 角色ID
	MenuId    uint        `json:"menuId"    orm:"menu_id"    ` // 菜单ID
}
