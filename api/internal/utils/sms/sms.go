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
		smsTypeVar, _ := daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `smsType`).Value(daoPlatform.Config.Columns().ConfigValue)
		smsType = smsTypeVar.String()
	}

	switch smsType {
	// case `smsOfAliyun`:
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`smsOfAliyunAccessKeyId`, `smsOfAliyunAccessKeySecret`, `smsOfAliyunEndpoint`, `smsOfAliyunSignName`, `smsOfAliyunTemplateCode`})
		return NewSmsOfAliyun(ctx, config.Map())
	}
}
