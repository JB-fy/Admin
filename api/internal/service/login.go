// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ILogin interface {
		// 获取登录加密字符串(前端登录操作用于加密密码后提交)
		Salt(ctx context.Context, account string) (salt string, err error)
		// 登录(平台后台管理员)
		PlatformAdmin(ctx context.Context, account string, password string) (token string, err error)
	}
)

var (
	localLogin ILogin
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}
