package controller

import (
	"api/api"
	apiLogin "api/api/platform/login"
	"api/internal/service"
	"context"
)

type Admin struct{}

func NewAdmin() *Admin {
	return &Admin{}
}

// 获取加密盐
func (controllerThis *Admin) EncryptStr(ctx context.Context, req *apiLogin.AdminEncryptStrReq) (res *api.CommonEncryptStrRes, err error) {
	encryptStr, err := service.Login().EncryptStr(ctx, `platform`, req.Account)
	if err != nil {
		return
	}
	res = &api.CommonEncryptStrRes{EncryptStr: encryptStr}
	return
}

// 登录
func (controllerThis *Admin) Login(ctx context.Context, req *apiLogin.AdminLoginReq) (res *api.CommonTokenRes, err error) {
	token, err := service.Login().Platform(ctx, req.Account, req.Password)
	if err != nil {
		return
	}
	res = &api.CommonTokenRes{Token: token}
	return
}
