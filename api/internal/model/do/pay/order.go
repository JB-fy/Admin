// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure of table pay_order for DAO operations like Where/Data.
type Order struct {
	g.Meta         `orm:"table:pay_order, do:true"`
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	OrderId        any         // 订单ID
	OrderNo        any         // 订单号
	RelOrderType   any         // 关联订单类型：0默认
	RelOrderUserId any         // 关联订单用户ID
	PayId          any         // 支付ID
	ChannelId      any         // 通道ID
	PayType        any         // 类型：0支付宝 1微信
	Amount         any         // 实付金额
	PayStatus      any         // 状态：0未付款 1已付款
	PayTime        *gtime.Time // 支付时间
	PayRate        any         // 费率
	ThirdOrderNo   any         // 第三方订单号
}
