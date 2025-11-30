package wx

import (
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils/wx/gzh"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

var (
	wxGzhMap   = map[string]*gzh.Wx{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	wxGzhMuMap sync.Map
)

func NewWxGzh(ctx context.Context) (wxGzh *gzh.Wx) {
	config := daoPlatform.Config.Get(ctx, `wx_gzh`).Map()
	wxGzhKey := gmd5.MustEncrypt(config)
	ok := false
	if wxGzh, ok = wxGzhMap[wxGzhKey]; ok { //先读一次（不加锁）
		return
	}
	muTmp, _ := wxGzhMuMap.LoadOrStore(wxGzhKey, &sync.Mutex{})
	mu := muTmp.(*sync.Mutex)
	mu.Lock()
	defer func() {
		mu.Unlock()
		wxGzhMuMap.Delete(wxGzhKey)
	}()
	if wxGzh, ok = wxGzhMap[wxGzhKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	wxGzh = gzh.NewWx(ctx, config)
	wxGzhMap[wxGzhKey] = wxGzh
	return
}
