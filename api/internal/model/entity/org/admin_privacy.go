// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package org

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminPrivacy is the golang structure for table admin_privacy.
type AdminPrivacy struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	AdminId   uint        `json:"adminId"   orm:"admin_id"   ` // 管理员ID
	Password  string      `json:"password"  orm:"password"   ` // 密码。md5保存
	Salt      string      `json:"salt"      orm:"salt"       ` // 密码盐
}
