// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package org

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	IsStop    uint        `json:"isStop"    orm:"is_stop"    ` // 停用：0否 1是
	AdminId   uint        `json:"adminId"   orm:"admin_id"   ` // 管理员ID
	OrgId     uint        `json:"orgId"     orm:"org_id"     ` // 机构ID
	IsSuper   uint        `json:"isSuper"   orm:"is_super"   ` // 超管：0否 1是
	Nickname  string      `json:"nickname"  orm:"nickname"   ` // 昵称
	Avatar    string      `json:"avatar"    orm:"avatar"     ` // 头像
	Phone     string      `json:"phone"     orm:"phone"      ` // 手机
	Email     string      `json:"email"     orm:"email"      ` // 邮箱
	Account   string      `json:"account"   orm:"account"    ` // 账号
}
