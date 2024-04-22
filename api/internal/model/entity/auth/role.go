// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	RoleId    uint        `json:"roleId"    orm:"role_id"    ` // 角色ID
	RoleName  string      `json:"roleName"  orm:"role_name"  ` // 名称
	SceneId   uint        `json:"sceneId"   orm:"scene_id"   ` // 场景ID
	TableId   uint        `json:"tableId"   orm:"table_id"   ` // 关联表ID。0表示平台创建，其它值根据sceneId对应不同表，表示由哪个机构或个人创建
	IsStop    uint        `json:"isStop"    orm:"is_stop"    ` // 停用：0否 1是
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
}
