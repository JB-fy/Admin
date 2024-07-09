// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	IsStop    uint        `json:"isStop"    orm:"is_stop"    ` // 停用：0否 1是
	RoleId    uint        `json:"roleId"    orm:"role_id"    ` // 角色ID
	RoleName  string      `json:"roleName"  orm:"role_name"  ` // 名称
	SceneId   uint        `json:"sceneId"   orm:"scene_id"   ` // 场景ID
	RelId     uint        `json:"relId"     orm:"rel_id"     ` // 关联ID。0表示平台创建，其它值根据scene_id对应不同表
}
