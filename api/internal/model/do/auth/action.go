// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Action is the golang structure of table auth_action for DAO operations like Where/Data.
type Action struct {
	g.Meta     `orm:"table:auth_action, do:true"`
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	IsStop     interface{} // 停用：0否 1是
	ActionId   interface{} // 操作ID
	ActionName interface{} // 名称
	ActionCode interface{} // 标识
	Remark     interface{} // 备注
}
