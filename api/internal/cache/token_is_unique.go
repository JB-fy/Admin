package cache

import (
	"api/internal/consts"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

var TokenIsUnique = tokenIsUnique{redis: g.Redis()}

type tokenIsUnique struct{ redis *gredis.Redis }

func (cacheThis *tokenIsUnique) cache() *gredis.Redis {
	return cacheThis.redis
}

func (cacheThis *tokenIsUnique) key(sceneId string, loginId string) string {
	return fmt.Sprintf(consts.CacheTokenIsUniqueFormat, sceneId, loginId)
}

func (cacheThis *tokenIsUnique) Set(ctx context.Context, sceneId string, loginId string, value string, ttl int64) (err error) {
	err = cacheThis.cache().SetEX(ctx, cacheThis.key(sceneId, loginId), value, ttl)
	return
}

func (cacheThis *tokenIsUnique) Get(ctx context.Context, sceneId string, loginId string) (value string, err error) {
	valueTmp, err := cacheThis.cache().Get(ctx, cacheThis.key(sceneId, loginId))
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
