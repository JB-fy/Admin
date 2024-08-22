// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure of table pay_order for DAO operations like Where/Data.
type Order struct {
	g.Meta         `orm:"table:pay_order, do:true"`
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	OrderId        interface{} // 订单ID
	OrderNo        interface{} // 订单号
	RelOrderType   interface{} // 关联订单类型：0默认
	RelOrderUserId interface{} // 关联订单用户ID
	PayId          interface{} // 支付ID
	ChannelId      interface{} // 通道ID
	PayType        interface{} // 类型：0支付宝 1微信
	Amount         interface{} // 实付金额
	PayStatus      interface{} // 状态：0未付款 1已付款
	PayTime        *gtime.Time // 支付时间
	PayRate        interface{} // 费率
	ThirdOrderNo   interface{} // 第三方订单号
}
