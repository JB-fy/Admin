package sms

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

type Sms interface {
	SendCode(ctx context.Context, phone string, code string) (err error)
	SendSms(ctx context.Context, phoneArr []string, message string, paramOpt ...any) (err error)
}

var (
	smsTypeDef = `smsOfAliyun`
	smsFuncMap = map[string]func(ctx context.Context, config map[string]any) Sms{
		`smsOfAliyun`: func(ctx context.Context, config map[string]any) Sms { return NewSmsOfAliyun(ctx, config) },
	}
	smsMap = map[string]Sms{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	smsMu  sync.Mutex
)

func NewSms(ctx context.Context, smsType string, config map[string]any) (sms Sms) {
	smsKey := smsType + gmd5.MustEncrypt(config)
	ok := false
	if sms, ok = smsMap[smsKey]; ok { //先读一次（不加锁）
		return
	}
	smsMu.Lock()
	defer smsMu.Unlock()
	if sms, ok = smsMap[smsKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	if _, ok = smsFuncMap[smsType]; !ok {
		smsType = smsTypeDef
	}
	sms = smsFuncMap[smsType](ctx, config)
	smsMap[smsKey] = sms
	return
}
