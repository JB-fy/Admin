// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IPlatformAdmin interface {
		// 获取加密盐
		Salt(ctx context.Context, account string) (saltStatic string, saltDynamic string, err error)
		// 登录
		Login(ctx context.Context, account string, password string) (token string, err error)
	}
)

var (
	localPlatformAdmin IPlatformAdmin
)

func PlatformAdmin() IPlatformAdmin {
	if localPlatformAdmin == nil {
		panic("implement not found for interface IPlatformAdmin, forgot register?")
	}
	return localPlatformAdmin
}

func RegisterPlatformAdmin(i IPlatformAdmin) {
	localPlatformAdmin = i
}
