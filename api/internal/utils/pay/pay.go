package pay

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

/* type OrderData struct {
	OrderNo   string
	Amount    float64
	NotifyUrl string
	Desc      string
} */

type Pay interface {
	Create(orderData map[string]interface{}) (orderInfo map[string]interface{}, err error)
}

func NewPay(ctx context.Context, payTypeTmp ...string) Pay {
	payType := ``
	if len(payTypeTmp) > 0 {
		payType = payTypeTmp[0]
	}

	switch payType {
	case `wxPay`: //微信
		config, _ := daoPlatform.Config.Get(ctx, []string{`wxPayAppId`, `wxPayMchid`, `wxPaySerialNo`, `wxPayApiV3Key`, `wxPayPrivateKey`})
		return NewWxPay(ctx, config)
	// case `aliPay`: //支付宝
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`aliPayAppId`, `aliPaySignType`, `aliPayPrivateKey`, `aliPayPublicKey`})
		return NewAliPay(ctx, config)
	}
}
