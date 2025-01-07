package pay

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
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
	OrderNo      string  //单号
	Amount       float64 //金额。单位：元
	ThirdOrderNo string  //第三方单号
}

type Pay interface {
	App(ctx context.Context, payReqData PayReqData) (payResData PayResData, err error)    // App支付
	H5(ctx context.Context, payReqData PayReqData) (payResData PayResData, err error)     // H5支付
	QRCode(ctx context.Context, payReqData PayReqData) (payResData PayResData, err error) // 扫码支付
	Jsapi(ctx context.Context, payReqData PayReqData) (payResData PayResData, err error)  // 小程序支付
	Notify(ctx context.Context, r *ghttp.Request) (notifyInfo NotifyInfo, err error)      // 回调验证
	NotifyRes(ctx context.Context, r *ghttp.Request, failMsg string)                      // 回调响应
}

var (
	payTypeDef uint = 0
	payFuncMap      = map[uint]func(ctx context.Context, config map[string]any) Pay{
		0: func(ctx context.Context, config map[string]any) Pay { return NewPayOfAli(ctx, config) },
		1: func(ctx context.Context, config map[string]any) Pay { return NewPayOfWx(ctx, config) },
	}
	payMap = map[string]Pay{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	payMu  sync.Mutex
)

func NewPay(ctx context.Context, payType uint, config map[string]any) (pay Pay) {
	payKey := gconv.String(payType) + gmd5.MustEncrypt(config)
	ok := false
	if pay, ok = payMap[payKey]; ok { //先读一次（不加锁）
		return
	}
	payMu.Lock()
	defer payMu.Unlock()
	if pay, ok = payMap[payKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	if _, ok = payFuncMap[payType]; !ok {
		payType = payTypeDef
	}
	pay = payFuncMap[payType](ctx, config)
	payMap[payKey] = pay
	return
}
