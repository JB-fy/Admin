package common

import (
	"context"

	"golang.org/x/sync/singleflight"
)

var (
	GetOrSetLocal = getOrSetLocal{}
)

type getOrSetLocal struct {
	sfg singleflight.Group
}

type localResult struct {
	value    any
	notExist bool
}

func (cacheThis *getOrSetLocal) GetOrSetLocal(ctx context.Context, key string, setFunc func() (value any, notExist bool, err error), getFunc func() (value any, notExist bool, err error)) (value any, notExist bool, err error) {
	value, notExist, err = getFunc()
	if !notExist || err != nil {
		return
	}
	resultTmp, err, _ := cacheThis.sfg.Do(key, func() (result any, err error) {
		value, notExist, err := setFunc()
		result = &localResult{value: value, notExist: notExist}
		return
	})
	result := resultTmp.(*localResult)
	value = result.value
	notExist = result.notExist
	return
}
