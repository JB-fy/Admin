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
func (controllerThis *Admin) Salt(ctx context.Context, req *apiLogin.AdminSaltReq) (res *api.CommonSaltRes, err error) {
	salt, err := service.Login().Salt(ctx, req.Account)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{Salt: salt}
	return
}

// 登录
func (controllerThis *Admin) Login(ctx context.Context, req *apiLogin.AdminLoginReq) (res *api.CommonTokenRes, err error) {
	token, err := service.Login().PlatformAdmin(ctx, req.Account, req.Password)
	if err != nil {
		return
	}
	res = &api.CommonTokenRes{Token: token}
	return
}
