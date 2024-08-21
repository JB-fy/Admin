// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Channel is the golang structure of table pay_channel for DAO operations like Where/Data.
type Channel struct {
	g.Meta      `orm:"table:pay_channel, do:true"`
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	IsStop      interface{} // 停用：0否 1是
	ChannelId   interface{} // 通道ID
	ChannelName interface{} // 名称
	ChannelIcon interface{} // 图标
	SceneId     interface{} // 场景ID
	PayId       interface{} // 支付ID
	PayMethod   interface{} // 方法：0APP支付 1H5支付 2扫码支付 3小程序支付
	Sort        interface{} // 排序值。从大到小排序
	TotalAmount interface{} // 总额
}
