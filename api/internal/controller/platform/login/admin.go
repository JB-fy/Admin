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
	saltStatic, saltDynamic, err := service.LoginPlatformAdmin().Salt(ctx, req.LoginName)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{SaltStatic: saltStatic, SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *Admin) Login(ctx context.Context, req *apiLogin.AdminLoginReq) (res *api.CommonTokenRes, err error) {
	token, err := service.LoginPlatformAdmin().Login(ctx, req.LoginName, req.Password)
	if err != nil {
		return
	}
	res = &api.CommonTokenRes{Token: token}
	return
}
