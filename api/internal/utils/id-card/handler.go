package id_card

import (
	daoPlatform "api/internal/dao/platform"
	"context"

	"github.com/gogf/gf/v2/frame/g"
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

	var config g.Map
	switch idCardType {
	// case `idCardOfAliyun`:
	default:
		config = daoPlatform.Config.GetOne(ctx, `idCardOfAliyun`).Map()
	}

	config[`idCardType`] = idCardType
	handlerObj.IdCard = NewIdCard(config)
	return handlerObj
}

func (handlerThis *Handler) Auth(idCardName string, idCardNo string) (idCardInfo IdCardInfo, err error) {
	return handlerThis.IdCard.Auth(handlerThis.Ctx, idCardName, idCardNo)
}
