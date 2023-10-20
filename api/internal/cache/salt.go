package cache

import (
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

type Salt struct {
	Ctx   context.Context
	Redis *gredis.Redis
	Key   string
}

// loginName 账号/手机
// sceneCodeS 场景标识。注意：当一个权限场景（auth_scene表）存在不同用户表的短信验证码时，第二个用户表需要传sceneCodeS（自定义），否则两个用户表缓存时互相覆盖导致BUG
func NewSalt(ctx context.Context, loginName string, sceneCodeS ...string) *Salt {
	sceneCode := ``
	if len(sceneCodeS) > 0 && sceneCodeS[0] != `` {
		sceneCode = sceneCodeS[0]
	} else {
		sceneInfo := utils.GetCtxSceneInfo(ctx)
		sceneCode = sceneInfo[`sceneCode`].String()
	}
	//可以做分库逻辑
	redis := g.Redis()
	return &Salt{
		Ctx:   ctx,
		Redis: redis,
		Key:   fmt.Sprintf(`salt_%s_%s`, sceneCode, loginName),
	}
}

func (cacheThis *Salt) Set(value string, ttl int64) (err error) {
	err = cacheThis.Redis.SetEX(cacheThis.Ctx, cacheThis.Key, value, ttl)
	return
}

func (cacheThis *Salt) Get() (value string, err error) {
	valueTmp, err := cacheThis.Redis.Get(cacheThis.Ctx, cacheThis.Key)
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
