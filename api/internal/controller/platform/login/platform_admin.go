package controller

import (
	"api/api"
	apiLogin "api/api/platform/login"
	"api/internal/service"
	"context"
)

type PlatformAdmin struct{}

func NewPlatformAdmin() *PlatformAdmin {
	return &PlatformAdmin{}
}

// 获取加密盐
func (controllerThis *PlatformAdmin) Salt(ctx context.Context, req *apiLogin.PlatformAdminSaltReq) (res *api.CommonSaltRes, err error) {
	saltStatic, saltDynamic, err := service.PlatformAdmin().Salt(ctx, req.Account)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{SaltStatic: saltStatic, SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *PlatformAdmin) Login(ctx context.Context, req *apiLogin.PlatformAdminLoginReq) (res *api.CommonTokenRes, err error) {
	token, err := service.PlatformAdmin().Login(ctx, req.Account, req.Password)
	if err != nil {
		return
	}
	res = &api.CommonTokenRes{Token: token}
	return
}
