package controller

import (
	"api/api"
	apiCurrent "api/api/app"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoUsers "api/internal/dao/users"
	"api/internal/utils"
	"api/internal/utils/email"
	"api/internal/utils/sms"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type Code struct{}

func NewCode() *Code {
	return &Code{}
}

// 发送验证码
func (controllerThis *Code) Send(ctx context.Context, req *apiCurrent.CodeSendReq) (res *api.CommonNoDataRes, err error) {
	switch req.Scene {
	case 0, 1, 2, 4:
		err = g.Validator().Rules(`phone`).Data(req.To).Run(ctx)
	case 10, 11, 12, 14:
		err = g.Validator().Rules(`email`).Data(req.To).Run(ctx)
	}
	if err != nil {
		return
	}

	to := req.To
	switch req.Scene {
	case 0, 2: //登录(手机)，密码找回(手机)
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Phone, to).One()
		if info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990000, ``)
			return
		}
		if info[daoUsers.Users.Columns().IsStop].Uint() == 1 {
			err = utils.NewErrorCode(ctx, 39990002, ``)
			return
		}
	case 1: //注册(手机)
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Phone, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991000, ``)
			return
		}
	case 3: //密码修改(手机)
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = loginInfo[daoUsers.Users.Columns().Phone].String()
		if to != `` {
			err = utils.NewErrorCode(ctx, 39991003, ``)
			return
		}
	case 4: //绑定(手机)
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		if loginInfo[daoUsers.Users.Columns().Phone].String() != `` {
			err = utils.NewErrorCode(ctx, 39991001, ``)
			return
		}
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Phone, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991002, ``)
			return
		}
	case 5: //解绑(手机)
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = loginInfo[daoUsers.Users.Columns().Phone].String()
		if to == `` {
			err = utils.NewErrorCode(ctx, 39991003, ``)
			return
		}
	case 10, 12: //登录(邮箱)，密码找回(邮箱)
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Email, to).One()
		if info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990000, ``)
			return
		}
		if info[daoUsers.Users.Columns().IsStop].Uint() == 1 {
			err = utils.NewErrorCode(ctx, 39990002, ``)
			return
		}
	case 11: //注册(邮箱)
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Email, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991010, ``)
			return
		}
	case 13: //密码修改(邮箱)
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = loginInfo[daoUsers.Users.Columns().Email].String()
		if to != `` {
			err = utils.NewErrorCode(ctx, 39991013, ``)
			return
		}
	case 14: //绑定(邮箱)
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		if loginInfo[daoUsers.Users.Columns().Email].String() != `` {
			err = utils.NewErrorCode(ctx, 39991011, ``)
			return
		}
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Email, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991012, ``)
			return
		}
	case 15: //解绑(邮箱)
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = loginInfo[daoUsers.Users.Columns().Email].String()
		if to == `` {
			err = utils.NewErrorCode(ctx, 39991013, ``)
			return
		}
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	code := grand.Digits(4)
	switch req.Scene {
	case 0, 1, 2, 3, 4, 5:
		err = sms.NewHandler(ctx).SendCode(to, code)
	case 10, 11, 12, 13, 14, 15:
		err = email.NewHandler(ctx).SendCode(to, code)
	}
	if err != nil {
		return
	}
	err = cache.NewCode(ctx, sceneId, to, req.Scene).Set(code, 5*60)
	return
}
