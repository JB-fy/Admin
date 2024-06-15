// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Org is the golang structure for table org.
type Org struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	IsStop    uint        `json:"isStop"    orm:"is_stop"    ` // 停用：0否 1是
	OrgId     uint        `json:"orgId"     orm:"org_id"     ` // 机构ID
	OrgName   string      `json:"orgName"   orm:"org_name"   ` // 机构名称
}
