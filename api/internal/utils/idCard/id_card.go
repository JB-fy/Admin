package idCard

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type IdCard interface {
	Auth(idCardName string, idCardNo string) (idCardInfo IdCardInfo, err error)
}

type IdCardInfo struct {
	Gender uint // 性别：0未设置 1男 2女
	// Birthday *gtime.Time // 生日
	Birthday string // 生日
	Address  string // 详细地址
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
