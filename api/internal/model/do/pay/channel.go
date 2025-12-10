// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Channel is the golang structure of table pay_channel for DAO operations like Where/Data.
type Channel struct {
	g.Meta      `orm:"table:pay_channel, do:true"`
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	IsStop      any         // 停用：0否 1是
	ChannelId   any         // 通道ID
	ChannelName any         // 名称
	ChannelIcon any         // 图标
	SceneId     any         // 场景ID
	PayId       any         // 支付ID
	PayMethod   any         // 支付方法：0APP支付 1H5支付 2扫码支付 3小程序支付
	Sort        any         // 排序值。从大到小排序
	TotalAmount any         // 总额
}
