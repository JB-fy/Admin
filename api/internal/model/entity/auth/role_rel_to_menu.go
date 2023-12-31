// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleRelToMenu is the golang structure for table role_rel_to_menu.
type RoleRelToMenu struct {
	RoleId    uint        `json:"roleId"    ` // 角色ID
	MenuId    uint        `json:"menuId"    ` // 菜单ID
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
}
