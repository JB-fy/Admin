package model

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type PayFunc func(ctx context.Context, config map[string]any) Pay

type Pay interface {
	App(ctx context.Context, payReq PayReq) (payRes PayRes, err error)               // App支付
	H5(ctx context.Context, payReq PayReq) (payRes PayRes, err error)                // H5支付
	QRCode(ctx context.Context, payReq PayReq) (payRes PayRes, err error)            // 扫码支付
	Jsapi(ctx context.Context, payReq PayReq) (payRes PayRes, err error)             // 小程序支付
	Notify(ctx context.Context, r *ghttp.Request) (notifyInfo NotifyInfo, err error) // 回调验证
	NotifyRes(ctx context.Context, r *ghttp.Request, failMsg string)                 // 回调响应
}

/* type Device string

const (
	DeviceUnknown Device = `unknown` //未知
	DeviceAndroid Device = `android` //安卓
	DeviceIOS     Device = `ios`     //苹果
) */

type PayReq struct {
	OrderNo   string  //单号
	Amount    float64 //金额。单位：元
	Desc      string  //描述
	ReturnUrl string  //同步回调地址。需要时传
	ClientIp  string  //客户端IP。需要时传
	Openid    string  //用户openid。JSAPI支付必传
	// Device    Device  //设备类型。需要时传。 unknown未知 android安卓 ios苹果
}

type PayRes struct {
	PayStr string //支付字符串
}

type NotifyInfo struct {
	OrderNo      string  //单号
	Amount       float64 //金额。单位：元
	ThirdOrderNo string  //第三方单号
}
