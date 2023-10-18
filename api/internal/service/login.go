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
		Salt(ctx context.Context, loginName string) (saltStatic string, saltDynamic string, err error)
		// 登录
		Login(ctx context.Context, loginName string, password string) (token string, err error)
	}
	ILoginUser interface {
		// 获取加密盐
		Salt(ctx context.Context, loginName string) (saltStatic string, saltDynamic string, err error)
		// 登录
		Login(ctx context.Context, loginName string, password string, code string) (token string, err error)
	}
)

var (
	localLoginUser          ILoginUser
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

func LoginUser() ILoginUser {
	if localLoginUser == nil {
		panic("implement not found for interface ILoginUser, forgot register?")
	}
	return localLoginUser
}

func RegisterLoginUser(i ILoginUser) {
	localLoginUser = i
}
