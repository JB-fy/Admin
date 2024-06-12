package controller

import (
	"api/api"
	apiCurrent "api/api/platform"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
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
	case 4: //绑定手机
		toType = `phone`
		to = req.Phone

		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filter(daoPlatform.Admin.Columns().Phone, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990006, ``)
			return
		}
	case 14: //绑定邮箱
		toType = `email`
		to = req.Email

		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filter(daoPlatform.Admin.Columns().Email, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990006, ``)
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
