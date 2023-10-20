package cache

import (
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

type Token struct {
	Ctx   context.Context
	Redis *gredis.Redis
	Key   string
}

// loginId 登录用户ID
// sceneCodeS 场景标识。注意：当一个权限场景（auth_scene表）存在不同用户表的短信验证码时，第二个用户表需要传sceneCodeS（自定义），否则两个用户表缓存时互相覆盖导致BUG
func NewToken(ctx context.Context, loginId uint, sceneCodeS ...string) *Token {
	sceneCode := ``
	if len(sceneCodeS) > 0 && sceneCodeS[0] != `` {
		sceneCode = sceneCodeS[0]
	} else {
		sceneInfo := utils.GetCtxSceneInfo(ctx)
		sceneCode = sceneInfo[`sceneCode`].String()
	}
	//可以做分库逻辑
	redis := g.Redis()
	return &Token{
		Ctx:   ctx,
		Redis: redis,
		Key:   fmt.Sprintf(`token_%s_%d`, sceneCode, loginId),
	}
}

func (cacheThis *Token) Set(value string, ttl int64) (err error) {
	err = cacheThis.Redis.SetEX(cacheThis.Ctx, cacheThis.Key, value, ttl)
	return
}

func (cacheThis *Token) Get() (value string, err error) {
	valueTmp, err := cacheThis.Redis.Get(cacheThis.Ctx, cacheThis.Key)
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
