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
	case 4:
		err = g.Validator().Rules(`phone`).Data(req.To).Run(ctx)
	case 14:
		err = g.Validator().Rules(`email`).Data(req.To).Run(ctx)
	}
	if err != nil {
		return
	}

	to := req.To
	switch req.Scene {
	case 4: //绑定手机
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filter(daoPlatform.Admin.Columns().Phone, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991002, ``)
			return
		}
	case 14: //绑定邮箱
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filter(daoPlatform.Admin.Columns().Email, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991012, ``)
			return
		}
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	code := grand.Digits(4)
	switch req.Scene {
	case 4:
		err = sms.NewSms(ctx).SendCode(to, code)
	case 14:
		err = email.NewEmail(ctx).SendCode(to, code)
	}
	if err != nil {
		return
	}
	err = cache.NewCode(ctx, sceneId, to, req.Scene).Set(code, 5*60)
	return
}
