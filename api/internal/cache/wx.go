package cache

import (
	"api/internal/consts"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

type wxGzhAccessToken struct {
	Ctx   context.Context
	Redis *gredis.Redis
	Key   string
}

// appId 微信公众号AppId
func NewWxGzhAccessToken(ctx context.Context, appId string) *wxGzhAccessToken {
	//可在这里写分库逻辑
	redis := g.Redis()
	return &wxGzhAccessToken{
		Ctx:   ctx,
		Redis: redis,
		Key:   fmt.Sprintf(consts.CacheWxGzhAccessToken, appId),
	}
}

func (cacheThis *wxGzhAccessToken) Set(value string, ttl int64) (err error) {
	err = cacheThis.Redis.SetEX(cacheThis.Ctx, cacheThis.Key, value, ttl)
	return
}

func (cacheThis *wxGzhAccessToken) Get() (value string, err error) {
	valueTmp, err := cacheThis.Redis.Get(cacheThis.Ctx, cacheThis.Key)
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
