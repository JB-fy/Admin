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

var deviceTypeMap = map[uint]string{0: `android`, 1: `ios`, 2: `mac_os`}

// 设备类型：0安卓 1苹果 2苹果电脑
func NewHandler(ctx context.Context, deviceType uint, pushTypeOpt ...string) model.Handler {
	handlerObj := &Handler{Ctx: ctx}
	pushType := ``
	if len(pushTypeOpt) > 0 {
		pushType = pushTypeOpt[0]
	} else {
		pushType = daoPlatform.Config.Get(ctx, `push_type`).String()
	}
	if _, ok := pushFuncMap[pushType]; !ok {
		pushType = pushTypeDef
	}
	config := daoPlatform.Config.Get(ctx, pushType).Map()
	switch pushType {
	case `push_of_tx`:
		deviceTypeStr, ok := deviceTypeMap[deviceType]
		if !ok {
			deviceTypeStr = `android`
		}
		config[`access_id`] = config[`access_id_of_`+deviceTypeStr]
		config[`secret_key`] = config[`secret_key_of_`+deviceTypeStr]
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
