// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Action is the golang structure for table action.
type Action struct {
	ActionId   uint        `json:"actionId"   orm:"action_id"   ` // 操作ID
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  ` // 更新时间
	IsStop     uint        `json:"isStop"     orm:"is_stop"     ` // 停用：0否 1是
	ActionName string      `json:"actionName" orm:"action_name" ` // 名称
	ActionCode string      `json:"actionCode" orm:"action_code" ` // 标识
	Remark     string      `json:"remark"     orm:"remark"      ` // 备注
}
