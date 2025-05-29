package one_click

import (
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils/one-click/wx"
	"api/internal/utils/one-click/yidun"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

var (
	oneClickOfWxMap = map[string]*wx.OneClick{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	oneClickOfWxMu  sync.Mutex
)

func NewOneClickOfWx(ctx context.Context) (oneClickOfWx *wx.OneClick) {
	config := daoPlatform.Config.GetOne(ctx, `one_click_of_wx`).Map()
	oneClickOfWxKey := gmd5.MustEncrypt(config)
	ok := false
	if oneClickOfWx, ok = oneClickOfWxMap[oneClickOfWxKey]; ok { //先读一次（不加锁）
		return
	}
	oneClickOfWxMu.Lock()
	defer oneClickOfWxMu.Unlock()
	if oneClickOfWx, ok = oneClickOfWxMap[oneClickOfWxKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	oneClickOfWx = wx.NewOneClick(ctx, config)
	oneClickOfWxMap[oneClickOfWxKey] = oneClickOfWx
	return
}

var (
	oneClickOfYidunMap = map[string]*yidun.OneClick{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	oneClickOfYidunMu  sync.Mutex
)

func NewOneClickOfYidun(ctx context.Context) (oneClickOfYidun *yidun.OneClick) {
	config := daoPlatform.Config.GetOne(ctx, `one_click_of_yidun`).Map()
	oneClickOfYidunKey := gmd5.MustEncrypt(config)
	ok := false
	if oneClickOfYidun, ok = oneClickOfYidunMap[oneClickOfYidunKey]; ok { //先读一次（不加锁）
		return
	}
	oneClickOfYidunMu.Lock()
	defer oneClickOfYidunMu.Unlock()
	if oneClickOfYidun, ok = oneClickOfYidunMap[oneClickOfYidunKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	oneClickOfYidun = yidun.NewOneClick(ctx, config)
	oneClickOfYidunMap[oneClickOfYidunKey] = oneClickOfYidun
	return
}
