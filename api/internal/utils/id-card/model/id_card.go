package model

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

type IdCardFunc func(ctx context.Context, config map[string]any) IdCard

type IdCard interface {
	Auth(ctx context.Context, idCardName string, idCardNo string) (idCardInfo IdCardInfo, err error)
}

type IdCardInfo struct {
	Birthday *gtime.Time // 生日
	Address  string      // 详细地址
	Gender   uint8       // 性别：0未设置 1男 2女
}
