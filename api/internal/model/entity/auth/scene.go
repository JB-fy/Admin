// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Scene is the golang structure for table scene.
type Scene struct {
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   ` // 更新时间
	IsStop      uint        `json:"isStop"      orm:"is_stop"      ` // 停用：0否 1是
	SceneId     string      `json:"sceneId"     orm:"scene_id"     ` // 场景ID
	SceneName   string      `json:"sceneName"   orm:"scene_name"   ` // 名称
	SceneConfig string      `json:"sceneConfig" orm:"scene_config" ` // 配置。JSON格式，根据场景设置
	Remark      string      `json:"remark"      orm:"remark"       ` // 备注
}
