package common

import (
	"api/internal/consts"
	"api/internal/utils/jbredis"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/redis/go-redis/v9"
)

var IsRun = isRun{}

type isRun struct{}

func (cacheThis *isRun) cache() redis.UniversalClient {
	return jbredis.DB()
}

func (cacheThis *isRun) key(key string) string {
	return fmt.Sprintf(consts.CACHE_IS_RUN, key)
}

func (cacheThis *isRun) IsRun(ctx context.Context, key string, ttl time.Duration) (isRun bool, runEndFunc func(), err error) {
	isRunKey := cacheThis.key(key)
	isRun, err = cacheThis.cache().SetNX(ctx, isRunKey, ``, ttl).Result()
	if err != nil {
		return
	}
	if !isRun {
		return
	}
	//保证操作执行完成之前，isRun不会过期
	timer := gtimer.AddSingleton(ctx, ttl-(3*time.Second), func(ctx context.Context) {
		cacheThis.cache().Expire(ctx, isRunKey, ttl)
	})
	runEndFunc = func() {
		timer.Close()
		cacheThis.cache().Del(ctx, isRunKey).Result()
	}
	return
}

func (cacheThis *isRun) Exist(ctx context.Context, key string) bool {
	row, _ := cacheThis.cache().Exists(ctx, cacheThis.key(key)).Result()
	return row > 0
}
