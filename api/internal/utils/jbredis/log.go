package jbredis

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/redis/go-redis/v9"
)

type HookLog struct {
	Group  string
	Config *redis.UniversalOptions
	Log    *glog.Logger
}

func (hook HookLog) DialHook(next redis.DialHook) redis.DialHook {
	return next
	/* return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	} */
}

func (hook HookLog) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	// return next
	return func(ctx context.Context, cmd redis.Cmder) error {
		startTime := time.Now()
		defer func() {
			hook.Log.Debug(ctx, fmt.Sprintf(`[REDIS] [%.3f ms] [%s] [%d] %s`,
				float64(time.Since(startTime).Microseconds())/1000,
				hook.Group,
				hook.Config.DB,
				cmd,
			))
		}()
		return next(ctx, cmd)
	}
}

func (hook HookLog) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	// return next
	return func(ctx context.Context, cmds []redis.Cmder) error {
		startTime := time.Now()
		defer func() {
			hook.Log.Debug(ctx, fmt.Sprintf(`[REDIS] [BATCH] [%s] [%d] %s`,
				hook.Group,
				hook.Config.DB,
				`BATCH START`,
			))
			for _, cmd := range cmds {
				hook.Log.Debug(ctx, fmt.Sprintf(`[REDIS] [BATCH] [%s] [%d] %s`,
					hook.Group,
					hook.Config.DB,
					cmd,
				))
			}
			hook.Log.Debug(ctx, fmt.Sprintf(`[REDIS] [BATCH] [%.3f ms] [%s] [%d] %s`,
				float64(time.Since(startTime).Microseconds())/1000,
				hook.Group,
				hook.Config.DB,
				`BATCH END`,
			))
		}()
		return next(ctx, cmds)
	}
}
