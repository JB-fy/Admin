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
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	ConfigKey   interface{} // 配置Key
	ConfigValue interface{} // 配置值
}
