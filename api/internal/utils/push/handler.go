package push

import (
	daoPlatform "api/internal/dao/platform"
	"context"

	"github.com/gogf/gf/v2/frame/g"
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

	var config g.Map
	switch pushType {
	// case `pushOfTx`:	//腾讯移动推送
	default:
		config = daoPlatform.Config.GetOne(ctx, `pushOfTx`).Map()
		switch deviceType {
		case 1: //IOS
			config[`accessID`] = config[`accessIDOfIos`]
			config[`secretKey`] = config[`secretKeyOfIos`]
		case 2: //MacOS
			config[`accessID`] = config[`accessIDOfMacOS`]
			config[`secretKey`] = config[`secretKeyOfMacOS`]
		// case 0: //安卓
		default:
			config[`accessID`] = config[`accessIDOfAndroid`]
			config[`secretKey`] = config[`secretKeyOfAndroid`]
		}
	}

	config[`pushType`] = pushType
	handlerObj.Push = NewPush(config)
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
