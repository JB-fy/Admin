// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table auth_role for DAO operations like Where/Data.
type Role struct {
	g.Meta    `orm:"table:auth_role, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsStop    any         // 停用：0否 1是
	RoleId    any         // 角色ID
	RoleName  any         // 名称
	SceneId   any         // 场景ID
	RelId     any         // 关联ID。0表示平台创建，其它值根据scene_id对应不同表
}
