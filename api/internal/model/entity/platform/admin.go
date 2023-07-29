// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	AdminId   uint        `json:"adminId"   ` // 管理员ID
	Phone     string      `json:"phone"     ` // 电话号码
	Account   string      `json:"account"   ` // 账号
	Password  string      `json:"password"  ` // 密码（md5保存）
	Salt      string      `json:"salt"      ` // 密码盐
	Nickname  string      `json:"nickname"  ` // 昵称
	Avatar    string      `json:"avatar"    ` // 头像
	IsStop    uint        `json:"isStop"    ` // 是否停用：0否 1是
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
}
