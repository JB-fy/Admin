package idCard

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
		idCardTypeVar, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(daoPlatform.Config.Columns().ConfigKey, `idCardType`).Value(daoPlatform.Config.Columns().ConfigValue)
		idCardType = idCardTypeVar.String()
	}

	switch idCardType {
	// case `aliyunIdCard`:
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunIdCardHost`, `aliyunIdCardPath`, `aliyunIdCardAppcode`})
		return NewAliyunIdCard(ctx, config)
	}
}
