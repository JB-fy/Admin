package sign

import (
	"api/internal/utils/sign/common"
	"api/internal/utils/sign/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/sync/singleflight"
)

var (
	signMap     sync.Map
	signSfg     singleflight.Group
	signFuncMap = map[uint8]model.SignFunc{
		0: common.NewSign,
	}
	signTypeDef uint8 = 0
)

func NewSign(ctx context.Context, signType uint8, config map[string]any) (obj model.Sign) {
	if _, ok := signFuncMap[signType]; !ok {
		signType = signTypeDef
	}
	key := gconv.String(signType) + gmd5.MustEncrypt(config)
	objTmp, ok := signMap.Load(key)
	if !ok {
		objTmp, _, _ = signSfg.Do(key, func() (obj any, err error) {
			obj = signFuncMap[signType](ctx, config)
			signMap.Store(key, obj)
			return
		})
	}
	obj = objTmp.(model.Sign)
	return
}
