package internal

import (
	"api/internal/consts"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	GetOrSet      = getOrSet{redis: g.Redis()}
	getOrSetMuMap sync.Map //存放所有缓存KEY的锁（当前服务器用）
)

type getOrSet struct{ redis *gredis.Redis }

func (cacheThis *getOrSet) cache() *gredis.Redis {
	return cacheThis.redis
}

func (cacheThis *getOrSet) key(key string) string {
	return fmt.Sprintf(consts.CACHE_IS_SET, key)
}

// 依据以下两个时间设置以下3个参数：valFunc运行速度 + 缓存写入到能读取之间的时间
// numLock	获取锁的重试次数。作用：当获取锁的服务器因报错，无法做缓存写入时，允许其它服务器重新获取锁，以保证缓存写入成功
// numRead	读取缓存的重试次数。作用：当未获取锁的服务器获取缓存时数据为空时，可以重试多次
// oneTime	读取缓存重试间隔时间，单位：毫秒
func (cacheThis *getOrSet) GetOrSet(ctx context.Context, key string, setFunc func() (value any, notExist bool, err error), getFunc func() (value any, notExist bool, err error), numLock int, numRead int, oneTime time.Duration) (value any, notExist bool, err error) {
	value, notExist, err = getFunc()
	if !notExist || err != nil {
		return
	}

	// 防止当前服务器并发
	getOrSetMuTmp, _ := getOrSetMuMap.LoadOrStore(key, &sync.Mutex{})
	getOrSetMu := getOrSetMuTmp.(*sync.Mutex)
	getOrSetMu.Lock()
	defer getOrSetMu.Unlock()

	// 防止不同服务器并发
	if numLock <= 0 {
		numLock = 3
	}
	if numRead <= 0 {
		numRead = 5
	}
	if oneTime <= 0 {
		oneTime = 200 * time.Millisecond
	}
	var isSetVal *gvar.Var
	isSetKey := cacheThis.key(key)
	isSetTtl := gconv.Int64(time.Duration(numLock*numRead) * oneTime / time.Second) //redis锁缓存Key时间
	isSetOption := gredis.SetOption{TTLOption: gredis.TTLOption{EX: &isSetTtl}, NX: true}
	for range numLock {
		isSetVal, err = cacheThis.cache().Set(ctx, isSetKey, ``, isSetOption)
		if err != nil {
			return
		}
		if isSetVal.Bool() {
			value, notExist, err = setFunc()
			if notExist || err != nil {
				cacheThis.cache().Del(ctx, isSetKey) //报错时，删除redis锁缓存Key，允许其它服务器重新尝试设置缓存
				return
			}
			cacheThis.cache().Expire(ctx, isSetKey, isSetTtl)
			return
		}
		// 等待读取数据
		for range numRead {
			time.Sleep(oneTime)
			value, notExist, err = getFunc()
			if !notExist || err != nil {
				return
			}
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

func (cacheThis *getOrSet) GetOrSetOfRedis(ctx context.Context, key string, setFunc func() (value any, notExist bool, err error), numLock int, numRead int, oneTime time.Duration) (value *gvar.Var, notExist bool, err error) {
	valueTmp, notExist, err := cacheThis.GetOrSet(ctx, key, func() (value any, notExist bool, err error) {
		value, notExist, err = setFunc()
		if notExist || err != nil {
			return
		}
		value = gvar.New(value)
		return
	}, func() (value any, notExist bool, err error) {
		value, err = cacheThis.cache().Get(ctx, key)
		if err != nil {
			return
		}
		notExist = value.(*gvar.Var).IsNil()
		return
	}, numLock, numRead, oneTime)
	value, _ = valueTmp.(*gvar.Var)
	return
}

// 删除时需同时删除redis竞争锁。建议：调用GetOrSet方法的缓存删除时也使用该方法。在缓存-删除-重设缓存三个步骤连续执行时，在第三步重设缓存会因redis竞争锁未删除报错：尝试多次查询缓存失败
func (cacheThis *getOrSet) Del(ctx context.Context, keyArr ...string) {
	isSetKeyArr := make([]string, len(keyArr))
	for index, key := range keyArr {
		isSetKeyArr[index] = cacheThis.key(key)
	}
	cacheThis.cache().Del(ctx, isSetKeyArr...)
}
