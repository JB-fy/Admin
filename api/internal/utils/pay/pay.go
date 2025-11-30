package pay

import (
	"api/internal/utils/pay/ali"
	"api/internal/utils/pay/model"
	"api/internal/utils/pay/wx"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	payMap     = map[string]model.Pay{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	payMuMap   sync.Map
	payFuncMap = map[uint]model.PayFunc{
		0: ali.NewPay,
		1: wx.NewPay,
	}
	payTypeDef uint = 0
)

func NewPay(ctx context.Context, payType uint, config map[string]any) (pay model.Pay) {
	payKey := gconv.String(payType) + gmd5.MustEncrypt(config)
	ok := false
	if pay, ok = payMap[payKey]; ok { //先读一次（不加锁）
		return
	}
	muTmp, _ := payMuMap.LoadOrStore(payKey, &sync.Mutex{})
	mu := muTmp.(*sync.Mutex)
	mu.Lock()
	defer func() {
		mu.Unlock()
		payMuMap.Delete(payKey)
	}()
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
