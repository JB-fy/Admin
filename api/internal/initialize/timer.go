package initialize

import (
	"api/internal/cache"
	"api/internal/utils/wx"
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
)

func initTimer(ctx context.Context) {
	if !g.Cfg().MustGet(ctx, `dev`).Bool() {
		if g.Cfg().MustGet(ctx, `timerServerNetworkIp`).String() != g.Cfg().MustGetWithEnv(ctx, `SERVER_NETWORK_IP`).String() {
			return
		}
	}

	// myTimerThis := myTimer{}
	// myTimerThis.CacheWxGzhAccessToken(ctx) //缓存微信公众号AccessToken（需要时启用，且公众号需设置IP白名单）
}

type myTimer struct{}

func (myTimerThis *myTimer) CacheWxGzhAccessToken(ctx context.Context) {
	wxGzhObj := wx.NewWxGzhByPfCfg(ctx)
	accessTokenInfo, err := wxGzhObj.AccessToken()
	if err != nil {
		g.Log().Error(ctx, `获取微信公众号AccessToken接口错误：`+err.Error(), err)
		gtimer.SetTimeout(ctx, 5*time.Second, myTimerThis.CacheWxGzhAccessToken) //5秒后重试
		return
	}
	err = cache.NewWxGzhAccessToken(ctx, wxGzhObj.AppId).Set(accessTokenInfo.AccessToken, int64(accessTokenInfo.ExpiresIn))
	if err != nil {
		g.Log().Error(ctx, `缓存微信公众号AccessToken错误：`+err.Error(), err)
		gtimer.SetTimeout(ctx, 5*time.Second, myTimerThis.CacheWxGzhAccessToken) //5秒后重试
		return
	}
	gtimer.SetTimeout(ctx, time.Duration(accessTokenInfo.ExpiresIn-30)*time.Second, myTimerThis.CacheWxGzhAccessToken) //提早30秒刷新缓存
}
