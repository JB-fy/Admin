package push

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type Handler struct {
	Ctx  context.Context
	Push Push
}

// 设备类型：0安卓 1苹果 2苹果电脑
func NewHandler(ctx context.Context, deviceType uint, pushTypeOpt ...string) *Handler {
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
	handlerObj.Push = NewPush(ctx, pushType, config)
	return handlerObj
}

type CodeTemplate struct {
	Subject  string `json:"subject"`
	Template string `json:"template"`
}

func (handlerThis *Handler) PushMsg(param PushParam) (err error) {
	return handlerThis.Push.PushMsg(handlerThis.Ctx, param)
}
func (handlerThis *Handler) TagHandle(param TagParam) (err error) {
	return handlerThis.Push.TagHandle(handlerThis.Ctx, param)
}
