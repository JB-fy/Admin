package controller

import (
	"api/api"
	apiCurrent "api/api/platform"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"api/internal/utils/jbctx"
	"api/internal/utils/token"
	"context"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

// 获取密码盐
func (controllerThis *Login) Salt(ctx context.Context, req *apiCurrent.LoginSaltReq) (res *api.CommonSaltRes, err error) {
	filter := g.Map{}
	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) == nil {
		filter[daoPlatform.Admin.Columns().Phone] = req.LoginName
	} else if g.Validator().Rules(`email`).Data(req.LoginName).Run(ctx) == nil {
		filter[daoPlatform.Admin.Columns().Email] = req.LoginName
	} else if g.Validator().Rules(`regex:^[\p{L}][\p{L}\p{N}_]{3,}$`).Data(req.LoginName).Run(ctx) == nil {
		filter[daoPlatform.Admin.Columns().Account] = req.LoginName
	} else {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filters(filter).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[daoPlatform.Admin.Columns().IsStop].Uint8() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	saltStatic, _ := daoPlatform.AdminPrivacy.CtxDaoModel(ctx).FilterPri(info[daoPlatform.Admin.Columns().AdminId]).ValueStr(daoPlatform.AdminPrivacy.Columns().Salt)
	if saltStatic == `` {
		err = utils.NewErrorCode(ctx, 39990004, ``)
		return
	}
	saltDynamic := grand.S(8)
	err = cache.Salt.Set(ctx, jbctx.GetSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), req.LoginName, saltDynamic, 5*time.Second)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{SaltStatic: saltStatic, SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *Login) Login(ctx context.Context, req *apiCurrent.LoginLoginReq) (res *api.CommonTokenRes, err error) {
	filter := g.Map{}
	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) == nil {
		filter[daoPlatform.Admin.Columns().Phone] = req.LoginName
	} else if g.Validator().Rules(`email`).Data(req.LoginName).Run(ctx) == nil {
		filter[daoPlatform.Admin.Columns().Email] = req.LoginName
	} else if g.Validator().Rules(`regex:^[\p{L}][\p{L}\p{N}_]{3,}$`).Data(req.LoginName).Run(ctx) == nil {
		filter[daoPlatform.Admin.Columns().Account] = req.LoginName
	} else {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filters(filter).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[daoPlatform.Admin.Columns().IsStop].Uint8() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	sceneInfo := jbctx.GetSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	if req.Password != `` { //密码
		password, _ := daoPlatform.AdminPrivacy.CtxDaoModel(ctx).FilterPri(info[daoPlatform.Admin.Columns().AdminId]).ValueStr(daoPlatform.AdminPrivacy.Columns().Password)
		if password == `` {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
		salt, _ := cache.Salt.Get(ctx, sceneId, req.LoginName)
		if salt == `` || gmd5.MustEncrypt(password+salt) != req.Password {
			err = utils.NewErrorCode(ctx, 39990001, ``)
			return
		}
	} else if req.SmsCode != `` { //短信验证码
		phone := info[daoPlatform.Admin.Columns().Phone].String()
		if phone == `` {
			err = utils.NewErrorCode(ctx, 39991003, ``)
			return
		}
		code, _ := cache.Code.Get(ctx, sceneId, phone, 0) //场景：0登录(手机)
		if code == `` || code != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
	} else if req.EmailCode != `` { //邮箱验证码
		email := info[daoPlatform.Admin.Columns().Email].String()
		if email == `` {
			err = utils.NewErrorCode(ctx, 39991013, ``)
			return
		}
		code, _ := cache.Code.Get(ctx, sceneId, email, 10) //场景：10登录(邮箱)
		if code == `` || code != req.EmailCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
	}

	token, err := token.NewHandler(ctx).Create(info[daoPlatform.Admin.Columns().AdminId].String(), nil)
	if err != nil {
		return
	}

	res = &api.CommonTokenRes{Token: token}
	return
}

// 注册
func (controllerThis *Login) Register(ctx context.Context, req *apiCurrent.LoginRegisterReq) (res *api.CommonTokenRes, err error) {
	data := g.Map{}
	sceneInfo := jbctx.GetSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	if req.Phone != `` {
		code, _ := cache.Code.Get(ctx, sceneId, req.Phone, 1) //场景：1注册(手机)
		if code == `` || code != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filter(daoPlatform.Admin.Columns().Phone, req.Phone).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991000, ``)
			return
		}
		data[daoPlatform.Admin.Columns().Phone] = req.Phone
		data[daoPlatform.Admin.Columns().Nickname] = req.Phone[:3] + gstr.Repeat(`*`, len(req.Phone)-7) + req.Phone[len(req.Phone)-4:]
	}
	if req.Email != `` {
		code, _ := cache.Code.Get(ctx, sceneId, req.Email, 11) //场景：11注册(邮箱)
		if code == `` || code != req.EmailCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filter(daoPlatform.Admin.Columns().Email, req.Email).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991010, ``)
			return
		}
		data[daoPlatform.Admin.Columns().Email] = req.Email
		data[daoPlatform.Admin.Columns().Nickname] = gstr.Split(req.Email, `@`)[0]
	}
	if req.Account != `` {
		info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filter(daoPlatform.Admin.Columns().Account, req.Account).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991020, ``)
			return
		}
		data[daoPlatform.Admin.Columns().Account] = req.Account
		accountRune := []rune(req.Account)
		accountRuneLen := len(accountRune)
		data[daoPlatform.Admin.Columns().Nickname] = string(accountRune[:1]) + gstr.Repeat(`*`, accountRuneLen-2) + string(accountRune[accountRuneLen-1:])
	}
	if req.Password != `` {
		data[daoPlatform.AdminPrivacy.Columns().Password] = req.Password
	}

	data[`role_id_arr`] = daoPlatform.Config.Get(ctx, `role_id_arr_of_platform_def`).Slice() //默认角色
	adminId, err := daoPlatform.Admin.CtxDaoModel(ctx).HookInsert(data).InsertAndGetId()
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
	sceneInfo := jbctx.GetSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	filter := g.Map{}
	if req.Phone != `` {
		if g.Validator().Rules(`phone`).Data(req.Phone).Run(ctx) != nil {
			err = utils.NewErrorCode(ctx, 89990000, ``)
			return
		}
		code, _ := cache.Code.Get(ctx, sceneId, req.Phone, 2) //场景：2密码找回(手机)
		if code == `` || code != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		filter[daoPlatform.Admin.Columns().Phone] = req.Phone
	} else if req.Email != `` {
		if g.Validator().Rules(`email`).Data(req.Email).Run(ctx) != nil {
			err = utils.NewErrorCode(ctx, 89990000, ``)
			return
		}
		code, _ := cache.Code.Get(ctx, sceneId, req.Email, 12) //场景：12密码找回(邮箱)
		if code == `` || code != req.EmailCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		filter[daoPlatform.Admin.Columns().Email] = req.Email
	}

	daoModelOrgAdmin := daoPlatform.Admin.CtxDaoModel(ctx).SetIdArr(filter)
	if len(daoModelOrgAdmin.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	_, err = daoModelOrgAdmin.HookUpdateOne(daoPlatform.AdminPrivacy.Columns().Password, req.Password).Update()
	return
}
