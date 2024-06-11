package pay

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

/* type Device string

const (
	DeviceUnknown Device = `unknown` //未知
	DeviceAndroid Device = `android` //安卓
	DeviceIOS     Device = `ios`     //苹果
) */

type PayData struct {
	OrderNo   string  //单号
	Amount    float64 //金额。单位：元
	Desc      string  //描述
	ReturnUrl string  //同步回调地址。需要时传
	ClientIp  string  //客户端IP。需要时传
	Openid    string  //用户openid。JSAPI支付必传
	// Device    Device  //设备类型。需要时传。 unknown未知 android安卓 ios苹果
}

type PayInfo struct {
	PayStr string //支付字符串
}

type NotifyInfo struct {
	OrderNo        string  //单号
	Amount         float64 //金额。单位：元
	OrderNoOfThird string  //第三方单号
}

type Pay interface {
	App(payData PayData) (orderInfo PayInfo, err error)         // App支付
	H5(payData PayData) (orderInfo PayInfo, err error)          // H5支付
	Jsapi(payData PayData) (orderInfo PayInfo, err error)       // JSAPI支付
	Notify(r *ghttp.Request) (notifyInfo NotifyInfo, err error) // 回调
	NotifyRes(r *ghttp.Request, failMsg string)                 // 回调响应处理
}

func NewPay(ctx context.Context, payTypeOpt ...string) Pay {
	payType := ``
	if len(payTypeOpt) > 0 {
		payType = payTypeOpt[0]
	}

	switch payType {
	case `payOfWx`: //微信
		return NewPayOfWx(ctx)
	// case `payOfAli`: //支付宝
	default:
		return NewPayOfAli(ctx)
	}
}
