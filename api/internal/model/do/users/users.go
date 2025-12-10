// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package users

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta    `orm:"table:users, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsStop    any         // 停用：0否 1是
	UserId    any         // 用户ID
	Nickname  any         // 昵称
	Avatar    any         // 头像
	Gender    any         // 性别：0未设置 1男 2女
	Birthday  *gtime.Time // 生日
	Address   any         // 地址
	Phone     any         // 手机
	Email     any         // 邮箱
	Account   any         // 账号
	WxOpenid  any         // 微信openid
	WxUnionid any         // 微信unionid
}
