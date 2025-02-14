package push

import (
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils/push/model"
	"context"
)

type Handler struct {
	Ctx  context.Context
	push model.Push
}

// 设备类型：0安卓 1苹果 2苹果电脑
func NewHandler(ctx context.Context, deviceType uint, pushTypeOpt ...string) model.Handler {
	handlerObj := &Handler{Ctx: ctx}
	pushType := ``
	if len(pushTypeOpt) > 0 {
		pushType = pushTypeOpt[0]
	} else {
		pushType = daoPlatform.Config.GetOne(ctx, `pushType`).String()
	}
	if _, ok := pushFuncMap[pushType]; !ok {
		pushType = pushTypeDef
	}
	config := daoPlatform.Config.GetOne(ctx, pushType).Map()
	switch pushType {
	case `pushOfTx`:
		deviceTypeStr, ok := map[uint]string{0: `Android`, 1: `Ios`, 2: `MacOS`}[deviceType]
		if !ok {
			deviceTypeStr = `Android`
		}
		config[`accessID`] = config[`accessIDOf`+deviceTypeStr]
		config[`secretKey`] = config[`secretKeyOf`+deviceTypeStr]
	}
	handlerObj.push = NewPush(ctx, pushType, config)
	return handlerObj
}

func (handlerThis *Handler) Push(param model.PushParam) (err error) {
	return handlerThis.push.Push(handlerThis.Ctx, param)
}
func (handlerThis *Handler) Tag(param model.TagParam) (err error) {
	return handlerThis.push.Tag(handlerThis.Ctx, param)
}
