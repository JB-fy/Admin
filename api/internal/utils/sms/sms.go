package sms

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type Sms interface {
	Send(phone string, code string) (err error)
	SendSms(phoneArr []string, templateParam string) (err error)
}

func NewSms(ctx context.Context) Sms {
	platformConfigColumns := daoPlatform.Config.Columns()
	smsType, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(platformConfigColumns.ConfigKey, `smsType`).Value(platformConfigColumns.ConfigValue)
	switch smsType.String() {
	// case `aliyunSms`:
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunSmsAccessKeyId`, `aliyunSmsAccessKeySecret`, `aliyunSmsEndpoint`, `aliyunSmsSignName`, `aliyunSmsTemplateCode`})
		return NewAliyunSms(ctx, config)
	}
}
