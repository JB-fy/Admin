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
	vodMu      sync.Mutex
	vodTypeDef = `vodOfAliyun`
	vodFuncMap = map[string]model.VodFunc{
		`vodOfAliyun`: aliyun.NewVod,
	}
)

func NewVod(ctx context.Context, vodType string, config map[string]any) (vod model.Vod) {
	vodKey := vodType + gmd5.MustEncrypt(config)
	ok := false
	if vod, ok = vodMap[vodKey]; ok { //先读一次（不加锁）
		return
	}
	vodMu.Lock()
	defer vodMu.Unlock()
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
