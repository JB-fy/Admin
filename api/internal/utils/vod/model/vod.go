package model

import (
	"context"
)

type VodFunc func(ctx context.Context, config map[string]any) Vod

type Vod interface {
	Sts(ctx context.Context, param VodParam) (stsInfo map[string]any, err error) // 获取Sts Token
}

type VodParam struct {
	ExpireTime int64 //签名有效时间。单位：秒
}
