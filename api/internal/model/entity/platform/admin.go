// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	AdminId   uint        `json:"adminId"   orm:"adminId"   ` // 管理员ID
	Phone     string      `json:"phone"     orm:"phone"     ` // 手机
	Account   string      `json:"account"   orm:"account"   ` // 账号
	Password  string      `json:"password"  orm:"password"  ` // 密码。md5保存
	Salt      string      `json:"salt"      orm:"salt"      ` // 加密盐
	Nickname  string      `json:"nickname"  orm:"nickname"  ` // 昵称
	Avatar    string      `json:"avatar"    orm:"avatar"    ` // 头像
	IsStop    uint        `json:"isStop"    orm:"isStop"    ` // 停用：0否 1是
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" orm:"createdAt" ` // 创建时间
}
