// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Cron is the golang structure of table platform_cron for DAO operations like Where/Data.
type Cron struct {
	g.Meta      `orm:"table:platform_cron, do:true"`
	CronId      interface{} // 定时器ID
	CronName    interface{} // 名称
	CronCode    interface{} // 标识
	CronPattern interface{} // 表达式
	Remark      interface{} // 备注
	IsStop      interface{} // 是否停用：0否 1是
	UpdatedAt   *gtime.Time // 更新时间
	CreatedAt   *gtime.Time // 创建时间
}
