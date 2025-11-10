package common

import (
	"api/internal/consts"
	"api/internal/utils/jbredis"
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

var GetOrSet = getOrSet{goCache: cache.New(0, consts.CACHE_LOCAL_INTERVAL_MINUTE*time.Minute)}

type getOrSet struct {
	goCache *cache.Cache
	muMap   sync.Map //存放所有缓存KEY的锁（当前服务器用）
}

func (cacheThis *getOrSet) cache() redis.UniversalClient {
	return jbredis.DB()
}

func (cacheThis *getOrSet) key(key string) string {
	return fmt.Sprintf(consts.CACHE_IS_SET, key)
}

func (cacheThis *getOrSet) retryInfo(numLock, numRead uint8, oneTime time.Duration) (uint8, uint8, time.Duration, time.Duration) {
	if numLock == 0 {
		numLock = 3
	}
	if numRead == 0 {
		numRead = 5
	}
	if oneTime <= 0 {
		oneTime = 200 * time.Millisecond
	}
	return numLock, numRead, oneTime, time.Duration(numLock*numRead) * oneTime
}

// 依据以下两个时间设置以下3个参数：valFunc运行速度 + 缓存写入到能读取之间的时间
// numLock	获取锁的重试次数。作用：当获取锁的服务器因报错，无法做缓存写入时，允许其它服务器重新获取锁，以保证缓存写入成功
// numRead	读取缓存的重试次数。作用：当未获取锁的服务器获取缓存时数据为空时，可以重试多次
// oneTime	读取缓存重试间隔时间，单位：毫秒
func (cacheThis *getOrSet) GetOrSet(ctx context.Context, key string, setFunc func() (value any, notExist bool, err error), getFunc func() (value any, notExist bool, err error), numLock, numRead uint8, oneTime time.Duration) (value any, notExist bool, err error) {
	value, notExist, err = getFunc()
	if !notExist || err != nil {
		return
	}

	// 防止当前服务器并发
	muTmp, _ := cacheThis.muMap.LoadOrStore(key, &sync.Mutex{})
	mu := muTmp.(*sync.Mutex)
	mu.Lock()
	defer func() {
		mu.Unlock()
		cacheThis.muMap.Delete(key)
	}()
	isSetKey := cacheThis.key(key)
	if _, isSetOfLocal := cacheThis.goCache.Get(isSetKey); isSetOfLocal { //当有协程（一般是第一个上锁成功的协程）执行setFunc设置缓存 或 发现缓存已存在 后，后续协程即可直接执行getFunc获取缓存
		value, notExist, err = getFunc()
		if !notExist || err != nil {
			return
		}
	}

	// 防止不同服务器并发
	numLock, numRead, oneTime, isSetTtl := cacheThis.retryInfo(numLock, numRead, oneTime)
	var isSet bool
	for range numLock {
		isSet, err = cacheThis.cache().SetNX(ctx, isSetKey, ``, isSetTtl).Result()
		if err != nil {
			return
		}
		if isSet {
			defer func() {
				if rec := recover(); rec != nil { //防止panic导致redis锁长时间没释放，造成频繁执行getFunc()方法
					err = errors.New(`设置缓存panic错误：` + gconv.String(rec) + `。栈信息：` + string(debug.Stack()))
					cacheThis.cache().Del(ctx, isSetKey) //报错时，删除redis锁缓存Key，允许其它服务器重新尝试设置缓存
					g.Log().Error(ctx, err.Error())
				} else if ctxErr := ctx.Err(); ctxErr != nil { //上下文取消或超时，删除redis锁缓存Key，允许其它服务器重新尝试设置缓存
					cacheThis.cache().Del(ctx, isSetKey)
					g.Log().Error(ctx, `上下文取消或超时：`+ctxErr.Error())
				}
			}()
			value, notExist, err = setFunc()
			if notExist || err != nil {
				cacheThis.cache().Del(ctx, isSetKey) //报错时，删除redis锁缓存Key，允许其它服务器重新尝试设置缓存
				return
			}
			cacheThis.goCache.Set(isSetKey, struct{}{}, isSetTtl)
			cacheThis.cache().Expire(ctx, isSetKey, isSetTtl)
			return
		}
		// 等待读取数据
		for range numRead {
			value, notExist, err = getFunc()
			if err != nil {
				return
			}
			if !notExist /* || err != nil */ {
				pttl, _ := cacheThis.cache().TTL(ctx, isSetKey).Result()
				cacheThis.goCache.Set(isSetKey, struct{}{}, time.Until(time.Now().Add(pttl)))
				return
			}
			// 放for前面执行。坏处：首次读取缓存有延迟；好处：减少缓存压力
			// 放for后面执行。好处：首次读取缓存没有延迟；坏处：增加缓存压力（通过当前服务器 上锁【getOrSetMu】 和 缓存【goCache】 可保证不会对缓存造成压力，除非服务器数量庞大）
			time.Sleep(oneTime)
		}
	}
	/*
		出现该错误的情况：
			1：所有服务器执行方法时都报错了。一般不大可能出现这种情况，概率极低
			2：redis服务负载过高，需要及时做优化了。解决方案：扩容或分库
	*/
	err = errors.New(`尝试多次查询缓存失败：` + key)
	return
}

// 删除时需同时删除redis竞争锁。建议：调用GetOrSet方法的缓存删除时也使用该方法。在缓存-删除-重设缓存三个步骤连续执行时，在第三步重设缓存会因redis竞争锁未删除报错：尝试多次查询缓存失败
func (cacheThis *getOrSet) Del(ctx context.Context, keyArr ...string) {
	isSetKeyArr := make([]string, len(keyArr))
	for index := range keyArr {
		isSetKeyArr[index] = cacheThis.key(keyArr[index])
		cacheThis.goCache.Delete(isSetKeyArr[index])
	}
	cacheThis.cache().Del(ctx, isSetKeyArr...)
}

// 错误最大缓存时间
func (cacheThis *getOrSet) MaxTtlOfErr(numLock, numRead uint8, oneTime time.Duration, defTtl time.Duration) time.Duration {
	_, _, _, isSetTtl := cacheThis.retryInfo(numLock, numRead, oneTime)
	if isSetTtl < defTtl {
		return defTtl
	}
	return isSetTtl
}
