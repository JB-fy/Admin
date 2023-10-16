// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table auth_role for DAO operations like Where/Data.
type Role struct {
	g.Meta    `orm:"table:auth_role, do:true"`
	RoleId    interface{} // 角色ID
	RoleName  interface{} // 名称
	SceneId   interface{} // 场景ID
	TableId   interface{} // 关联表ID。0表示平台创建，其它值根据sceneId对应不同表，表示由哪个机构或个人创建
	IsStop    interface{} // 停用：0否 1是
	UpdatedAt *gtime.Time // 更新时间
	CreatedAt *gtime.Time // 创建时间
}
