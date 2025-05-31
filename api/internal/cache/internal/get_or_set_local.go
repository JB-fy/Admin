package internal

import (
	"context"
	"sync"
)

var (
	GetOrSetLocal = getOrSetLocal{}
)

type getOrSetLocal struct {
	muMap sync.Map //存放所有缓存KEY的锁（当前服务器用）
}

func (cacheThis *getOrSetLocal) GetOrSetLocal(ctx context.Context, key string, setFunc func() (value any, notExist bool, err error), getFunc func() (value any, notExist bool, err error)) (value any, notExist bool, err error) {
	value, notExist, err = getFunc()
	if !notExist || err != nil { //先读一次（不加锁）
		return
	}

	// 防止当前服务器并发
	muTmp, _ := cacheThis.muMap.LoadOrStore(key, &sync.Mutex{})
	mu := muTmp.(*sync.Mutex)
	mu.Lock()
	defer mu.Unlock()
	value, notExist, err = getFunc() // 再读一次（加锁），防止重复初始化
	if !notExist || err != nil {
		return
	}

	value, notExist, err = setFunc()
	return
}
