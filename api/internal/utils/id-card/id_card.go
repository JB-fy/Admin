package id_card

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type IdCardInfo struct {
	Gender uint // 性别：0未设置 1男 2女
	// Birthday *gtime.Time // 生日
	Birthday string // 生日
	Address  string // 详细地址
}

type IdCard interface {
	Auth(idCardName string, idCardNo string) (idCardInfo IdCardInfo, err error)
}

func NewIdCard(ctx context.Context, idCardTypeOpt ...string) IdCard {
	idCardType := ``
	if len(idCardTypeOpt) > 0 {
		idCardType = idCardTypeOpt[0]
	} else {
		idCardType, _ = daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `idCardType`).ValueStr(daoPlatform.Config.Columns().ConfigValue)
	}

	switch idCardType {
	// case `idCardOfAliyun`:
	default:
		return NewIdCardOfAliyun(ctx)
	}
}
