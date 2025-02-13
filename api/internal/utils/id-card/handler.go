package id_card

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type Handler struct {
	Ctx    context.Context
	IdCard IdCard
}

func NewHandler(ctx context.Context, idCardTypeOpt ...string) *Handler {
	handlerObj := &Handler{Ctx: ctx}
	idCardType := ``
	if len(idCardTypeOpt) > 0 {
		idCardType = idCardTypeOpt[0]
	} else {
		idCardType = daoPlatform.Config.GetOne(ctx, `idCardType`).String()
	}
	if _, ok := idCardFuncMap[idCardType]; !ok {
		idCardType = idCardTypeDef
	}
	config := daoPlatform.Config.GetOne(ctx, idCardType).Map()
	handlerObj.IdCard = NewIdCard(ctx, idCardType, config)
	return handlerObj
}

func (handlerThis *Handler) Auth(idCardName string, idCardNo string) (idCardInfo IdCardInfo, err error) {
	return handlerThis.IdCard.Auth(handlerThis.Ctx, idCardName, idCardNo)
}
