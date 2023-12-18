package vod

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type VodParam struct {
	ExpireTime int64 //签名有效时间。单位：秒
}

type Vod interface {
	Sts(param VodParam) (stsInfo map[string]interface{}, err error) // 获取Sts Token
}

func CreateVodParam() (param VodParam) {
	param = VodParam{
		ExpireTime: 50 * 60,
	}
	return
}

func NewVod(ctx context.Context, vodTypeTmp ...string) Vod {
	vodType := ``
	if len(vodTypeTmp) > 0 {
		vodType = vodTypeTmp[0]
	} else {
		vodTypeVar, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(daoPlatform.Config.Columns().ConfigKey, `vodType`).Value(daoPlatform.Config.Columns().ConfigValue)
		vodType = vodTypeVar.String()
	}

	switch vodType {
	// case `aliyunVod`:
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunVodAccessKeyId`, `aliyunVodAccessKeySecret`, `aliyunVodEndpoint`, `aliyunVodRoleArn`})
		return NewAliyunVod(ctx, config)
	}
}
