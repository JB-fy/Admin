package initialize

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

func initDb(ctx context.Context) {
	dbList := g.Cfg().MustGet(ctx, `database`).Map()
	redis := gcache.NewAdapterRedis(g.Redis())
	for k := range dbList {
		switch k {
		case `logger`:
			continue
		default:
			g.DB(k).GetCache().SetAdapter(redis)
		}
	}
}
