package sms

import (
	"api/internal/consts"
	daoConfig "api/internal/dao/config"
	"api/internal/utils/sms/model"
	"context"
)

type Handler struct {
	Ctx context.Context
	sms model.Sms
}

func NewHandler(ctx context.Context, smsTypeOpt ...string) model.Handler {
	handlerObj := &Handler{Ctx: ctx}
	smsType := ``
	if len(smsTypeOpt) > 0 {
		smsType = smsTypeOpt[0]
	} else {
		smsType = daoConfig.Config.Get(ctx, consts.SCENE_ID_PLATFORM, 0, `sms_type`).String()
	}
	if _, ok := smsFuncMap[smsType]; !ok {
		smsType = smsTypeDef
	}
	config := daoConfig.Config.Get(ctx, consts.SCENE_ID_PLATFORM, 0, smsType).Map()
	handlerObj.sms = NewSms(ctx, smsType, config)
	return handlerObj
}

func (handlerThis *Handler) SendCode(phone string, code string) (err error) {
	return handlerThis.sms.SendCode(handlerThis.Ctx, phone, code)
}

func (handlerThis *Handler) SendSms(phoneArr []string, message string, paramOpt ...any) (err error) {
	return handlerThis.sms.SendSms(handlerThis.Ctx, phoneArr, message, paramOpt...)
}
