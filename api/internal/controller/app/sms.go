package controller

import (
	"api/api"
	apiCurrent "api/api/app"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoUsers "api/internal/dao/users"
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
	case 0, 2: //登录，密码找回
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Phone, phone).One()
		if info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990000, ``)
			return
		}
		if info[daoUsers.Users.Columns().IsStop].Uint() == 1 {
			err = utils.NewErrorCode(ctx, 39990002, ``)
			return
		}
	case 1: //注册
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Phone, phone).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
	case 3: //密码修改
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		phone = loginInfo[daoUsers.Users.Columns().Phone].String()
		if phone != `` {
			err = utils.NewErrorCode(ctx, 39990007, ``)
			return
		}
	case 4: //绑定手机
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		if loginInfo[daoUsers.Users.Columns().Phone].String() != `` {
			err = utils.NewErrorCode(ctx, 39990005, ``)
			return
		}
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Phone, phone).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990006, ``)
			return
		}
	case 5: //解绑手机
		loginInfo := utils.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		phone = loginInfo[daoUsers.Users.Columns().Phone].String()
		if phone == `` {
			err = utils.NewErrorCode(ctx, 39990007, ``)
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
