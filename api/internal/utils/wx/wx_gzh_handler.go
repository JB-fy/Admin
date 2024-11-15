package wx

import (
	daoPlatform "api/internal/dao/platform"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

var (
	wxGzhMap = map[string]*WxGzh{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	wxGzhMu  sync.Mutex
)

func NewWxGzhHandler(ctx context.Context) (wxGzh *WxGzh) {
	config, _ := daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `wxGzh`).ValueMap(daoPlatform.Config.Columns().ConfigValue)

	wxGzhKey := gmd5.MustEncrypt(config)

	ok := false
	if wxGzh, ok = wxGzhMap[wxGzhKey]; ok { //先读一次（不加锁）
		return
	}
	wxGzhMu.Lock()
	defer wxGzhMu.Unlock()
	if wxGzh, ok = wxGzhMap[wxGzhKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}

	wxGzh = NewWxGzh(config)
	wxGzhMap[wxGzhKey] = wxGzh
	return NewWxGzh(config)
}
