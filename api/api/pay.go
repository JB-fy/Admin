package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type PayInfo struct {
	PayId   *uint   `json:"pay_id,omitempty" dc:"支付ID"`
	PayName *string `json:"pay_name,omitempty" dc:"名称"`
	PayIcon *string `json:"pay_icon,omitempty" dc:"图标"`
}

/*--------列表 开始--------*/
type PayListReq struct {
	g.Meta   `path:"/list" method:"post" tags:"支付" sm:"列表"`
	PayScene uint `json:"pay_scene,omitempty" v:"required|in:0,1,2,10,11,20" dc:"支付场景：0APP 1H5 2扫码 10微信小程序 11微信公众号 20支付宝小程序"`
}

type PayListRes struct {
	List []PayInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------支付 开始--------*/
type PayPayReq struct {
	g.Meta   `path:"/pay" method:"post" tags:"支付" sm:"支付"`
	PayScene uint   `json:"pay_scene,omitempty" v:"required|in:0,1,2,10,11,20" dc:"支付场景：0APP 1H5 2扫码 10微信小程序 11微信公众号 20支付宝小程序"`
	PayId    uint   `json:"pay_id" v:"required|between:1,4294967295" dc:"支付ID"`
	OrderNo  string `json:"order_no" v:"required|max-length:60" dc:"订单号"`
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
