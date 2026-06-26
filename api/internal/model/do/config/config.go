// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure of table config for DAO operations like Where/Data.
type Config struct {
	g.Meta      `orm:"table:config, do:true"`
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	SceneId     any         // 场景ID
	RelId       any         // 关联ID。根据scene_id对应不同表
	ConfigKey   any         // 配置键
	ConfigValue any         // 配置值
}
