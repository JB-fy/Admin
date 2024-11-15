package vod

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

type VodParam struct {
	ExpireTime int64 //签名有效时间。单位：秒
}

type Vod interface {
	Sts(ctx context.Context, param VodParam) (stsInfo map[string]any, err error) // 获取Sts Token
}

var (
	vodMap = map[string]Vod{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	vodMu  sync.Mutex
)

func NewVod(config map[string]any) (vod Vod) {
	vodKey := gmd5.MustEncrypt(config)

	ok := false
	if vod, ok = vodMap[vodKey]; ok { //先读一次（不加锁）
		return
	}
	vodMu.Lock()
	defer vodMu.Unlock()
	if vod, ok = vodMap[vodKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}

	switch gconv.String(config[`vodType`]) {
	// case `vodOfAliyun`:
	default:
		vod = NewVodOfAliyun(config)
	}
	vodMap[vodKey] = vod
	return
}
