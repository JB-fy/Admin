package pay

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type PayData struct {
	OrderNo string  //单号
	Amount  float64 //金额。单位：元
	Desc    string  //描述
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
	App(payData PayData) (orderInfo PayInfo, err error) //App支付
	Notify() (notifyInfo NotifyInfo, err error)         // 回调
	NotifyRes(failMsg string)                           // 回调响应处理
}

func NewPay(ctx context.Context, payTypeOpt ...string) Pay {
	payType := ``
	if len(payTypeOpt) > 0 {
		payType = payTypeOpt[0]
	}

	switch payType {
	case `payOfWx`: //微信
		config, _ := daoPlatform.Config.Get(ctx, []string{`payOfWxAppId`, `payOfWxMchid`, `payOfWxSerialNo`, `payOfWxApiV3Key`, `payOfWxPrivateKey`, `payOfWxNotifyUrl`})
		return NewPayOfWx(ctx, config)
	// case `payOfAli`: //支付宝
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`payOfAliAppId`, `payOfAliSignType`, `payOfAliPrivateKey`, `payOfAliPublicKey`, `payOfAliNotifyUrl`, `payOfAliReturnUrl`})
		return NewPayOfAli(ctx, config)
	}
}
