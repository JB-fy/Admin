package initialize

import (
	"api/internal/cache"
	"api/internal/consts"
	"api/internal/utils"
	"api/internal/utils/wx"
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gtimer"
)

func initTimer(ctx context.Context) {
	if !utils.IsDev(ctx) && g.Cfg().MustGet(ctx, `timerServerNetworkIp`).String() != genv.Get(consts.ENV_SERVER_NETWORK_IP).String() {
		return
	}

	// myTimerThis := myTimer{}
	// myTimerThis.CACHE_WX_GZH_ACCESS_TOKEN(ctx) //缓存微信公众号AccessToken（需要时启用，且公众号需设置IP白名单）
}

type myTimer struct{}

func (myTimerThis *myTimer) CACHE_WX_GZH_ACCESS_TOKEN(ctx context.Context) {
	wxGzhObj := wx.NewWxGzh(ctx)
	accessTokenInfo, err := wxGzhObj.AccessToken(ctx)
	if err != nil {
		g.Log().Error(ctx, `获取微信公众号AccessToken接口错误：`+err.Error(), err)
		gtimer.SetTimeout(ctx, 5*time.Second, myTimerThis.CACHE_WX_GZH_ACCESS_TOKEN) //5秒后重试
		return
	}
	err = cache.WxGzhAccessToken.Set(ctx, wxGzhObj.AppId, accessTokenInfo.AccessToken, accessTokenInfo.ExpiresIn)
	if err != nil {
		g.Log().Error(ctx, `缓存微信公众号AccessToken错误：`+err.Error(), err)
		gtimer.SetTimeout(ctx, 5*time.Second, myTimerThis.CACHE_WX_GZH_ACCESS_TOKEN) //5秒后重试
		return
	}
	gtimer.SetTimeout(ctx, time.Duration(accessTokenInfo.ExpiresIn-30)*time.Second, myTimerThis.CACHE_WX_GZH_ACCESS_TOKEN) //提早30秒刷新缓存
}
