package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type PayChannelInfo struct {
	ChannelId   *uint   `json:"channel_id,omitempty" dc:"通道ID"`
	ChannelName *string `json:"channel_name,omitempty" dc:"名称"`
	ChannelIcon *string `json:"channel_icon,omitempty" dc:"图标"`
}

/*--------列表 开始--------*/
type PayChannelListReq struct {
	g.Meta `path:"/list" method:"post" tags:"支付" sm:"列表"`
	CommonAllTokenHeaderReq
	SceneId uint `json:"scene_id,omitempty" v:"required|min:1" dc:"支付场景ID"`
}

type PayChannelListRes struct {
	List []PayChannelInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------新增 开始--------*/
type PayOrderCreateReq struct {
	g.Meta `path:"/order/create" method:"post" tags:"支付" sm:"新增"`
	CommonAllTokenHeaderReq
	OrderType *uint8 `json:"order_type,omitempty" v:"required|in:0" dc:"订单类型：0默认。值对应的请求参数必传"`
	// Amount    *float64 `json:"amount,omitempty" v:"required-if:OrderType,0|between:0,99999999.99" dc:"实付金额"`
	// ExtData   *string  `json:"ext_data,omitempty" v:"max-length:120" dc:"扩展数据"`
	Param0 *struct {
		Amount *float64 `json:"amount,omitempty" v:"required|between:0,99999999.99" dc:"金额"`
	} `json:"param_0,omitempty" v:"required-if:OrderType,0" dc:"请求参数0"`
}

type PayOrderCreateRes struct {
	Info PayOrderInfo `json:"info" dc:"详情"`
}

type PayOrderInfo struct {
	OrderId   *uint       `json:"order_id,omitempty" dc:"订单ID"`
	OrderNo   *string     `json:"order_no,omitempty" dc:"订单号"`
	OrderType *uint       `json:"order_type,omitempty" dc:"订单类型：0默认"`
	Amount    *float64    `json:"amount,omitempty" dc:"实付金额"`
	ExtData   *string     `json:"ext_data,omitempty" dc:"扩展数据"`
	OrderIp   *string     `json:"order_ip,omitempty" dc:"订单IP"`
	CreatedAt *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
}

/*--------新增 结束--------*/

/*--------支付 开始--------*/
type PayPayReq struct {
	g.Meta `path:"/pay" method:"post" tags:"支付" sm:"支付"`
	CommonAllTokenHeaderReq
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
	CommonHeaderReq
	PayId uint `json:"pay_id" v:"required|between:1,4294967295" in:"path" dc:"支付ID"`
}

/*--------回调 结束--------*/
