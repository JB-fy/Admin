package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------支付 开始--------*/
type PayPayReq struct {
	g.Meta  `path:"/pay" method:"post" tags:"支付" sm:"支付"`
	OrderNo string `json:"orderNo" v:"required|max-length:60" dc:"订单号"`
	PayType *uint  `json:"payType" v:"required|in:0,1,2,10,11,12" dc:"支付类型：0APP支付(支付宝) 1H5支付(支付宝) 2小程序支付(支付宝) 10APP支付(微信) 11H5支付(微信) 12小程序支付(微信)"`
}

type PayPayRes struct {
	PayStr string `json:"payStr" dc:"支付字符串"`
}

/*--------支付 结束--------*/

/*--------回调 开始--------*/
type PayNotifyReq struct {
	g.Meta  `path:"/notify/:payType" method:"get,post" tags:"支付" sm:"回调"`
	PayType string `json:"payType" v:"required|in:payOfAli,payOfWx" in:"path" dc:"支付方式"`
}

/*--------回调 结束--------*/
