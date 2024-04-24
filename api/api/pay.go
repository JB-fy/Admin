package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------列表 开始--------*/
type PayListReq struct {
	g.Meta `path:"/list" method:"post" tags:"支付" sm:"列表"`
	// UseScene uint `json:"use_scene" v:"required|in:0" dc:"使用场景：0手机APP 1手机浏览器 2电脑浏览器"`
}

type PayListRes struct {
	List []PayListItem `json:"list" dc:"列表"`
}

type PayListItem struct {
	PayMethod uint   `json:"pay_method" dc:"支付方式：1APP支付(支付宝) 2H5支付(支付宝) 3JSAPI支付(支付宝) 11APP支付(微信) 12H5支付(微信) 13JSAPI支付(微信)"`
	PayName   string `json:"pay_name" dc:"名称"`
	PayIcon   string `json:"pay_icon" dc:"图标"`
}

/*--------列表 结束--------*/

/*--------支付 开始--------*/
type PayPayReq struct {
	g.Meta    `path:"/pay" method:"post" tags:"支付" sm:"支付"`
	OrderNo   string `json:"order_no" v:"required|max-length:60" dc:"订单号"`
	PayMethod uint   `json:"pay_method" v:"required|in:1,2,3,11,12,13" dc:"支付类型：1APP支付(支付宝) 2H5支付(支付宝) 3JSAPI支付(支付宝) 11APP支付(微信) 12H5支付(微信) 13JSAPI支付(微信)"`
}

type PayPayRes struct {
	PayStr string `json:"payStr" dc:"支付字符串"`
}

/*--------支付 结束--------*/

/*--------回调 开始--------*/
type PayNotifyReq struct {
	g.Meta  `path:"/notify/:pay_type" method:"get,post" tags:"支付" sm:"回调"`
	PayType string `json:"pay_type" v:"required|in:payOfAli,payOfWx" in:"path" dc:"支付方式"`
}

/*--------回调 结束--------*/
