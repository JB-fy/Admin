package pay

import (
	"api/internal/utils/pay/ali"
	"api/internal/utils/pay/model"
	"api/internal/utils/pay/wx"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/sync/singleflight"
)

var (
	payMap     sync.Map
	paySfg     singleflight.Group
	payFuncMap = map[uint]model.PayFunc{
		0: ali.NewPay,
		1: wx.NewPay,
	}
	payTypeDef uint = 0
)

func NewPay(ctx context.Context, payType uint, config map[string]any) (obj model.Pay) {
	if _, ok := payFuncMap[payType]; !ok {
		payType = payTypeDef
	}
	key := gconv.String(payType) + gmd5.MustEncrypt(config)
	objTmp, ok := payMap.Load(key)
	if !ok {
		objTmp, _, _ = paySfg.Do(key, func() (obj any, err error) {
			obj = payFuncMap[payType](ctx, config)
			payMap.Store(key, obj)
			return
		})
	}
	obj = objTmp.(model.Pay)
	return
}
