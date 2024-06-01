// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleRelToAction is the golang structure for table role_rel_to_action.
type RoleRelToAction struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	RoleId    uint        `json:"roleId"    orm:"role_id"    ` // 角色ID
	ActionId  uint        `json:"actionId"  orm:"action_id"  ` // 操作ID
}
