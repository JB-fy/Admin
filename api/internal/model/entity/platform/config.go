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
	ConfigKey   string      `json:"configKey"   ` // 配置项Key
	ConfigValue string      `json:"configValue" ` // 配置项值（设置大点。以后可能需要保存富文本内容，如公司简介或协议等等）
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
}
