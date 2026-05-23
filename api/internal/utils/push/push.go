package push

import (
	"api/internal/utils/push/model"
	"api/internal/utils/push/tx"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"golang.org/x/sync/singleflight"
)

var (
	pushMap     sync.Map
	pushSfg     singleflight.Group
	pushFuncMap = map[string]model.PushFunc{
		`push_of_tx`: tx.NewPush,
	}
	pushTypeDef = `push_of_tx`
)

func NewPush(ctx context.Context, pushType string, config map[string]any) (obj model.Push) {
	if _, ok := pushFuncMap[pushType]; !ok {
		pushType = pushTypeDef
	}
	key := pushType + gmd5.MustEncrypt(config)
	objTmp, ok := pushMap.Load(key)
	if !ok {
		objTmp, _, _ = pushSfg.Do(key, func() (obj any, err error) {
			obj = pushFuncMap[pushType](ctx, config)
			pushMap.Store(key, obj)
			return
		})
	}
	obj = objTmp.(model.Push)
	return
}
