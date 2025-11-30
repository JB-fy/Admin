package sign

import (
	"api/internal/utils/sign/common"
	"api/internal/utils/sign/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	signMap     = map[string]model.Sign{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	signMuMap   sync.Map
	signFuncMap = map[uint8]model.SignFunc{
		0: common.NewSign,
	}
	signTypeDef uint8 = 0
)

func NewSign(ctx context.Context, signType uint8, config map[string]any) (sign model.Sign) {
	signKey := gconv.String(signType) + gmd5.MustEncrypt(config)
	ok := false
	if sign, ok = signMap[signKey]; ok { //先读一次（不加锁）
		return
	}
	muTmp, _ := signMuMap.LoadOrStore(signKey, &sync.Mutex{})
	mu := muTmp.(*sync.Mutex)
	mu.Lock()
	defer func() {
		mu.Unlock()
		signMuMap.Delete(signKey)
	}()
	if sign, ok = signMap[signKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	if _, ok = signFuncMap[signType]; !ok {
		signType = signTypeDef
	}
	sign = signFuncMap[signType](ctx, config)
	signMap[signKey] = sign
	return
}
