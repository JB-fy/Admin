// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package org

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure of table org_admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta    `orm:"table:org_admin, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsStop    any         // 停用：0否 1是
	AdminId   any         // 管理员ID
	OrgId     any         // 机构ID
	IsSuper   any         // 超管：0否 1是
	Nickname  any         // 昵称
	Avatar    any         // 头像
	Phone     any         // 手机
	Email     any         // 邮箱
	Account   any         // 账号
}
