package cache

import (
	"api/internal/consts"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

// appId 微信公众号AppId
var WxGzhAccessToken = wxGzhAccessToken{redis: g.Redis()}

type wxGzhAccessToken struct{ redis *gredis.Redis }

func (cacheThis *wxGzhAccessToken) cache() *gredis.Redis {
	return cacheThis.redis
}

func (cacheThis *wxGzhAccessToken) key(appId string) string {
	return fmt.Sprintf(consts.CACHE_WX_GZH_ACCESS_TOKEN, appId)
}

func (cacheThis *wxGzhAccessToken) Set(ctx context.Context, appId string, value string, ttl int64) (err error) {
	err = cacheThis.cache().SetEX(ctx, cacheThis.key(appId), value, ttl)
	return
}

func (cacheThis *wxGzhAccessToken) Get(ctx context.Context, appId string) (value string, err error) {
	valueTmp, err := cacheThis.cache().Get(ctx, cacheThis.key(appId))
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
