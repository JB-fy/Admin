// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderRel is the golang structure of table pay_order_rel for DAO operations like Where/Data.
type OrderRel struct {
	g.Meta         `orm:"table:pay_order_rel, do:true"`
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	OrderId        any         // 订单ID
	RelOrderType   any         // 关联订单类型：0默认
	RelOrderId     any         // 关联订单ID
	RelOrderNo     any         // 关联订单号
	RelOrderUserId any         // 关联订单用户ID
	RelOrderAmount any         // 关联订单实付金额
}
