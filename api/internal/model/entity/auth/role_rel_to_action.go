// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleRelToAction is the golang structure for table role_rel_to_action.
type RoleRelToAction struct {
	RoleId    uint        `json:"roleId"    orm:"roleId"    ` // 角色ID
	ActionId  uint        `json:"actionId"  orm:"actionId"  ` // 操作ID
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" orm:"createdAt" ` // 创建时间
}
