// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package config

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure for table config.
type Config struct {
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   ` // 更新时间
	SceneId     string      `json:"sceneId"     orm:"scene_id"     ` // 场景ID
	RelId       uint        `json:"relId"       orm:"rel_id"       ` // 关联ID。根据scene_id对应不同表
	ConfigKey   string      `json:"configKey"   orm:"config_key"   ` // 配置键
	ConfigValue string      `json:"configValue" orm:"config_value" ` // 配置值
}
