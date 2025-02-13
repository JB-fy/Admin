// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package pay

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Channel is the golang structure for table channel.
type Channel struct {
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   ` // 更新时间
	IsStop      uint        `json:"isStop"      orm:"is_stop"      ` // 停用：0否 1是
	ChannelId   uint        `json:"channelId"   orm:"channel_id"   ` // 通道ID
	ChannelName string      `json:"channelName" orm:"channel_name" ` // 名称
	ChannelIcon string      `json:"channelIcon" orm:"channel_icon" ` // 图标
	SceneId     uint        `json:"sceneId"     orm:"scene_id"     ` // 场景ID
	PayId       uint        `json:"payId"       orm:"pay_id"       ` // 支付ID
	PayMethod   uint        `json:"payMethod"   orm:"pay_method"   ` // 支付方法：0APP支付 1H5支付 2扫码支付 3小程序支付
	Sort        uint        `json:"sort"        orm:"sort"         ` // 排序值。从大到小排序
	TotalAmount float64     `json:"totalAmount" orm:"total_amount" ` // 总额
}
