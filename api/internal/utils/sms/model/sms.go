package model

import "context"

type SmsFunc func(ctx context.Context, config map[string]any) Sms

type Sms interface {
	SendCode(ctx context.Context, phone string, code string) (err error)
	SendSms(ctx context.Context, phoneArr []string, message string, paramOpt ...any) (err error)
}
