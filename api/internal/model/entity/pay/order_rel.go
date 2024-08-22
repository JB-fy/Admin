// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderRel is the golang structure for table order_rel.
type OrderRel struct {
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"        ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"        ` // 更新时间
	OrderId        uint        `json:"orderId"        orm:"order_id"          ` // 订单ID
	RelOrderType   uint        `json:"relOrderType"   orm:"rel_order_type"    ` // 关联订单类型
	RelOrderId     uint        `json:"relOrderId"     orm:"rel_order_id"      ` // 关联订单ID
	RelOrderNo     string      `json:"relOrderNo"     orm:"rel_order_no"      ` // 关联订单号
	RelOrderUserId uint        `json:"relOrderUserId" orm:"rel_order_user_id" ` // 关联订单用户ID
	RelOrderAmount float64     `json:"relOrderAmount" orm:"rel_order_amount"  ` // 关联订单实付金额
}
