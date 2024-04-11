// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleRelOfPlatformAdmin is the golang structure for table role_rel_of_platform_admin.
type RoleRelOfPlatformAdmin struct {
	RoleId    uint        `json:"roleId"    orm:"roleId"    ` // 角色ID
	AdminId   uint        `json:"adminId"   orm:"adminId"   ` // 管理员ID
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" orm:"createdAt" ` // 创建时间
}
