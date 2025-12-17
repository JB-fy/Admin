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
}
