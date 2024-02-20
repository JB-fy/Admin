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

func NewVod(ctx context.Context, vodTypeOpt ...string) Vod {
	vodType := ``
	if len(vodTypeOpt) > 0 {
		vodType = vodTypeOpt[0]
	} else {
		vodTypeVar, _ := daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `vodType`).Value(daoPlatform.Config.Columns().ConfigValue)
		vodType = vodTypeVar.String()
	}

	switch vodType {
	// case `vodOfAliyun`:
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`vodOfAliyunAccessKeyId`, `vodOfAliyunAccessKeySecret`, `vodOfAliyunEndpoint`, `vodOfAliyunRoleArn`})
		return NewVodOfAliyun(ctx, config.Map())
	}
}
