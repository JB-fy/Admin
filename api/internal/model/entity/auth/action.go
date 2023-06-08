// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Action is the golang structure for table action.
type Action struct {
	ActionId   uint        `json:"actionId"   ` // 权限操作ID
	ActionName string      `json:"actionName" ` // 名称
	ActionCode string      `json:"actionCode" ` // 标识（代码中用于判断权限）
	Remark     string      `json:"remark"     ` // 备注
	IsStop     uint        `json:"isStop"     ` // 是否停用：0否 1是
	UpdateAt   *gtime.Time `json:"updateAt"   ` // 更新时间
	CreateAt   *gtime.Time `json:"createAt"   ` // 创建时间
}
