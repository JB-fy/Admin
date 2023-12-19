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

func NewPay(ctx context.Context, payTypeOpt ...string) Pay {
	payType := ``
	if len(payTypeOpt) > 0 {
		payType = payTypeOpt[0]
	}

	switch payType {
	case `payOfWx`: //微信
		config, _ := daoPlatform.Config.Get(ctx, []string{`payOfWxAppId`, `payOfWxMchid`, `payOfWxSerialNo`, `payOfWxApiV3Key`, `payOfWxPrivateKey`})
		return NewPayOfWx(ctx, config)
	// case `payOfAli`: //支付宝
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`payOfAliAppId`, `payOfAliSignType`, `payOfAliPrivateKey`, `payOfAliPublicKey`})
		return NewPayOfAli(ctx, config)
	}
}
