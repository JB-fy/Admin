package controller

import (
	"api/api"
	apiCurrent "api/api/platform"
	"api/internal/cache"
	"api/internal/consts"
	daoAdmin "api/internal/dao/admin"
	daoConfig "api/internal/dao/config"
	"api/internal/utils"
	"api/internal/utils/jbctx"
	"api/internal/utils/token"
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

// 获取密码盐
func (controllerThis *Login) Salt(ctx context.Context, req *apiCurrent.LoginSaltReq) (res *api.CommonSaltRes, err error) {
	loginName := req.LoginName
	filter := g.Map{daoAdmin.Admin.Columns().AdminType: req.AdminType}
	if g.Validator().Rules(`phone`).Data(loginName).Run(ctx) == nil {
		filter[daoAdmin.Admin.Columns().Phone] = req.LoginName
	} else if g.Validator().Rules(`email`).Data(loginName).Run(ctx) == nil {
		filter[daoAdmin.Admin.Columns().Email] = req.LoginName
	} else if g.Validator().Rules(`regex:^[\p{L}][\p{L}\p{N}_]{3,}$`).Data(loginName).Run(ctx) == nil {
		filter[daoAdmin.Admin.Columns().Account] = req.LoginName
	} else {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	info, _ := daoAdmin.Admin.CtxDaoModel(ctx).Filters(filter).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[daoAdmin.Admin.Columns().IsStop].Uint8() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	saltStatic, _ := daoAdmin.Privacy.CtxDaoModel(ctx).FilterPri(info[daoAdmin.Privacy.Columns().AdminId]).ValueStr(daoAdmin.Privacy.Columns().Salt)
	if saltStatic == `` {
		err = utils.NewErrorCode(ctx, 39990004, ``)
		return
	}
	saltDynamic := grand.S(8)
	err = cache.Salt.Set(ctx, jbctx.GetSceneId(ctx).String(), req.LoginName, req.AdminType, saltDynamic, 5*time.Second)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{SaltStatic: saltStatic, SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *Login) Login(ctx context.Context, req *apiCurrent.LoginLoginReq) (res *api.CommonTokenRes, err error) {
	loginName := daoAdmin.Admin.GetLoginName(req.LoginName)
	filter := g.Map{daoAdmin.Admin.Columns().AdminType: req.AdminType}
	if g.Validator().Rules(`phone`).Data(loginName).Run(ctx) == nil {
		filter[daoAdmin.Admin.Columns().Phone] = req.LoginName
	} else if g.Validator().Rules(`email`).Data(loginName).Run(ctx) == nil {
		filter[daoAdmin.Admin.Columns().Email] = req.LoginName
	} else if g.Validator().Rules(`regex:^[\p{L}][\p{L}\p{N}_]{3,}$`).Data(loginName).Run(ctx) == nil {
		filter[daoAdmin.Admin.Columns().Account] = req.LoginName
	} else {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	info, _ := daoAdmin.Admin.CtxDaoModel(ctx).Filters(filter).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[daoAdmin.Admin.Columns().IsStop].Uint8() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	if req.Password != `` { //密码
		password, _ := daoAdmin.Privacy.CtxDaoModel(ctx).FilterPri(info[daoAdmin.Privacy.Columns().AdminId]).ValueStr(daoAdmin.Privacy.Columns().Password)
		if password == `` {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
		salt, _ := cache.Salt.Get(ctx, jbctx.GetSceneId(ctx).String(), req.LoginName, req.AdminType)
		if salt == `` || gmd5.MustEncrypt(password+salt) != req.Password {
			err = utils.NewErrorCode(ctx, 39990001, ``)
			return
		}
	} else if req.SmsCode != `` { //短信验证码
		phone := info[daoAdmin.Admin.Columns().Phone].String()
		if phone == `` {
			err = utils.NewErrorCode(ctx, 39991003, ``)
			return
		}
		code, _ := cache.Code.Get(ctx, jbctx.GetSceneId(ctx).String(), phone, req.AdminType, 0) //场景：0登录(手机)
		if code == `` || code != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
	} else if req.EmailCode != `` { //邮箱验证码
		email := info[daoAdmin.Admin.Columns().Email].String()
		if email == `` {
			err = utils.NewErrorCode(ctx, 39991013, ``)
			return
		}
		code, _ := cache.Code.Get(ctx, jbctx.GetSceneId(ctx).String(), email, req.AdminType, 10) //场景：10登录(邮箱)
		if code == `` || code != req.EmailCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
	}

	token, err := token.NewHandler(ctx).Create(info[daoAdmin.Admin.Columns().AdminId].String(), nil)
	if err != nil {
		return
	}

	res = &api.CommonTokenRes{Token: token}
	return
}

// 注册
func (controllerThis *Login) Register(ctx context.Context, req *apiCurrent.LoginRegisterReq) (res *api.CommonTokenRes, err error) {
	data := g.Map{daoAdmin.Admin.Columns().AdminType: req.AdminType}
	if req.Phone != `` {
		code, _ := cache.Code.Get(ctx, jbctx.GetSceneId(ctx).String(), req.Phone, req.AdminType, 1) //场景：1注册(手机)
		if code == `` || code != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		info, _ := daoAdmin.Admin.CtxDaoModel(ctx).Filters(map[string]any{daoAdmin.Admin.Columns().Phone: req.Phone, daoAdmin.Admin.Columns().AdminType: req.AdminType}).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991000, ``)
			return
		}
		data[daoAdmin.Admin.Columns().Phone] = req.Phone
		data[daoAdmin.Admin.Columns().Nickname] = req.Phone[:3] + strings.Repeat(`*`, len(req.Phone)-7) + req.Phone[len(req.Phone)-4:]
	}
	if req.Email != `` {
		code, _ := cache.Code.Get(ctx, jbctx.GetSceneId(ctx).String(), req.Email, req.AdminType, 11) //场景：11注册(邮箱)
		if code == `` || code != req.EmailCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		info, _ := daoAdmin.Admin.CtxDaoModel(ctx).Filters(map[string]any{daoAdmin.Admin.Columns().Email: req.Email, daoAdmin.Admin.Columns().AdminType: req.AdminType}).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991010, ``)
			return
		}
		data[daoAdmin.Admin.Columns().Email] = req.Email
		data[daoAdmin.Admin.Columns().Nickname], _, _ = strings.Cut(req.Email, `@`)
	}
	if req.Account != `` {
		info, _ := daoAdmin.Admin.CtxDaoModel(ctx).Filters(map[string]any{daoAdmin.Admin.Columns().Account: req.Account, daoAdmin.Admin.Columns().AdminType: req.AdminType}).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991020, ``)
			return
		}
		data[daoAdmin.Admin.Columns().Account] = req.Account
		accountRune := []rune(req.Account)
		accountRuneLen := len(accountRune)
		data[daoAdmin.Admin.Columns().Nickname] = string(accountRune[:1]) + strings.Repeat(`*`, accountRuneLen-2) + string(accountRune[accountRuneLen-1:])
	}
	if req.Password != `` {
		data[daoAdmin.Privacy.Columns().Password] = req.Password
	}

	switch req.AdminType {
	case 0: //平台
		data[daoAdmin.Admin.Columns().IsSuper] = 0 //不允许注册超级管理员
		data[daoAdmin.Admin.Columns().SceneId] = consts.SCENE_ID_PLATFORM
		data[`role_id_arr`] = daoConfig.Config.Get(ctx, consts.SCENE_ID_PLATFORM, 0, `role_id_arr_of_platform_def`).Slice() //默认角色
	default:
		err = utils.NewErrorCode(ctx, 39999995, ``)
		return
	}
	adminId, err := daoAdmin.Admin.CtxDaoModel(ctx).HookInsert(data).InsertAndGetId()
	if err != nil {
		return
	}

	token, err := token.NewHandler(ctx).Create(gconv.String(adminId), nil)
	if err != nil {
		return
	}

	res = &api.CommonTokenRes{Token: token}
	return
}

// 密码找回
func (controllerThis *Login) PasswordRecovery(ctx context.Context, req *apiCurrent.LoginPasswordRecoveryReq) (res *api.CommonNoDataRes, err error) {
	filter := g.Map{daoAdmin.Admin.Columns().AdminType: req.AdminType}
	if req.Phone != `` {
		if g.Validator().Rules(`phone`).Data(daoAdmin.Admin.GetLoginName(req.Phone)).Run(ctx) != nil {
			err = utils.NewErrorCode(ctx, 89990000, ``)
			return
		}
		code, _ := cache.Code.Get(ctx, jbctx.GetSceneId(ctx).String(), req.Phone, req.AdminType, 2) //场景：2密码找回(手机)
		if code == `` || code != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		filter[daoAdmin.Admin.Columns().Phone] = req.Phone
	} else if req.Email != `` {
		if g.Validator().Rules(`email`).Data(daoAdmin.Admin.GetLoginName(req.Email)).Run(ctx) != nil {
			err = utils.NewErrorCode(ctx, 89990000, ``)
			return
		}
		code, _ := cache.Code.Get(ctx, jbctx.GetSceneId(ctx).String(), req.Email, req.AdminType, 12) //场景：12密码找回(邮箱)
		if code == `` || code != req.EmailCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		filter[daoAdmin.Admin.Columns().Email] = req.Email
	}

	daoModelOrgAdmin := daoAdmin.Admin.CtxDaoModel(ctx).SetIdArr(filter)
	if len(daoModelOrgAdmin.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	_, err = daoModelOrgAdmin.HookUpdateOne(daoAdmin.Privacy.Columns().Password, req.Password).Update()
	return
}
