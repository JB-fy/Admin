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
		idCardType, _ = daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `idCardType`).ValueStr(daoPlatform.Config.Columns().ConfigValue)
	}

	var config g.Map
	switch idCardType {
	// case `idCardOfAliyun`:
	default:
		config, _ = daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `idCardOfAliyun`).ValueMap(daoPlatform.Config.Columns().ConfigValue)
	}

	config[`idCardType`] = idCardType
	handlerObj.IdCard = NewIdCard(config)
	return handlerObj
}

func (handlerThis *Handler) Auth(idCardName string, idCardNo string) (idCardInfo IdCardInfo, err error) {
	return handlerThis.IdCard.Auth(handlerThis.Ctx, idCardName, idCardNo)
}
