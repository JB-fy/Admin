package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type PayChannelInfo struct {
	ChannelId   *uint   `json:"channel_id,omitempty" dc:"通道ID"`
	ChannelName *string `json:"channel_name,omitempty" dc:"名称"`
	ChannelIcon *string `json:"channel_icon,omitempty" dc:"图标"`
}

/*--------列表 开始--------*/
type PayChannelListReq struct {
	g.Meta  `path:"/list" method:"post" tags:"支付" sm:"列表"`
	SceneId uint `json:"scene_id,omitempty" v:"required|min:1" dc:"支付场景ID"`
}

type PayChannelListRes struct {
	List []PayChannelInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------支付 开始--------*/
type PayPayReq struct {
	g.Meta    `path:"/pay" method:"post" tags:"支付" sm:"支付"`
	ChannelId uint   `json:"channel_id" v:"required|between:1,4294967295" dc:"通道ID"`
	OrderId   uint   `json:"order_id" v:"required-without:OrderNo|between:1,4294967295" dc:"订单ID。订单ID和订单号二选一"`
	OrderNo   string `json:"order_no" v:"required-without:OrderId|max-length:60" dc:"订单号。订单ID和订单号二选一"`
}

type PayPayRes struct {
	PayStr string `json:"payStr" dc:"支付字符串"`
}

/*--------支付 结束--------*/

/*--------回调 开始--------*/
type PayNotifyReq struct {
	g.Meta `path:"/notify/:pay_id" method:"get,post" tags:"支付" sm:"回调"`
	PayId  uint `json:"pay_id" v:"required|between:1,4294967295" in:"path" dc:"支付ID"`
}

/*--------回调 结束--------*/
