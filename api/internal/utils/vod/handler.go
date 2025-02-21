package vod

import (
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils/vod/model"
	"context"
)

type Handler struct {
	Ctx   context.Context
	Scene string //上传场景。default默认。根据自身需求扩展，用于确定上传通道和上传参数
	vod   model.Vod
}

func NewHandler(ctx context.Context, scene string, vodTypeOpt ...string) model.Handler {
	handlerObj := &Handler{
		Ctx:   ctx,
		Scene: scene,
	}
	vodType := ``
	if len(vodTypeOpt) > 0 {
		vodType = vodTypeOpt[0]
	} else {
		vodType = daoPlatform.Config.GetOne(ctx, `vodType`).String()
	}
	if _, ok := vodFuncMap[vodType]; !ok {
		vodType = vodTypeDef
	}
	config := daoPlatform.Config.GetOne(ctx, vodType).Map()
	handlerObj.vod = NewVod(ctx, vodType, config)
	return handlerObj
}

func (handlerThis *Handler) Sts() (stsInfo map[string]any, err error) {
	return handlerThis.vod.Sts(handlerThis.Ctx, handlerThis.createVodParam())
}

func (handlerThis *Handler) createVodParam() (param model.VodParam) {
	switch handlerThis.Scene {
	default:
		param = model.VodParam{
			ExpireTime: 60 * 60,
		}
	}
	return
}
