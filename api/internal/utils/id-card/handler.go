package id_card

import (
	"api/internal/consts"
	daoConfig "api/internal/dao/config"
	"api/internal/utils/id-card/model"
	"context"
)

type Handler struct {
	Ctx    context.Context
	idCard model.IdCard
}

func NewHandler(ctx context.Context, idCardTypeOpt ...string) model.Handler {
	handlerObj := &Handler{Ctx: ctx}
	idCardType := ``
	if len(idCardTypeOpt) > 0 {
		idCardType = idCardTypeOpt[0]
	} else {
		idCardType = daoConfig.Config.Get(ctx, consts.SCENE_ID_PLATFORM, 0, `id_card_type`).String()
	}
	if _, ok := idCardFuncMap[idCardType]; !ok {
		idCardType = idCardTypeDef
	}
	config := daoConfig.Config.Get(ctx, consts.SCENE_ID_PLATFORM, 0, idCardType).Map()
	handlerObj.idCard = NewIdCard(ctx, idCardType, config)
	return handlerObj
}

func (handlerThis *Handler) Auth(idCardName string, idCardNo string) (idCardInfo model.IdCardInfo, err error) {
	return handlerThis.idCard.Auth(handlerThis.Ctx, idCardName, idCardNo)
}
