// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleRelOfPlatformAdmin is the golang structure of table auth_role_rel_of_platform_admin for DAO operations like Where/Data.
type RoleRelOfPlatformAdmin struct {
	g.Meta    `orm:"table:auth_role_rel_of_platform_admin, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	AdminId   interface{} // 管理员ID
	RoleId    interface{} // 角色ID
}
