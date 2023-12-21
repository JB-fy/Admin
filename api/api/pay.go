package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------回调 开始--------*/
type PayNotifyReq struct {
	g.Meta  `path:"/notify/:payType" method:"get,post" tags:"支付" sm:"回调"`
	PayType string `json:"payType" v:"required|in:payOfAli,payOfWx" in:"path" dc:"支付方式"`
}

/*--------回调 结束--------*/
