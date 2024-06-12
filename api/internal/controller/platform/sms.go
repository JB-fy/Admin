package controller

import (
	"api/api"
	apiCurrent "api/api/app"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"api/internal/utils/sms"
	"context"

	"github.com/gogf/gf/v2/util/grand"
)

type Sms struct{}

func NewSms() *Sms {
	return &Sms{}
}

// 发送短信
func (controllerThis *Sms) Send(ctx context.Context, req *apiCurrent.SmsSendReq) (res *api.CommonNoDataRes, err error) {
	phone := req.Phone
	switch req.UseScene {
	case 4: //绑定手机
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filter(daoPlatform.Admin.Columns().Phone, phone).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990006, ``)
			return
		}
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	smsCode := grand.Digits(4)
	err = sms.NewSms(ctx).Send(phone, smsCode)
	if err != nil {
		return
	}
	err = cache.NewSms(ctx, sceneCode, phone, req.UseScene).Set(smsCode, 5*60)
	return
}
