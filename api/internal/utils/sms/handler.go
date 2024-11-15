package sms

import (
	daoPlatform "api/internal/dao/platform"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type Handler struct {
	Ctx context.Context
	Sms Sms
}

func NewHandler(ctx context.Context, smsTypeOpt ...string) *Handler {
	handlerObj := &Handler{Ctx: ctx}

	smsType := ``
	if len(smsTypeOpt) > 0 {
		smsType = smsTypeOpt[0]
	} else {
		smsType = daoPlatform.Config.GetOne(ctx, `smsType`).String()
	}

	var config g.Map
	switch smsType {
	// case `smsOfAliyun`:
	default:
		config = daoPlatform.Config.GetOne(ctx, `smsOfAliyun`).Map()
	}

	config[`smsType`] = smsType
	handlerObj.Sms = NewSms(config)
	return handlerObj
}

func (handlerThis *Handler) SendCode(phone string, code string) (err error) {
	return handlerThis.Sms.SendCode(handlerThis.Ctx, phone, code)
}

func (handlerThis *Handler) SendSms(phoneArr []string, message string, paramOpt ...any) (err error) {
	return handlerThis.Sms.SendSms(handlerThis.Ctx, phoneArr, message, paramOpt...)
}
