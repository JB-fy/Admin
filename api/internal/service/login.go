// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ILoginPlatformAdmin interface {
		// 获取加密盐
		Salt(ctx context.Context, account string) (saltStatic string, saltDynamic string, err error)
		// 登录
		Login(ctx context.Context, account string, password string) (token string, err error)
	}
)

var (
	localLoginPlatformAdmin ILoginPlatformAdmin
)

func LoginPlatformAdmin() ILoginPlatformAdmin {
	if localLoginPlatformAdmin == nil {
		panic("implement not found for interface ILoginPlatformAdmin, forgot register?")
	}
	return localLoginPlatformAdmin
}

func RegisterLoginPlatformAdmin(i ILoginPlatformAdmin) {
	localLoginPlatformAdmin = i
}
