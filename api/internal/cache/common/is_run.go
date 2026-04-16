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

func (cacheThis *isRun) IsRun(ctx context.Context, key string, ttl time.Duration) (isRun bool, runEndFunc func(isDel bool), err error) {
	// return cacheThis.cache().SetNX(ctx, cacheThis.key(key), ``, ttl).Result()
	isRunKey := cacheThis.key(key)
	isRun, err = cacheThis.cache().SetNX(ctx, isRunKey, ``, ttl).Result()
	if !isRun || err != nil {
		return
	}
	runEndFunc = func(isDel bool) {
		if isDel {
			cacheThis.cache().Del(ctx, isRunKey).Result()
		}
	}
	return
}

// 运行耗时任务，需定时刷新缓存时间，防止运行期间，锁过期失效
func (cacheThis *isRun) IsRunAndRefreshTTL(ctx context.Context, key string, ttl time.Duration, checkRunResultFuncOpt ...func(ctx context.Context) (isResult bool, err error)) (isRun, isResult bool, runEndFunc func(isDel bool), err error) {
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
		cacheThis.cache().PExpire(ctx, isRunKey, ttl)
	})
	runEndFunc = func(isDel bool) {
		timer.Close()
		if isDel {
			cacheThis.cache().Del(ctx, isRunKey).Result()
		}
	}
	return
}

func (cacheThis *isRun) IsRunAndRefreshTTLWithRunFunc(ctx context.Context, key string, ttl time.Duration, isDel bool, isRunGo bool, waitTimeGo time.Duration, runFunc func(ctx context.Context) (err error), checkRunResultFuncOpt ...func(ctx context.Context) (isResult bool, err error)) (isRun, isResult bool, err error) {
	isRun, isResult, runEndFunc, err := cacheThis.IsRunAndRefreshTTL(ctx, key, ttl, checkRunResultFuncOpt...)
	if !isRun || isResult || err != nil {
		return
	}
	if !isRunGo {
		defer func() {
			runEndFunc(isDel)
		}()
		err = runFunc(ctx)
		return
	}
	if waitTimeGo <= 0 {
		go func() {
			defer func() {
				runEndFunc(isDel)
			}()
			runFunc(ctx)
		}()
		return
	}
	ctxOfWait, cancel := context.WithTimeout(ctx, waitTimeGo)
	defer cancel()
	ch := make(chan error, 1)
	go func() {
		defer func() {
			runEndFunc(isDel)
		}()
		ch <- runFunc(ctx)
	}()
	select {
	case err = <-ch:
	case <-ctxOfWait.Done():
	}
	return
}

func (cacheThis *isRun) Exist(ctx context.Context, key string) bool {
	row, _ := cacheThis.cache().Exists(ctx, cacheThis.key(key)).Result()
	return row > 0
}

func (cacheThis *isRun) Del(ctx context.Context, key string) (int64, error) {
	return cacheThis.cache().Del(ctx, cacheThis.key(key)).Result()
}
