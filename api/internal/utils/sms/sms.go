package sms

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type Sms interface {
	Send(ctx context.Context, phone string, code string) (err error)
	SendSms(ctx context.Context, phoneArr []string, templateParam string) (err error)
}

func NewSms(ctx context.Context) Sms {
	platformConfigColumns := daoPlatform.Config.Columns()
	smsType, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(platformConfigColumns.ConfigKey, `smsType`).Value(platformConfigColumns.ConfigValue)
	switch smsType.String() {
	case `aliyunSms`:
		config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunSmsAccessKeyId`, `aliyunSmsAccessKeySecret`, `aliyunSmsEndpoint`, `aliyunSmsSignName`, `aliyunSmsTemplateCode`})
		return NewAliyunSms(ctx, config)
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunSmsAccessKeyId`, `aliyunSmsAccessKeySecret`, `aliyunSmsEndpoint`, `aliyunSmsSignName`, `aliyunSmsTemplateCode`})
		return NewAliyunSms(ctx, config)
	}
}
