package initialize

import (
	"api/internal/consts"
	daoAuth "api/internal/dao/auth"
	daoPay "api/internal/dao/pay"
	daoUpload "api/internal/dao/upload"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
)

func initCron(ctx context.Context) {
	myCronThis := myCron{}
	myCronThis.SetEnv(ctx)                                      //必须先执行一次，在内存中初始化环境变量
	gcron.Add(ctx, `0 */15 * * * *`, myCronThis.SetEnv, `Test`) //15分钟更新一次。所有服务器都需要启动该定时器

	if !utils.IsDev(ctx) && g.Cfg().MustGet(ctx, `cronServerNetworkIp`).String() != g.Cfg().MustGetWithEnv(ctx, consts.SERVER_NETWORK_IP).String() {
		return
	}

	// gcron.AddSingleton(ctx, `*/5 * * * * *`, myCronThis.Test, `Test`)
}

type myCron struct{}

func (myCron) SetEnv(ctx context.Context) {
	daoAuth.Scene.CacheSet(ctx)

	daoUpload.Upload.CacheSet(ctx)

	daoPay.Scene.CacheSet(ctx)
	daoPay.Channel.CacheSet(ctx)
	daoPay.Pay.CacheSet(ctx)
}

func (myCron) Test(ctx context.Context) {
	fmt.Println(gtime.Now().Format(`Y-m-d H:i:s`))
}
