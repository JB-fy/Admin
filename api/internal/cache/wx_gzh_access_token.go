package cache

import (
	"api/internal/consts"
	"api/internal/utils/jbredis"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var WxGzhAccessToken = wxGzhAccessToken{}

type wxGzhAccessToken struct{}

func (cacheThis *wxGzhAccessToken) cache() redis.UniversalClient {
	return jbredis.DB()
}

func (cacheThis *wxGzhAccessToken) key(appId string) string {
	return fmt.Sprintf(consts.CACHE_WX_GZH_ACCESS_TOKEN, appId)
}

func (cacheThis *wxGzhAccessToken) Set(ctx context.Context, appId string, value string, ttl time.Duration) error {
	return cacheThis.cache().SetEx(ctx, cacheThis.key(appId), value, ttl).Err()
}

func (cacheThis *wxGzhAccessToken) Get(ctx context.Context, appId string) (string, error) {
	return cacheThis.cache().Get(ctx, cacheThis.key(appId)).Result()
}
