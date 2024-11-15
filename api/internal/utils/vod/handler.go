package vod

import (
	daoPlatform "api/internal/dao/platform"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type Handler struct {
	Ctx   context.Context
	Scene string //上传场景。default默认。根据自身需求扩展，用于确定上传通道和上传参数
	Vod   Vod
}

func NewHandler(ctx context.Context, scene string, vodTypeOpt ...string) *Handler {
	handlerObj := &Handler{
		Ctx:   ctx,
		Scene: scene,
	}

	vodType := ``
	if len(vodTypeOpt) > 0 {
		vodType = vodTypeOpt[0]
	} else {
		vodType, _ = daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `vodType`).ValueStr(daoPlatform.Config.Columns().ConfigValue)
	}

	var config g.Map
	switch vodType {
	// case `vodOfAliyun`:
	default:
		config, _ = daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `vodOfAliyun`).ValueMap(daoPlatform.Config.Columns().ConfigValue)
	}

	config[`vodType`] = vodType
	handlerObj.Vod = NewVod(config)
	return handlerObj
}

func (handlerThis *Handler) Sts() (stsInfo map[string]any, err error) {
	return handlerThis.Vod.Sts(handlerThis.Ctx, handlerThis.createVodParam())
}

func (handlerThis *Handler) createVodParam() (param VodParam) {
	switch handlerThis.Scene {
	default:
		param = VodParam{
			ExpireTime: 50 * 60,
		}
	}
	return
}
