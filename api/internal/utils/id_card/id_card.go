package id_card

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type IdCard interface {
	Auth(idCardName string, idCardNo string) (idCardInfo map[string]interface{}, err error)
}

func NewIdCard(ctx context.Context) IdCard {
	platformConfigColumns := daoPlatform.Config.Columns()
	idCardType, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(platformConfigColumns.ConfigKey, `idCardType`).Value(platformConfigColumns.ConfigValue)
	switch idCardType.String() {
	case `aliyunIdCard`:
		config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunIdCardHost`, `aliyunIdCardPath`, `aliyunIdCardAppcode`})
		return NewAliyunIdCard(ctx, config)
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunIdCardHost`, `aliyunIdCardPath`, `aliyunIdCardAppcode`})
		return NewAliyunIdCard(ctx, config)
	}
}
