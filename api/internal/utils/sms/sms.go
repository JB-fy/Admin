package sms

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

type Sms interface {
	SendCode(ctx context.Context, phone string, code string) (err error)
	SendSms(ctx context.Context, phoneArr []string, message string, paramOpt ...any) (err error)
}

var (
	smsMap = map[string]Sms{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	smsMu  sync.Mutex
)

func NewSms(config map[string]any) (sms Sms) {
	smsKey := gmd5.MustEncrypt(config)

	ok := false
	if sms, ok = smsMap[smsKey]; ok { //先读一次（不加锁）
		return
	}
	smsMu.Lock()
	defer smsMu.Unlock()
	if sms, ok = smsMap[smsKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}

	switch gconv.String(config[`smsType`]) {
	// case `smsOfAliyun`:
	default:
		sms = NewSmsOfAliyun(config)
	}
	smsMap[smsKey] = sms
	return
}
