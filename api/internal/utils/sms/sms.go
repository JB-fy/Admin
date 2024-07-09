package sms

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type Sms interface {
	SendCode(phone string, code string) (err error)
	SendSms(phoneArr []string, message string, paramOpt ...any) (err error)
}

func NewSms(ctx context.Context, smsTypeOpt ...string) Sms {
	smsType := ``
	if len(smsTypeOpt) > 0 {
		smsType = smsTypeOpt[0]
	} else {
		smsType, _ = daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `smsType`).ValueStr(daoPlatform.Config.Columns().ConfigValue)
	}

	switch smsType {
	// case `smsOfAliyun`:
	default:
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`smsOfAliyun`})
		config := configTmp[`smsOfAliyun`].Map()
		return NewSmsOfAliyun(ctx, config)
	}
}
