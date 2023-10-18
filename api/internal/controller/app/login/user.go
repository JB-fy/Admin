package controller

import (
	"api/api"
	apiLogin "api/api/app/login"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

// 获取加密盐
func (controllerThis *User) Salt(ctx context.Context, req *apiLogin.UserSaltReq) (res *api.CommonSaltRes, err error) {
	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) != nil && g.Validator().Rules(`passport`).Data(req.LoginName).Run(ctx) != nil {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}
	saltStatic, saltDynamic, err := service.LoginUser().Salt(ctx, req.LoginName)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{SaltStatic: saltStatic, SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *User) Login(ctx context.Context, req *apiLogin.UserLoginReq) (res *api.CommonTokenRes, err error) {
	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) != nil && g.Validator().Rules(`passport`).Data(req.LoginName).Run(ctx) != nil {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}
	token, err := service.LoginUser().Login(ctx, req.LoginName, req.Password, req.Code)
	if err != nil {
		return
	}
	res = &api.CommonTokenRes{Token: token}
	return
}
