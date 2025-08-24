package cache

import (
	"api/internal/consts"
	"api/internal/utils/jbredis"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var TokenIsUnique = tokenIsUnique{}

type tokenIsUnique struct{}

func (cacheThis *tokenIsUnique) cache() redis.UniversalClient {
	return jbredis.DB()
}

func (cacheThis *tokenIsUnique) key(sceneId string, loginId string) string {
	return fmt.Sprintf(consts.CACHE_TOKEN_IS_UNIQUE, sceneId, loginId)
}

func (cacheThis *tokenIsUnique) Set(ctx context.Context, sceneId string, loginId string, value string, ttl time.Duration) error {
	return cacheThis.cache().SetEx(ctx, cacheThis.key(sceneId, loginId), value, ttl).Err()
}

func (cacheThis *tokenIsUnique) Get(ctx context.Context, sceneId string, loginId string) (string, error) {
	return cacheThis.cache().Get(ctx, cacheThis.key(sceneId, loginId)).Result()
}
