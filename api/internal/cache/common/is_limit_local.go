package common

import (
	"api/internal/utils"
	"context"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var IsLimitLocal = isLimitLocal{}

type isLimitLocal struct {
	chanMap sync.Map
	chanSfg singleflight.Group
}

func (cacheThis *isLimitLocal) GetChan(key string, size uint) (ch chan struct{}) {
	chanKey := key
	// chanKey := fmt.Sprintf(`%s:%d`, key, size)
	chTmp, ok := cacheThis.chanMap.Load(chanKey)
	if !ok {
		chTmp, _, _ = cacheThis.chanSfg.Do(chanKey, func() (ch any, err error) {
			ch = make(chan struct{}, size)
			cacheThis.chanMap.Store(chanKey, ch)
			return
		})
	}
	ch = chTmp.(chan struct{})
	return
}

func (cacheThis *isLimitLocal) Acquire(ctx context.Context, ch chan struct{}, waitTime time.Duration) (err error) {
	if waitTime > 0 {
		timer := time.NewTimer(waitTime)
		defer timer.Stop()
		select {
		case ch <- struct{}{}:
		case <-timer.C:
			err = utils.NewErrorCode(ctx, 99999998, ``)
		case <-ctx.Done():
			err = utils.NewErrorCode(ctx, 99999998, ctx.Err().Error())
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
