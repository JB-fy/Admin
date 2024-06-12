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

	"github.com/gogf/gf/v2/util/grand"
)

type Code struct{}

func NewCode() *Code {
	return &Code{}
}

// 发送验证码
func (controllerThis *Code) Send(ctx context.Context, req *apiCurrent.CodeSendReq) (res *api.CommonNoDataRes, err error) {
	toType := ``
	to := ``
	switch req.Scene {
	case 0, 2: //登录(手机)，密码找回(手机)
		toType = `phone`
		to = req.Phone

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
		toType = `phone`
		to = req.Phone

		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Phone, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991000, ``)
			return
		}
	case 3: //密码修改(手机)
		toType = `phone`

		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = loginInfo[daoUsers.Users.Columns().Phone].String()
		if to != `` {
			err = utils.NewErrorCode(ctx, 39990007, ``)
			return
		}
	case 4: //绑定(手机)
		toType = `phone`
		to = req.Phone

		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		if loginInfo[daoUsers.Users.Columns().Phone].String() != `` {
			err = utils.NewErrorCode(ctx, 39990005, ``)
			return
		}
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Phone, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990006, ``)
			return
		}
	case 5: //解绑(手机)
		toType = `phone`

		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = loginInfo[daoUsers.Users.Columns().Phone].String()
		if to == `` {
			err = utils.NewErrorCode(ctx, 39990007, ``)
			return
		}
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	code := grand.Digits(4)
	switch toType {
	case `phone`:
		err = sms.NewSms(ctx).SendCode(to, code)
	case `email`:
		err = email.NewEmail(ctx).SendCode(to, code)
	default:
		err = utils.NewErrorCode(ctx, 89999998, ``)
	}
	if err != nil {
		return
	}
	err = cache.NewCode(ctx, sceneCode, to, req.Scene).Set(code, 5*60)
	return
}
