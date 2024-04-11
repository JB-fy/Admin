// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure for table config.
type Config struct {
	ConfigKey   string      `json:"configKey"   orm:"configKey"   ` // 配置Key
	ConfigValue string      `json:"configValue" orm:"configValue" ` // 配置值
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updatedAt"   ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"createdAt"   ` // 创建时间
}
