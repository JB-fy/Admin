// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package org

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Org is the golang structure of table org for DAO operations like Where/Data.
type Org struct {
	g.Meta    `orm:"table:org, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsStop    any         // 停用：0否 1是
	OrgId     any         // 机构ID
	OrgName   any         // 机构名称
}
