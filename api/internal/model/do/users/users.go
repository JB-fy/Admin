// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta    `orm:"table:users, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsStop    interface{} // 停用：0否 1是
	UserId    interface{} // 用户ID
	Nickname  interface{} // 昵称
	Avatar    interface{} // 头像
	Gender    interface{} // 性别：0未设置 1男 2女
	Birthday  *gtime.Time // 生日
	Address   interface{} // 地址
	Phone     interface{} // 手机
	Email     interface{} // 邮箱
	Account   interface{} // 账号
	WxOpenId  interface{} // 微信openId
	WxUnionId interface{} // 微信unionId
}
