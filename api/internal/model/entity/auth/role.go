// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	RoleId    uint        `json:"roleId"    ` // 角色ID
	RoleName  string      `json:"roleName"  ` // 名称
	SceneId   uint        `json:"sceneId"   ` // 场景ID
	TableId   uint        `json:"tableId"   ` // 关联表ID。0表示平台创建，其它值根据sceneId对应不同表，表示由哪个机构或个人创建
	IsStop    uint        `json:"isStop"    ` // 停用：0否 1是
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
}
