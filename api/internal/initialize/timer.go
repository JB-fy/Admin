package initialize

import (
	"api/internal/cache"
	one_click "api/internal/utils/one-click"
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
)

func initTimer(ctx context.Context) {
	if !g.Cfg().MustGet(ctx, `dev`).Bool() {
		if g.Cfg().MustGet(ctx, `cronServerNetworkIp`).String() != g.Cfg().MustGetWithEnv(ctx, `SERVER_NETWORK_IP`).String() {
			return
		}
	}

	// myTimerThis := myTimer{}
	// myTimerThis.CacheCgiAccessTokenOfWx(ctx) //缓存微信授权token（需要时启用）
}

type myTimer struct{}

func (myTimerThis *myTimer) CacheCgiAccessTokenOfWx(ctx context.Context) {
	oneClickObj := one_click.NewOneClickOfWx(ctx)
	accessTokenInfo, err := oneClickObj.CgiAccessToken()
	if err != nil {
		gtimer.SetTimeout(ctx, 5*time.Second, myTimerThis.CacheCgiAccessTokenOfWx) //5秒后重试
		return
	}
	err = cache.NewCgiAccessTokenOfWx(ctx, oneClickObj.AppId).Set(accessTokenInfo.AccessToken, int64(accessTokenInfo.ExpiresIn))
	if err != nil {
		gtimer.SetTimeout(ctx, 5*time.Second, myTimerThis.CacheCgiAccessTokenOfWx) //5秒后重试
		return
	}
	gtimer.SetTimeout(ctx, time.Duration(accessTokenInfo.ExpiresIn-30)*time.Second, myTimerThis.CacheCgiAccessTokenOfWx) //提早30秒刷新缓存
}
