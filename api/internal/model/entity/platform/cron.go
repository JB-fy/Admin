// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Cron is the golang structure for table cron.
type Cron struct {
	CronId      uint        `json:"cronId"      ` // 定时器ID
	CronName    string      `json:"cronName"    ` // 名称
	CronCode    string      `json:"cronCode"    ` // 标识
	CronPattern string      `json:"cronPattern" ` // 表达式
	Remark      string      `json:"remark"      ` // 备注
	IsStop      uint        `json:"isStop"      ` // 是否停用：0否 1是
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
}
