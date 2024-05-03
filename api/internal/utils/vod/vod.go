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
		vodType, _ = daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `vodType`).ValueStr(daoPlatform.Config.Columns().ConfigValue)
	}

	switch vodType {
	// case `vodOfAliyun`:
	default:
		return NewVodOfAliyun(ctx)
	}
}
