package cache

import (
	"api/internal/consts"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

var TokenActive = tokenActive{redis: g.Redis()}

type tokenActive struct{ redis *gredis.Redis }

func (cacheThis *tokenActive) cache() *gredis.Redis {
	return cacheThis.redis
}

func (cacheThis *tokenActive) key(sceneId string, loginId string) string {
	return fmt.Sprintf(consts.CacheTokenActiveFormat, sceneId, loginId)
}

func (cacheThis *tokenActive) Set(ctx context.Context, sceneId string, loginId string, ttl int64) (err error) {
	err = cacheThis.cache().SetEX(ctx, cacheThis.key(sceneId, loginId), ttl, ttl)
	return
}

func (cacheThis *tokenActive) Get(ctx context.Context, sceneId string, loginId string) (isExists int64, err error) {
	isExists, err = cacheThis.cache().Exists(ctx, cacheThis.key(sceneId, loginId))
	return
}
