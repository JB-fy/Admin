// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ActionRelToScene is the golang structure of table auth_action_rel_to_scene for DAO operations like Where/Data.
type ActionRelToScene struct {
	g.Meta    `orm:"table:auth_action_rel_to_scene, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	ActionId  interface{} // 操作ID
	SceneId   interface{} // 场景ID
}
