package initialize

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func initDb(ctx context.Context) {
	for k := range g.Cfg().MustGet(ctx, `database`).Map() {
		switch k {
		case `logger`:
			continue
		default:
			// g.DB(k).GetCache().SetAdapter(gcache.NewAdapterRedis(g.Redis()))
		}
	}
}
