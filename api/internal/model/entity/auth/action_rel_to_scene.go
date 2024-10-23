// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActionRelToScene is the golang structure for table action_rel_to_scene.
type ActionRelToScene struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	ActionId  uint        `json:"actionId"  orm:"action_id"  ` // 操作ID
	SceneId   uint        `json:"sceneId"   orm:"scene_id"   ` // 场景ID
}
