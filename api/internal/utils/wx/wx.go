package wx

import (
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils/wx/gzh"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"golang.org/x/sync/singleflight"
)

var (
	wxGzhMap sync.Map
	wxGzhSfg singleflight.Group
)

func NewWxGzh(ctx context.Context) (obj *gzh.Wx) {
	config := daoPlatform.Config.Get(ctx, `wx_gzh`).Map()
	key := gmd5.MustEncrypt(config)
	objTmp, ok := wxGzhMap.Load(key)
	if !ok {
		objTmp, _, _ = wxGzhSfg.Do(key, func() (obj any, err error) {
			obj = gzh.NewWx(ctx, config)
			wxGzhMap.Store(key, obj)
			return
		})
	}
	obj = objTmp.(*gzh.Wx)
	return
}
