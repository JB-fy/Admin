// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package platform

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminPrivacy is the golang structure of table platform_admin_privacy for DAO operations like Where/Data.
type AdminPrivacy struct {
	g.Meta    `orm:"table:platform_admin_privacy, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	AdminId   any         // 管理员ID
	Password  any         // 密码。md5保存
	Salt      any         // 密码盐
}
