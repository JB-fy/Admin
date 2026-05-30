package initialize

import (
	cacheCommon "api/internal/cache/common"
	"api/internal/utils/jbredis"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func initRedis(ctx context.Context) {
	for group, config := range g.Cfg().MustGet(ctx, `redisDb`).Map() {
		jbredis.AddDB(ctx, group, gconv.Map(config))
	}
	cacheCommon.InitIsLimit(ctx)

	/* // 只有指定IP的服务器才继续执行，否则会造成重复处理
	if !(utils.IsDev(ctx) || g.Cfg().MustGet(ctx, `masterServerNetworkIpArr.0`).String() == genv.Get(consts.ENV_SERVER_NETWORK_IP).String()) {
		return
	}
	sub.Add(ctx, `default`, jbredis.DB(), `__keyevent@0__:expired`) */
}
