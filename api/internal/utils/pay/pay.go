package pay

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type PayData struct {
	OrderNo   string  //单号
	Amount    float64 //金额。单位：元
	NotifyUrl string  //异步回调地址
	Desc      string  //描述
}

type PayInfo struct {
	PayStr string
}

type Pay interface {
	App(payData PayData) (orderInfo PayInfo, err error) //App支付
}

func NewPay(ctx context.Context, payTypeOpt ...string) Pay {
	payType := ``
	if len(payTypeOpt) > 0 {
		payType = payTypeOpt[0]
	}

	switch payType {
	case `payOfWx`: //微信
		config, _ := daoPlatform.Config.Get(ctx, []string{`payOfWxAppId`, `payOfWxMchid`, `payOfWxSerialNo`, `payOfWxApiV3Key`, `payOfWxCertPath`})
		return NewPayOfWx(ctx, config)
	// case `payOfAli`: //支付宝
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`payOfAliAppId`, `payOfAliSignType`, `payOfAliPrivateKey`, `payOfAliPublicKey`})
		return NewPayOfAli(ctx, config)
	}
}
