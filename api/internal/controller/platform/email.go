package controller

import (
	"api/api"
	apiCurrent "api/api/app"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"api/internal/utils/email"
	"context"

	"github.com/gogf/gf/v2/util/grand"
)

type Email struct{}

func NewEmail() *Email {
	return &Email{}
}

// 发送短信
func (controllerThis *Email) Send(ctx context.Context, req *apiCurrent.EmailSendReq) (res *api.CommonNoDataRes, err error) {
	toEmail := req.Email
	switch req.UseScene {
	case 4: //绑定邮箱
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filter(daoPlatform.Admin.Columns().Phone, toEmail).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990006, ``)
			return
		}
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	emailCode := grand.Digits(4)
	err = email.NewEmail(ctx).SendCode(toEmail, emailCode)
	if err != nil {
		return
	}
	err = cache.NewEmail(ctx, sceneCode, toEmail, req.UseScene).Set(emailCode, 5*60)
	return
}
