// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Action is the golang structure of table auth_action for DAO operations like Where/Data.
type Action struct {
	g.Meta     `orm:"table:auth_action, do:true"`
	ActionId   interface{} // 权限操作ID
	ActionName interface{} // 名称
	ActionCode interface{} // 标识（代码中用于判断权限）
	Remark     interface{} // 备注
	IsStop     interface{} // 停用：0否 1是
	UpdatedAt  *gtime.Time // 更新时间
	CreatedAt  *gtime.Time // 创建时间
}
