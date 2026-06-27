package controller

import (
	"api/api"
	apiCurrent "api/api/org"
	"api/internal/cache"
	daoAdmin "api/internal/dao/admin"
	"api/internal/utils"
	"api/internal/utils/email"
	"api/internal/utils/jbctx"
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
	loginName := daoAdmin.Admin.GetLoginName(req.To)
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
	adminType := req.AdminType
	switch req.Scene {
	case 0, 2: //登录(手机)，密码找回(手机)
		info, _ := daoAdmin.Admin.CtxDaoModel(ctx).Filters(map[string]any{daoAdmin.Admin.Columns().Phone: to, daoAdmin.Admin.Columns().AdminType: adminType}).One()
		if info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990000, ``)
			return
		}
		if info[daoAdmin.Admin.Columns().IsStop].Uint8() == 1 {
			err = utils.NewErrorCode(ctx, 39990002, ``)
			return
		}
	case 1: //注册(手机)
		info, _ := daoAdmin.Admin.CtxDaoModel(ctx).Filters(map[string]any{daoAdmin.Admin.Columns().Phone: to, daoAdmin.Admin.Columns().AdminType: adminType}).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991000, ``)
			return
		}
	case 3: //密码修改(手机)
		loginInfo := jbctx.GetLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		/* if loginInfo[daoAdmin.Admin.Columns().Phone].String() != `` {
			err = utils.NewErrorCode(ctx, 39991001, ``)
			return
		} */
		to = loginInfo[daoAdmin.Admin.Columns().Phone].String()
		if to != `` {
			err = utils.NewErrorCode(ctx, 39991003, ``)
			return
		}
		loginName = daoAdmin.Admin.GetLoginName(to)
		adminType = loginInfo[daoAdmin.Admin.Columns().AdminType].Uint8()
	case 4: //绑定(手机)
		loginInfo := jbctx.GetLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = daoAdmin.Admin.JoinLoginName(loginInfo[daoAdmin.Admin.Columns().RelId].Uint(), loginInfo[daoAdmin.Admin.Columns().IsSuper].Uint8(), to)
		adminType = loginInfo[daoAdmin.Admin.Columns().AdminType].Uint8()
		info, _ := daoAdmin.Admin.CtxDaoModel(ctx).Filters(map[string]any{daoAdmin.Admin.Columns().Phone: to, daoAdmin.Admin.Columns().AdminType: adminType}).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991002, ``)
			return
		}
	case 5: //解绑(手机)
		loginInfo := jbctx.GetLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = loginInfo[daoAdmin.Admin.Columns().Phone].String()
		if to == `` {
			err = utils.NewErrorCode(ctx, 39991003, ``)
			return
		}
		loginName = daoAdmin.Admin.GetLoginName(to)
		adminType = loginInfo[daoAdmin.Admin.Columns().AdminType].Uint8()
	case 10, 12: //登录(邮箱)，密码找回(邮箱)
		info, _ := daoAdmin.Admin.CtxDaoModel(ctx).Filters(map[string]any{daoAdmin.Admin.Columns().Email: to, daoAdmin.Admin.Columns().AdminType: adminType}).One()
		if info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990000, ``)
			return
		}
		if info[daoAdmin.Admin.Columns().IsStop].Uint8() == 1 {
			err = utils.NewErrorCode(ctx, 39990002, ``)
			return
		}
	case 11: //注册(邮箱)
		info, _ := daoAdmin.Admin.CtxDaoModel(ctx).Filters(map[string]any{daoAdmin.Admin.Columns().Email: to, daoAdmin.Admin.Columns().AdminType: adminType}).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991010, ``)
			return
		}
	case 13: //密码修改(邮箱)
		loginInfo := jbctx.GetLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = loginInfo[daoAdmin.Admin.Columns().Email].String()
		if to != `` {
			err = utils.NewErrorCode(ctx, 39991013, ``)
			return
		}
		loginName = daoAdmin.Admin.GetLoginName(to)
		adminType = loginInfo[daoAdmin.Admin.Columns().AdminType].Uint8()
	case 14: //绑定(邮箱)
		loginInfo := jbctx.GetLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		/* if loginInfo[daoAdmin.Admin.Columns().Email].String() != `` {
			err = utils.NewErrorCode(ctx, 39991011, ``)
			return
		} */
		to = daoAdmin.Admin.JoinLoginName(loginInfo[daoAdmin.Admin.Columns().RelId].Uint(), loginInfo[daoAdmin.Admin.Columns().IsSuper].Uint8(), to)
		adminType = loginInfo[daoAdmin.Admin.Columns().AdminType].Uint8()
		info, _ := daoAdmin.Admin.CtxDaoModel(ctx).Filters(map[string]any{daoAdmin.Admin.Columns().Email: to, daoAdmin.Admin.Columns().AdminType: adminType}).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991012, ``)
			return
		}
	case 15: //解绑(邮箱)
		loginInfo := jbctx.GetLoginInfo(ctx)
		if loginInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39994000, ``)
			return
		}
		to = loginInfo[daoAdmin.Admin.Columns().Email].String()
		if to == `` {
			err = utils.NewErrorCode(ctx, 39991013, ``)
			return
		}
		loginName = daoAdmin.Admin.GetLoginName(to)
		adminType = loginInfo[daoAdmin.Admin.Columns().AdminType].Uint8()
	}

	code := grand.Digits(4)
	switch req.Scene {
	case 0, 1, 2, 3, 4, 5:
		err = sms.NewHandler(ctx).SendCode(loginName, code)
	case 10, 11, 12, 13, 14, 15:
		err = email.NewHandler(ctx).SendCode(loginName, code)
	default:
		err = utils.NewErrorCode(ctx, 39999995, ``)
		return
	}
	if err != nil {
		return
	}
	err = cache.Code.Set(ctx, jbctx.GetSceneId(ctx).String(), to, adminType, req.Scene, code, 5*time.Minute)
	return
}
