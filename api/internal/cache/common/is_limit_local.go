package common

import (
	"api/internal/utils"
	"context"
	"sync"
	"time"
)

var IsLimitLocal = isLimitLocal{
	chanMap: map[string]chan struct{}{},
}

type isLimitLocal struct {
	chanMap   map[string]chan struct{}
	chanMuMap sync.Map
}

func (cacheThis *isLimitLocal) GetChan(key string, size uint) (ch chan struct{}) {
	chanKey := key
	// chanKey := fmt.Sprintf(`%s:%d`, key, size)
	ok := false
	if ch, ok = cacheThis.chanMap[chanKey]; ok { //先读一次（不加锁）
		return
	}
	muTmp, _ := cacheThis.chanMuMap.LoadOrStore(chanKey, &sync.Mutex{})
	mu := muTmp.(*sync.Mutex)
	mu.Lock()
	defer func() {
		mu.Unlock()
		cacheThis.chanMuMap.Delete(chanKey)
	}()
	if ch, ok = cacheThis.chanMap[chanKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	ch = make(chan struct{}, size)
	cacheThis.chanMap[chanKey] = ch
	return
}

func (cacheThis *isLimitLocal) Acquire(ctx context.Context, ch chan struct{}, waitTime time.Duration) (err error) {
	if waitTime > 0 {
		timer := time.NewTimer(waitTime)
		defer timer.Stop()
		select {
		case ch <- struct{}{}:
		case <-ctx.Done():
			err = utils.NewErrorCode(ctx, 99999998, ctx.Err().Error())
		case <-timer.C:
			err = utils.NewErrorCode(ctx, 99999998, ``)
		}
		return
	}
	select {
	case ch <- struct{}{}:
	case <-ctx.Done():
		err = utils.NewErrorCode(ctx, 99999998, ctx.Err().Error())
	default:
		err = utils.NewErrorCode(ctx, 99999998, ``)
	}
	return
}

func (cacheThis *isLimitLocal) Release(ctx context.Context, ch chan struct{}) {
	select {
	case <-ch:
	default: //重复调用 或 管道为空时忽略，防止堵塞
	}
}
