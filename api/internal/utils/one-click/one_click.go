package one_click

import (
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils/one-click/wx"
	"api/internal/utils/one-click/yidun"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"golang.org/x/sync/singleflight"
)

var (
	oneClickOfWxMap sync.Map
	oneClickOfWxSfg singleflight.Group
)

func NewOneClickOfWx(ctx context.Context) (obj *wx.OneClick) {
	config := daoPlatform.Config.Get(ctx, `one_click_of_wx`).Map()
	key := gmd5.MustEncrypt(config)
	objTmp, ok := oneClickOfWxMap.Load(key)
	if !ok {
		objTmp, _, _ = oneClickOfWxSfg.Do(key, func() (obj any, err error) {
			obj = wx.NewOneClick(ctx, config)
			oneClickOfWxMap.Store(key, obj)
			return
		})
	}
	obj = objTmp.(*wx.OneClick)
	return
}

var (
	oneClickOfYidunMap sync.Map
	oneClickOfYidunSfg singleflight.Group
)

func NewOneClickOfYidun(ctx context.Context) (obj *yidun.OneClick) {
	config := daoPlatform.Config.Get(ctx, `one_click_of_yidun`).Map()
	key := gmd5.MustEncrypt(config)
	objTmp, ok := oneClickOfYidunMap.Load(key)
	if !ok {
		objTmp, _, _ = oneClickOfYidunSfg.Do(key, func() (obj any, err error) {
			obj = yidun.NewOneClick(ctx, config)
			oneClickOfYidunMap.Store(key, obj)
			return
		})
	}
	obj = objTmp.(*yidun.OneClick)
	return
}
