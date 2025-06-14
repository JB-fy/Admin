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
	IsStop    interface{} // 停用：0否 1是
	AdminId   interface{} // 管理员ID
	OrgId     interface{} // 机构ID
	IsSuper   interface{} // 超管：0否 1是
	Nickname  interface{} // 昵称
	Avatar    interface{} // 头像
	Phone     interface{} // 手机
	Email     interface{} // 邮箱
	Account   interface{} // 账号
}
