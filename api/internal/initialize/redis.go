package initialize

import (
	// daoCh "api/internal/dao/ch"
	// daoAuth "api/internal/dao/auth"
	// daoOrg "api/internal/dao/org"
	// daoPlt "api/internal/dao/plt"
	"api/internal/utils/redis"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func initRedis(ctx context.Context) {
	for group, config := range g.Cfg().MustGet(ctx, `redisDb`).Map() {
		redis.AddDB(ctx, group, gconv.Map(config))
	}
}
