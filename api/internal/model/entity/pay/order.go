// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package pay

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Order is the golang structure for table order.
type Order struct {
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"        ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"        ` // 更新时间
	OrderId        uint        `json:"orderId"        orm:"order_id"          ` // 订单ID
	OrderNo        string      `json:"orderNo"        orm:"order_no"          ` // 订单号
	RelOrderType   uint        `json:"relOrderType"   orm:"rel_order_type"    ` // 关联订单类型：0默认
	RelOrderUserId uint        `json:"relOrderUserId" orm:"rel_order_user_id" ` // 关联订单用户ID
	PayId          uint        `json:"payId"          orm:"pay_id"            ` // 支付ID
	ChannelId      uint        `json:"channelId"      orm:"channel_id"        ` // 通道ID
	PayType        uint        `json:"payType"        orm:"pay_type"          ` // 类型：0支付宝 1微信
	Amount         float64     `json:"amount"         orm:"amount"            ` // 实付金额
	PayStatus      uint        `json:"payStatus"      orm:"pay_status"        ` // 状态：0未付款 1已付款
	PayTime        *gtime.Time `json:"payTime"        orm:"pay_time"          ` // 支付时间
	PayRate        float64     `json:"payRate"        orm:"pay_rate"          ` // 费率
	ThirdOrderNo   string      `json:"thirdOrderNo"   orm:"third_order_no"    ` // 第三方订单号
}
