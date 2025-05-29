package sms

import (
	"api/internal/utils/sms/aliyun"
	"api/internal/utils/sms/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

var (
	smsMap     = map[string]model.Sms{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	smsMu      sync.Mutex
	smsTypeDef = `sms_of_aliyun`
	smsFuncMap = map[string]model.SmsFunc{
		`sms_of_aliyun`: aliyun.NewSms,
	}
)

func NewSms(ctx context.Context, smsType string, config map[string]any) (sms model.Sms) {
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
