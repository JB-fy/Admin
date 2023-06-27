// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Corn is the golang structure of table platform_corn for DAO operations like Where/Data.
type Corn struct {
	g.Meta      `orm:"table:platform_corn, do:true"`
	CornId      interface{} // 定时器ID
	CornCode    interface{} // 标识
	CornName    interface{} // 名称
	CornPattern interface{} // 表达式
	Remark      interface{} // 备注
	IsStop      interface{} // 是否停用：0否 1是
	UpdatedAt   *gtime.Time // 更新时间
	CreatedAt   *gtime.Time // 创建时间
}
