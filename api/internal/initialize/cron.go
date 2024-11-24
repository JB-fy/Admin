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
	/*--------数据库中某些配置表极少修改，统一缓存在本机内存中，能极大增加服务器性能，减少数据库压力（注意：服务启动时，就必须先运行一次，缓存到内存中） 开始--------*/
	daoAuth.Scene.CacheSet(ctx)
	gcron.AddSingleton(ctx, `50 0 3 * * *`, daoAuth.Scene.CacheSet, `AuthSceneCacheSet`) //每天晚上3点刷新一次
	// 表数据很小，无需这样做，且会导致数据修改无法立即生效。确实需要减轻数据库压力时可以使用
	// daoAuth.Menu.CacheSet(ctx)
	// gcron.AddSingleton(ctx, `50 */15 * * * *`, daoAuth.Menu.CacheSet, `AuthMenuCacheSet`) //每15分钟刷新一次
	// daoAuth.Action.CacheSet(ctx)
	// gcron.AddSingleton(ctx, `50 */15 * * * *`, daoAuth.Action.CacheSet, `AuthActionCacheSet`) //每15分钟刷新一次

	daoUpload.Upload.CacheSet(ctx)
	gcron.AddSingleton(ctx, `40 */30 * * * *`, daoUpload.Upload.CacheSet, `UploadCacheSet`) //每30分钟刷新一次

	myCronThis.PayCacheSet(ctx)
	gcron.AddSingleton(ctx, `30 */15 * * * *`, myCronThis.PayCacheSet, `PayCacheSet`) //每15分钟刷新一次
	/*--------数据库中某些配置表极少修改，统一缓存在本机内存中，能极大增加服务器性能，减少数据库压力（注意：服务启动时，就必须先运行一次，缓存到内存中） 结束--------*/

	// 部分定时任务不允许全部服务器都开启，只有指定IP的服务器才能开启。比如任务存在数据库先读后改的逻辑时，多服务器同时开启任务，会存在重复处理的问题
	if !utils.IsDev(ctx) && g.Cfg().MustGet(ctx, `cronServerNetworkIp`).String() != g.Cfg().MustGetWithEnv(ctx, consts.SERVER_NETWORK_IP).String() {
		return
	}

	// gcron.AddSingleton(ctx, `*/5 * * * * *`, myCronThis.Test, `Test`)
}

type myCron struct{}

func (myCron) PayCacheSet(ctx context.Context) {
	daoPay.Scene.CacheSet(ctx)
	daoPay.Channel.CacheSet(ctx)
	daoPay.Pay.CacheSet(ctx)
}

func (myCron) Test(ctx context.Context) {
	fmt.Println(gtime.Now().Format(`Y-m-d H:i:s`))
}
