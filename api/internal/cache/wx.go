package cache

import (
	"api/internal/consts"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

type WxGzhAccessToken struct {
	Ctx   context.Context
	Redis *gredis.Redis
	Key   string
}

// sceneCode 场景标识。注意：在同一权限场景下，存在互相覆盖BUG时，须自定义sceneCode规避
// loginName 账号/手机
func NewWxGzhAccessToken(ctx context.Context, appId string) *WxGzhAccessToken {
	//可在这里写分库逻辑
	redis := g.Redis()
	return &WxGzhAccessToken{
		Ctx:   ctx,
		Redis: redis,
		Key:   fmt.Sprintf(consts.CacheWxGzhAccessToken, appId),
	}
}

func (cacheThis *WxGzhAccessToken) Set(value string, ttl int64) (err error) {
	err = cacheThis.Redis.SetEX(cacheThis.Ctx, cacheThis.Key, value, ttl)
	return
}

func (cacheThis *WxGzhAccessToken) Get() (value string, err error) {
	valueTmp, err := cacheThis.Redis.Get(cacheThis.Ctx, cacheThis.Key)
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
