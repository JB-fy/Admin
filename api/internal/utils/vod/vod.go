package vod

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

type VodParam struct {
	ExpireTime int64 //签名有效时间。单位：秒
}

type Vod interface {
	Sts(ctx context.Context, param VodParam) (stsInfo map[string]any, err error) // 获取Sts Token
}

var (
	vodTypeDef = `vodOfAliyun`
	vodFuncMap = map[string]func(ctx context.Context, config map[string]any) Vod{
		`vodOfAliyun`: func(ctx context.Context, config map[string]any) Vod { return NewVodOfAliyun(ctx, config) },
	}
	vodMap = map[string]Vod{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	vodMu  sync.Mutex
)

func NewVod(ctx context.Context, vodType string, config map[string]any) (vod Vod) {
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
