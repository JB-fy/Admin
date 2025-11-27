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

var IsRun = isRun{addSecond: 5 * time.Second}

type isRun struct{ addSecond time.Duration }

func (cacheThis *isRun) cache() redis.UniversalClient {
	return jbredis.DB()
}

func (cacheThis *isRun) key(key string) string {
	return fmt.Sprintf(consts.CACHE_IS_RUN, key)
}

func (cacheThis *isRun) IsRunNotRunFunc(ctx context.Context, key string, ttl time.Duration, checkRunResultFuncOpt ...func() (isRun bool, err error)) (isRun bool, runEndFunc func(), err error) {
	if ttl == 0 {
		ttl = 10 * time.Second
	} else if ttl < cacheThis.addSecond { //不能小于addSecond
		ttl = ttl + cacheThis.addSecond
	}
	isRunKey := cacheThis.key(key)
	if len(checkRunResultFuncOpt) > 0 && checkRunResultFuncOpt[0] != nil {
		row, _ := cacheThis.cache().Exists(ctx, isRunKey).Result()
		if row > 0 {
			return
		}
		isRun, err = checkRunResultFuncOpt[0]()
		if isRun || err != nil {
			return
		}
	}
	isRun, err = cacheThis.cache().SetNX(ctx, isRunKey, ``, ttl).Result()
	if !isRun || err != nil {
		return
	}
	//保证操作执行完成之前，isRun不会过期
	timer := gtimer.AddSingleton(ctx, ttl-cacheThis.addSecond, func(ctx context.Context) {
		cacheThis.cache().Expire(ctx, isRunKey, ttl)
	})
	runEndFunc = func() {
		timer.Close()
		cacheThis.cache().Del(ctx, isRunKey).Result()
	}
	return
}

func (cacheThis *isRun) IsRun(ctx context.Context, key string, ttl time.Duration, runFunc func() (err error), checkRunResultFuncOpt ...func() (isResult bool, err error)) (isRun bool, err error) {
	isRun, runEndFunc, err := cacheThis.IsRunNotRunFunc(ctx, key, ttl, checkRunResultFuncOpt...)
	if !isRun || err != nil {
		return
	}
	err = runFunc()
	runEndFunc()
	return
}

func (cacheThis *isRun) Exist(ctx context.Context, key string) bool {
	row, _ := cacheThis.cache().Exists(ctx, cacheThis.key(key)).Result()
	return row > 0
}
