package sms

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type Sms interface {
	Send(phone string, code string) (err error)
	SendSms(phoneArr []string, templateParam string) (err error)
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
		return NewSmsOfAliyun(ctx)
	}
}
