// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure for table config.
type Config struct {
	ConfigId    uint        `json:"configId"    ` // 配置ID
	ConfigKey   string      `json:"configKey"   ` // 配置Key
	ConfigValue string      `json:"configValue" ` // 配置值
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
}
