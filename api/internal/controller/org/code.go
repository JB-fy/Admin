package controller

import (
	"api/api"
	apiCurrent "api/api/org"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoOrg "api/internal/dao/org"
	"api/internal/utils"
	"api/internal/utils/email"
	get_or_set_ctx "api/internal/utils/get-or-set-ctx"
	"api/internal/utils/sms"
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type Code struct{}

func NewCode() *Code {
	return &Code{}
}

// 发送验证码
func (controllerThis *Code) Send(ctx context.Context, req *apiCurrent.CodeSendReq) (res *api.CommonNoDataRes, err error) {
	loginName := daoOrg.Admin.GetLoginName(req.To)
	switch req.Scene {
	case 0, 1, 2, 4:
		err = g.Validator().Rules(`phone`).Data(loginName).Run(ctx)
	case 10, 11, 12, 14:
		err = g.Validator().Rules(`email`).Data(loginName).Run(ctx)
	}
	if err != nil {
		return
	}

	to := req.To
	switch req.Scene {
	case 0, 2: //登录(手机)，密码找回(手机)
		info, _ := daoOrg.Admin.CtxDaoModel(ctx).Filter(daoOrg.Admin.Columns().Phone, to).One()
		if info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990000, ``)
			return
		}
		if info[daoOrg.Admin.Columns().IsStop].Uint8() == 1 {
			err = utils.NewErrorCode(ctx, 39990002, ``)
			return
		}
	case 1: //注册(手机)
		info, _ := daoOrg.Admin.CtxDaoModel(ctx).Filter(daoOrg.Admin.Columns().Phone, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991000, ``)
			return
		}
	case 3: //密码修改(手机)
		loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		/* if loginInfo[daoOrg.Admin.Columns().Phone].String() != `` {
			err = utils.NewErrorCode(ctx, 39991001, ``)
			return
		} */
		to = loginInfo[daoOrg.Admin.Columns().Phone].String()
		if to != `` {
			err = utils.NewErrorCode(ctx, 39991003, ``)
			return
		}
		loginName = daoOrg.Admin.GetLoginName(to)
	case 4: //绑定(手机)
		loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		if loginInfo[daoOrg.Admin.Columns().IsSuper].Uint8() == 0 {
			to = daoOrg.Admin.JoinLoginName(loginInfo[daoOrg.Admin.Columns().OrgId].Uint(), to)
		}
		info, _ := daoOrg.Admin.CtxDaoModel(ctx).Filter(daoOrg.Admin.Columns().Phone, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991002, ``)
			return
		}
	case 5: //解绑(手机)
		loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = loginInfo[daoOrg.Admin.Columns().Phone].String()
		if to == `` {
			err = utils.NewErrorCode(ctx, 39991003, ``)
			return
		}
		loginName = daoOrg.Admin.GetLoginName(to)
	case 10, 12: //登录(邮箱)，密码找回(邮箱)
		info, _ := daoOrg.Admin.CtxDaoModel(ctx).Filter(daoOrg.Admin.Columns().Email, to).One()
		if info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990000, ``)
			return
		}
		if info[daoOrg.Admin.Columns().IsStop].Uint8() == 1 {
			err = utils.NewErrorCode(ctx, 39990002, ``)
			return
		}
	case 11: //注册(邮箱)
		info, _ := daoOrg.Admin.CtxDaoModel(ctx).Filter(daoOrg.Admin.Columns().Email, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991010, ``)
			return
		}
	case 13: //密码修改(邮箱)
		loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = loginInfo[daoOrg.Admin.Columns().Email].String()
		if to != `` {
			err = utils.NewErrorCode(ctx, 39991013, ``)
			return
		}
		loginName = daoOrg.Admin.GetLoginName(to)
	case 14: //绑定(邮箱)
		loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		/* if loginInfo[daoOrg.Admin.Columns().Email].String() != `` {
			err = utils.NewErrorCode(ctx, 39991011, ``)
			return
		} */
		if loginInfo[daoOrg.Admin.Columns().IsSuper].Uint8() == 0 {
			to = daoOrg.Admin.JoinLoginName(loginInfo[daoOrg.Admin.Columns().OrgId].Uint(), to)
		}
		info, _ := daoOrg.Admin.CtxDaoModel(ctx).Filter(daoOrg.Admin.Columns().Email, to).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991012, ``)
			return
		}
	case 15: //解绑(邮箱)
		loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = loginInfo[daoOrg.Admin.Columns().Email].String()
		if to == `` {
			err = utils.NewErrorCode(ctx, 39991013, ``)
			return
		}
		loginName = daoOrg.Admin.GetLoginName(to)
	}

	code := grand.Digits(4)
	switch req.Scene {
	case 0, 1, 2, 3, 4, 5:
		err = sms.NewHandler(ctx).SendCode(loginName, code)
	case 10, 11, 12, 13, 14, 15:
		err = email.NewHandler(ctx).SendCode(loginName, code)
	}
	if err != nil {
		return
	}
	err = cache.Code.Set(ctx, get_or_set_ctx.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), to, req.Scene, code, 5*time.Minute)
	return
}
