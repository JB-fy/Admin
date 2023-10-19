package controller

import (
	"api/api"
	apiCurrent "api/api/app"
	"api/internal/consts"
	"api/internal/dao"
	daoUser "api/internal/dao/user"
	"api/internal/utils"
	"api/internal/utils/sms"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
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
		info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{`phone`: phone}).GetModel().One()
		if info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990000, ``)
			return
		}
		if info[`isStop`].Int() == 1 {
			err = utils.NewErrorCode(ctx, 39990002, ``)
			return
		}
	case 1: //注册
		info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{`phone`: phone}).GetModel().One()
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
		phone = loginInfo[`phone`].String()
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
		if loginInfo[`phone`].String() != `` {
			err = utils.NewErrorCode(ctx, 39990005, ``)
			return
		}
		info, _ := dao.NewDaoHandler(ctx, &daoUser.User).Filter(g.Map{`phone`: phone}).GetModel().One()
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
		phone = loginInfo[`phone`].String()
		if phone == `` {
			err = utils.NewErrorCode(ctx, 39990007, ``)
			return
		}
	}

	smsCode := grand.Digits(4)
	err = sms.NewSms(ctx).Send(phone, smsCode)
	if err != nil {
		return
	}
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[`sceneCode`].String()
	smsKey := fmt.Sprintf(consts.CacheSmsFormat, sceneCode, phone, req.UseScene)
	err = g.Redis().SetEX(ctx, smsKey, smsCode, 5*60)
	return
}
