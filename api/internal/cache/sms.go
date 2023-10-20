package cache

import (
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

// sceneCode 场景标识
// phone 手机
// useScene 使用场景
func NewSms(ctx context.Context, sceneCode string, phone string, useScene int) *Sms {
	//可以做分库逻辑
	redis := g.Redis()
	return &Sms{
		Ctx:   ctx,
		Redis: redis,
		Key:   fmt.Sprintf(`sms_%s_%s_%d`, sceneCode, phone, useScene),
	}
}

func (cacheThis *Sms) SetSmsCode(smsCode string, ttl int64) (err error) {
	err = cacheThis.Redis.SetEX(cacheThis.Ctx, cacheThis.Key, smsCode, ttl)
	return
}

func (cacheThis *Sms) GetSmsCode() (smsCode string, err error) {
	smsCodeVar, err := cacheThis.Redis.Get(cacheThis.Ctx, cacheThis.Key)
	if err != nil {
		return
	}
	smsCode = smsCodeVar.String()
	return
}
