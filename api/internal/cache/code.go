package cache

import (
	"api/internal/consts"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

type code struct {
	Ctx   context.Context
	Redis *gredis.Redis
	Key   string
}

// sceneId 场景ID。注意：在同一权限场景下，存在互相覆盖BUG时，须自定义sceneId规避
// to 手机/邮箱
// scene 场景
func NewCode(ctx context.Context, sceneId string, to string, scene uint) *code {
	//可在这里写分库逻辑
	redis := g.Redis()
	return &code{
		Ctx:   ctx,
		Redis: redis,
		Key:   fmt.Sprintf(consts.CacheCodeFormat, sceneId, to, scene),
	}
}

func (cacheThis *code) Set(value string, ttl int64) (err error) {
	err = cacheThis.Redis.SetEX(cacheThis.Ctx, cacheThis.Key, value, ttl)
	return
}

func (cacheThis *code) Get() (value string, err error) {
	valueTmp, err := cacheThis.Redis.Get(cacheThis.Ctx, cacheThis.Key)
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
