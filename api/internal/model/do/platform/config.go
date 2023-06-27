// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure of table platform_config for DAO operations like Where/Data.
type Config struct {
	g.Meta      `orm:"table:platform_config, do:true"`
	ConfigId    interface{} // 配置ID
	ConfigKey   interface{} // 配置项Key
	ConfigValue interface{} // 配置项值（设置大点。以后可能需要保存富文本内容，如公司简介或协议等等）
	UpdatedAt   *gtime.Time // 更新时间
	CreatedAt   *gtime.Time // 创建时间
}
