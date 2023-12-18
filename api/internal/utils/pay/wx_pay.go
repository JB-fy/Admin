package pay

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type WxPay struct {
	Ctx        context.Context
	AppId      string `json:"wxPayAppId"`
	Mchid      string `json:"wxPayMchid"`
	SerialNo   string `json:"wxPaySerialNo"`
	APIv3Key   string `json:"wxPayApiV3Key"`
	PrivateKey string `json:"wxPayPrivateKey"`
}

func NewWxPay(ctx context.Context, config map[string]interface{}) *WxPay {
	wxPayObj := WxPay{
		Ctx: ctx,
	}
	gconv.Struct(config, &wxPayObj)
	return &wxPayObj
}

func (payThis *WxPay) Create(orderData map[string]interface{}) (orderInfo map[string]interface{}, err error) {
	return
}
