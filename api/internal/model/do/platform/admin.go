// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure of table platform_admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta    `orm:"table:platform_admin, do:true"`
	AdminId   interface{} // 管理员ID
	Phone     interface{} // 电话号码
	Account   interface{} // 账号
	Password  interface{} // 密码（md5保存）
	Salt      interface{} // 加密盐
	Nickname  interface{} // 昵称
	Avatar    interface{} // 头像
	IsStop    interface{} // 停用：0否 1是
	UpdatedAt *gtime.Time // 更新时间
	CreatedAt *gtime.Time // 创建时间
}
