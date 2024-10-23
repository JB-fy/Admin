// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleRelOfPlatformAdmin is the golang structure for table role_rel_of_platform_admin.
type RoleRelOfPlatformAdmin struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	AdminId   uint        `json:"adminId"   orm:"admin_id"   ` // 管理员ID
	RoleId    uint        `json:"roleId"    orm:"role_id"    ` // 角色ID
}
