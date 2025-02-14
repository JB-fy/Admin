package push

import (
	"api/internal/utils/push/model"
	"api/internal/utils/push/tx"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

var (
	pushMap     = map[string]model.Push{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	pushMu      sync.Mutex
	pushTypeDef = `pushOfTx`
	pushFuncMap = map[string]model.PushFunc{
		`pushOfTx`: tx.NewPush,
	}
)

func NewPush(ctx context.Context, pushType string, config map[string]any) (push model.Push) {
	pushKey := pushType + gmd5.MustEncrypt(config)
	ok := false
	if push, ok = pushMap[pushKey]; ok { //先读一次（不加锁）
		return
	}
	pushMu.Lock()
	defer pushMu.Unlock()
	if push, ok = pushMap[pushKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	if _, ok = pushFuncMap[pushType]; !ok {
		pushType = pushTypeDef
	}
	push = pushFuncMap[pushType](ctx, config)
	pushMap[pushKey] = push
	return
}
