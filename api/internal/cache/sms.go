package cache

import (
	"api/internal/consts"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

type Sms struct {
	Ctx   context.Context
	Redis *gredis.Redis
	Key   string
}

// phone 手机
// useScene 使用场景
// sceneCode 场景标识。注意：在同一权限场景下，存在互相覆盖BUG时，须自定义sceneCode规避
func NewSms(ctx context.Context, phone string, useScene int, sceneCode ...string) *Sms {
	//可以做分库逻辑
	redis := g.Redis()
	return &Sms{
		Ctx:   ctx,
		Redis: redis,
		Key:   fmt.Sprintf(consts.CacheSmsFormat, utils.GetSceneCode(ctx, sceneCode...), phone, useScene),
	}
}

func (cacheThis *Sms) Set(value string, ttl int64) (err error) {
	err = cacheThis.Redis.SetEX(cacheThis.Ctx, cacheThis.Key, value, ttl)
	return
}

func (cacheThis *Sms) Get() (value string, err error) {
	valueTmp, err := cacheThis.Redis.Get(cacheThis.Ctx, cacheThis.Key)
	if err != nil {
		return
	}
	value = valueTmp.String()
	return
}
