package pay

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type Pay interface {
	Create(orderData map[string]interface{}) (orderInfo map[string]interface{}, err error)
}

// 支付方式：0-支付宝 1-微信
func NewPay(ctx context.Context, payType uint) Pay {
	switch payType {
	case 1: //微信
		config, _ := daoPlatform.Config.Get(ctx, []string{`wxPayAppId`, `wxPayMchid`, `wxPaySerialNo`, `wxPayApiV3Key`, `wxPayPrivateKey`})
		return NewWxPay(ctx, config)
	// case 0: //支付宝
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`aliPayAppId`, `aliPaySignType`, `aliPayPrivateKey`, `aliPayPublicKey`})
		return NewAliPay(ctx, config)
	}
}
