package pay

import (
	daoPay "api/internal/dao/pay"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
)

/* type Device string

const (
	DeviceUnknown Device = `unknown` //未知
	DeviceAndroid Device = `android` //安卓
	DeviceIOS     Device = `ios`     //苹果
) */

type PayReqData struct {
	OrderNo   string  //单号
	Amount    float64 //金额。单位：元
	Desc      string  //描述
	ReturnUrl string  //同步回调地址。需要时传
	ClientIp  string  //客户端IP。需要时传
	Openid    string  //用户openid。JSAPI支付必传
	// Device    Device  //设备类型。需要时传。 unknown未知 android安卓 ios苹果
}

type PayResData struct {
	PayStr string //支付字符串
}

type NotifyInfo struct {
	OrderNo        string  //单号
	Amount         float64 //金额。单位：元
	OrderNoOfThird string  //第三方单号
}

type Pay interface {
	App(payReqData PayReqData) (payResData PayResData, err error)    // App支付
	H5(payReqData PayReqData) (payResData PayResData, err error)     // H5支付
	QRCode(payReqData PayReqData) (payResData PayResData, err error) // 扫码支付
	Jsapi(payReqData PayReqData) (payResData PayResData, err error)  // 小程序支付
	Notify(r *ghttp.Request) (notifyInfo NotifyInfo, err error)      // 回调验证
	NotifyRes(r *ghttp.Request, failMsg string)                      // 回调响应
}

func NewPay(ctx context.Context, payInfo gdb.Record) Pay {
	config := payInfo[daoPay.Pay.Columns().PayConfig].Map()
	config[`notifyUrl`] = utils.GetRequestUrl(ctx, 0) + `/pay/notify/` + payInfo[daoPay.Pay.Columns().PayId].String()

	switch payInfo[daoPay.Pay.Columns().PayType].Uint() {
	case 1: //微信
		return NewPayOfWx(ctx, config)
	// case 0: //支付宝
	default:
		return NewPayOfAli(ctx, config)
	}
}
