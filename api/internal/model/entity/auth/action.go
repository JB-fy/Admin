// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Action is the golang structure for table action.
type Action struct {
	ActionId   uint        `json:"actionId"   orm:"actionId"   ` // 操作ID
	ActionName string      `json:"actionName" orm:"actionName" ` // 名称
	ActionCode string      `json:"actionCode" orm:"actionCode" ` // 标识
	Remark     string      `json:"remark"     orm:"remark"     ` // 备注
	IsStop     uint        `json:"isStop"     orm:"isStop"     ` // 停用：0否 1是
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updatedAt"  ` // 更新时间
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"createdAt"  ` // 创建时间
}
