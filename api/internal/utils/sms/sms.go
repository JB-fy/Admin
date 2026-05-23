package sms

import (
	"api/internal/utils/sms/aliyun"
	"api/internal/utils/sms/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"golang.org/x/sync/singleflight"
)

var (
	smsMap     sync.Map
	smsSfg     singleflight.Group
	smsFuncMap = map[string]model.SmsFunc{
		`sms_of_aliyun`: aliyun.NewSms,
	}
	smsTypeDef = `sms_of_aliyun`
)

func NewSms(ctx context.Context, smsType string, config map[string]any) (obj model.Sms) {
	if _, ok := smsFuncMap[smsType]; !ok {
		smsType = smsTypeDef
	}
	key := smsType + gmd5.MustEncrypt(config)
	objTmp, ok := smsMap.Load(key)
	if !ok {
		objTmp, _, _ = smsSfg.Do(key, func() (obj any, err error) {
			obj = smsFuncMap[smsType](ctx, config)
			smsMap.Store(key, obj)
			return
		})
	}
	obj = objTmp.(model.Sms)
	return
}
