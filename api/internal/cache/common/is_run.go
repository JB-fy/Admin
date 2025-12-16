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

var IsRun = isRun{advSecond: 5 * time.Second}

type isRun struct{ advSecond time.Duration }

func (cacheThis *isRun) cache() redis.UniversalClient {
	return jbredis.DB()
}

func (cacheThis *isRun) key(key string) string {
	return fmt.Sprintf(consts.CACHE_IS_RUN, key)
}

func (cacheThis *isRun) IsRunNotRunFunc(ctx context.Context, key string, ttl time.Duration, checkRunResultFuncOpt ...func(ctx context.Context) (isResult bool, err error)) (isRun, isResult bool, runEndFunc func(), err error) {
	if ttl == 0 {
		ttl = 2 * cacheThis.advSecond
	} else if ttl <= cacheThis.advSecond { //不能小于advSecond
		ttl = ttl + cacheThis.advSecond
	}
	isRunKey := cacheThis.key(key)
	if len(checkRunResultFuncOpt) > 0 && checkRunResultFuncOpt[0] != nil {
		row, _ := cacheThis.cache().Exists(ctx, isRunKey).Result()
		if row > 0 {
			return
		}
		isResult, err = checkRunResultFuncOpt[0](ctx)
		if isResult || err != nil {
			return
		}
	}
	isRun, err = cacheThis.cache().SetNX(ctx, isRunKey, ``, ttl).Result()
	if !isRun || err != nil {
		return
	}
	//保证操作执行完成之前，isRun不会过期
	timer := gtimer.AddSingleton(ctx, ttl-cacheThis.advSecond, func(ctx context.Context) {
		cacheThis.cache().Expire(ctx, isRunKey, ttl)
	})
	runEndFunc = func() {
		timer.Close()
		cacheThis.cache().Del(ctx, isRunKey).Result()
	}
	return
}

func (cacheThis *isRun) IsRun(ctx context.Context, key string, ttl time.Duration, isRunGo bool, runFunc func(ctx context.Context) (err error), checkRunResultFuncOpt ...func(ctx context.Context) (isResult bool, err error)) (isRun, isResult bool, err error) {
	isRun, isResult, runEndFunc, err := cacheThis.IsRunNotRunFunc(ctx, key, ttl, checkRunResultFuncOpt...)
	if !isRun || isResult || err != nil {
		return
	}
	if isRunGo {
		go func() {
			defer func() { runEndFunc() }()
			runFunc(ctx)
		}()
		return
	}
	defer func() { runEndFunc() }()
	err = runFunc(ctx)
	return
}

func (cacheThis *isRun) Exist(ctx context.Context, key string) bool {
	row, _ := cacheThis.cache().Exists(ctx, cacheThis.key(key)).Result()
	return row > 0
}
