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

func NewVod(ctx context.Context) Vod {
	platformConfigColumns := daoPlatform.Config.Columns()
	vodType, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(platformConfigColumns.ConfigKey, `vodType`).Value(platformConfigColumns.ConfigValue)
	switch vodType.String() {
	// case `aliyunVod`:
	default:
		// config, _ := daoPlatform.Config.Get(ctx, []string{`localVodUrl`, `localVodSignKey`, `localVodFileSaveDir`, `localVodFileUrlPrefix`})
		config := map[string]interface{}{
			`aliyunVodAccessKeyId`:     `LTAI5tFVtUyYRbetup1hxjjX`,
			`aliyunVodAccessKeySecret`: `pcSsoLLkZPfnTWRAtx0inYhJcpwNRC`,
			`aliyunVodRoleArn`:         `acs:ram::1359390739767110:role/aliyunvodtokengeneratorrole`,
			`aliyunVodEndpoint`:        `sts.cn-shanghai.aliyuncs.com`,
		}
		return NewAliyunVod(ctx, config)
	}
}
