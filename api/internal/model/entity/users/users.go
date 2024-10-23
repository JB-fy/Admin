// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package users

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	IsStop    uint        `json:"isStop"    orm:"is_stop"    ` // 停用：0否 1是
	UserId    uint        `json:"userId"    orm:"user_id"    ` // 用户ID
	Nickname  string      `json:"nickname"  orm:"nickname"   ` // 昵称
	Avatar    string      `json:"avatar"    orm:"avatar"     ` // 头像
	Gender    uint        `json:"gender"    orm:"gender"     ` // 性别：0未设置 1男 2女
	Birthday  *gtime.Time `json:"birthday"  orm:"birthday"   ` // 生日
	Address   string      `json:"address"   orm:"address"    ` // 地址
	Phone     string      `json:"phone"     orm:"phone"      ` // 手机
	Email     string      `json:"email"     orm:"email"      ` // 邮箱
	Account   string      `json:"account"   orm:"account"    ` // 账号
	WxOpenid  string      `json:"wxOpenid"  orm:"wx_openid"  ` // 微信openid
	WxUnionid string      `json:"wxUnionid" orm:"wx_unionid" ` // 微信unionid
}
