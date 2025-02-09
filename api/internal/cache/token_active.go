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
	return fmt.Sprintf(consts.CACHE_TOKEN_ACTIVE, sceneId, loginId)
}

func (cacheThis *tokenActive) Set(ctx context.Context, sceneId string, loginId string, ttl int64) (err error) {
	err = cacheThis.cache().SetEX(ctx, cacheThis.key(sceneId, loginId), ``, ttl)
	return
}

func (cacheThis *tokenActive) Reset(ctx context.Context, sceneId string, loginId string, ttl int64) (isSet bool, err error) {
	isSetVal, err := cacheThis.cache().Set(ctx, cacheThis.key(sceneId, loginId), ``, gredis.SetOption{TTLOption: gredis.TTLOption{EX: &ttl}, XX: true})
	if err != nil {
		return
	}
	isSet = isSetVal.Bool()
	return
}
