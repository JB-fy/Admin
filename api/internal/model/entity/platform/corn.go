// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Corn is the golang structure for table corn.
type Corn struct {
	CornId      uint        `json:"cornId"      ` // 定时器ID
	CornCode    string      `json:"cornCode"    ` // 标识
	CornName    string      `json:"cornName"    ` // 名称
	CornPattern string      `json:"cornPattern" ` // 表达式
	Remark      string      `json:"remark"      ` // 备注
	IsStop      uint        `json:"isStop"      ` // 是否停用：0否 1是
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
}
