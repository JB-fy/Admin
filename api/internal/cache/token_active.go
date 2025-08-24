package cache

import (
	"api/internal/consts"
	"api/internal/utils/jbredis"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var TokenActive = tokenActive{}

type tokenActive struct{}

func (cacheThis *tokenActive) cache() redis.UniversalClient {
	return jbredis.DB()
}

func (cacheThis *tokenActive) key(sceneId string, loginId string) string {
	return fmt.Sprintf(consts.CACHE_TOKEN_ACTIVE, sceneId, loginId)
}

func (cacheThis *tokenActive) Set(ctx context.Context, sceneId string, loginId string, ttl time.Duration) error {
	return cacheThis.cache().SetEx(ctx, cacheThis.key(sceneId, loginId), ``, ttl).Err()
}

func (cacheThis *tokenActive) Reset(ctx context.Context, sceneId string, loginId string, ttl time.Duration) (bool, error) {
	return cacheThis.cache().SetXX(ctx, cacheThis.key(sceneId, loginId), ``, ttl).Result()
}
