package pay

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type PayOfWx struct {
	Ctx        context.Context
	AppId      string `json:"payOfWxAppId"`
	Mchid      string `json:"payOfWxMchid"`
	SerialNo   string `json:"payOfWxSerialNo"`
	APIv3Key   string `json:"payOfWxApiV3Key"`
	PrivateKey string `json:"payOfWxPrivateKey"`
}

func NewPayOfWx(ctx context.Context, config map[string]interface{}) *PayOfWx {
	payOfWxObj := PayOfWx{
		Ctx: ctx,
	}
	gconv.Struct(config, &payOfWxObj)
	return &payOfWxObj
}

func (payThis *PayOfWx) Create(orderData map[string]interface{}) (orderInfo map[string]interface{}, err error) {
	return
}
