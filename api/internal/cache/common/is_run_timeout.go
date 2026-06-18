package common

import (
	"api/internal/utils"
	"context"
	"time"
)

var IsRunTimeout = isRunTimeout{}

type isRunTimeout struct{}

type runResult struct {
	result any
	err    error
}

func (cacheThis *isRunTimeout) IsRunTimeout(ctx context.Context, waitTime time.Duration, runFunc func(ctx context.Context) (value any, err error)) (isRunTimeout bool, result any, err error) {
	ch := make(chan *runResult, 1)
	timer := time.NewTimer(waitTime)
	defer timer.Stop()
	go func() {
		runResult := &runResult{}
		runResult.result, runResult.err = runFunc(ctx)
		ch <- runResult
	}()
	select {
	case runResult := <-ch:
		result = runResult.result
		err = runResult.err
	case <-timer.C:
		isRunTimeout = true
		err = utils.NewErrorCode(ctx, 99990000, ``)
	case <-ctx.Done():
		err = utils.NewErrorCode(ctx, 99999997, ctx.Err().Error())
	}
	return
}
