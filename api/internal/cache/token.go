package cache

import (
	"api/internal/consts"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

type TokenActive struct {
	Ctx   context.Context
	Redis *gredis.Redis
	Key   string
}

func NewTokenActive(ctx context.Context, sceneCode string, loginId string) *TokenActive {
	//可在这里写分库逻辑
	redis := g.Redis()
	return &TokenActive{
		Ctx:   ctx,
		Redis: redis,
		Key:   fmt.Sprintf(consts.CacheTokenActiveFormat, sceneCode, loginId),
	}
}

func (cacheThis *TokenActive) Set(ttl int64) (err error) {
	err = cacheThis.Redis.SetEX(cacheThis.Ctx, cacheThis.Key, ttl, ttl)
	return
}

func (cacheThis *TokenActive) Get() (isExists int64, err error) {
	isExists, err = cacheThis.Redis.Exists(cacheThis.Ctx, cacheThis.Key)
	return
}

type TokenIsUnique struct {
	Ctx   context.Context
	Redis *gredis.Redis
	Key   string
}

func NewTokenIsUnique(ctx context.Context, sceneCode string, loginId string) *TokenIsUnique {
	//可在这里写分库逻辑
	redis := g.Redis()
	return &TokenIsUnique{
		Ctx:   ctx,
		Redis: redis,
		Key:   fmt.Sprintf(consts.CacheTokenIsUniqueFormat, sceneCode, loginId),
	}
}

func (cacheThis *TokenIsUnique) Set(value string, ttl int64) (err error) {
	err = cacheThis.Redis.SetEX(cacheThis.Ctx, cacheThis.Key, value, ttl)
	return
}

func (cacheThis *TokenIsUnique) Get() (value string, err error) {
	valueTmp, err := cacheThis.Redis.Get(cacheThis.Ctx, cacheThis.Key)
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
