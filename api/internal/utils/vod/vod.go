package vod

import (
	"api/internal/utils/vod/aliyun"
	"api/internal/utils/vod/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

var (
	vodMap     = map[string]model.Vod{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	vodMuMap   sync.Map
	vodFuncMap = map[string]model.VodFunc{
		`vod_of_aliyun`: aliyun.NewVod,
	}
	vodTypeDef = `vod_of_aliyun`
)

func NewVod(ctx context.Context, vodType string, config map[string]any) (vod model.Vod) {
	vodKey := vodType + gmd5.MustEncrypt(config)
	ok := false
	if vod, ok = vodMap[vodKey]; ok { //先读一次（不加锁）
		return
	}
	muTmp, _ := vodMuMap.LoadOrStore(vodKey, &sync.Mutex{})
	mu := muTmp.(*sync.Mutex)
	mu.Lock()
	defer func() {
		mu.Unlock()
		vodMuMap.Delete(vodKey)
	}()
	if vod, ok = vodMap[vodKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	if _, ok = vodFuncMap[vodType]; !ok {
		vodType = vodTypeDef
	}
	vod = vodFuncMap[vodType](ctx, config)
	vodMap[vodKey] = vod
	return
}
