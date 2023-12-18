package pay

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type AliPay struct {
	Ctx        context.Context
	AppId      string `json:"aliPayAppId"`
	SignType   string `json:"aliPaySignType"`
	PrivateKey string `json:"aliPayPrivateKey"`
	PublicKey  string `json:"aliPayPublicKey"`
}

func NewAliPay(ctx context.Context, config map[string]interface{}) *AliPay {
	aliPayObj := AliPay{
		Ctx: ctx,
	}
	gconv.Struct(config, &aliPayObj)
	return &aliPayObj
}

func (payThis *AliPay) Create(orderData map[string]interface{}) (orderInfo map[string]interface{}, err error) {
	return
}
