package controller

import (
	"api/api"
	apiCurrent "api/api/org"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoOrg "api/internal/dao/org"
	"api/internal/utils"
	"api/internal/utils/token"
	"context"

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
	loginName := daoOrg.Admin.GetLoginName(req.LoginName)
	filter := g.Map{}
	if g.Validator().Rules(`phone`).Data(loginName).Run(ctx) == nil {
		filter[daoOrg.Admin.Columns().Phone] = req.LoginName
	} else if g.Validator().Rules(`email`).Data(loginName).Run(ctx) == nil {
		filter[daoOrg.Admin.Columns().Email] = req.LoginName
	} else if g.Validator().Rules(`regex:^[\p{L}][\p{L}\p{N}_]{3,}$`).Data(loginName).Run(ctx) == nil {
		filter[daoOrg.Admin.Columns().Account] = req.LoginName
	} else {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	info, _ := daoOrg.Admin.CtxDaoModel(ctx).Filters(filter).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[daoOrg.Admin.Columns().IsStop].Uint8() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	saltDynamic := grand.S(8)
	err = cache.Salt.Set(ctx, sceneId, req.LoginName, saltDynamic, 5)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{SaltStatic: info[daoOrg.Admin.Columns().Salt].String(), SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *Login) Login(ctx context.Context, req *apiCurrent.LoginLoginReq) (res *api.CommonTokenRes, err error) {
	loginName := daoOrg.Admin.GetLoginName(req.LoginName)
	filter := g.Map{}
	if g.Validator().Rules(`phone`).Data(loginName).Run(ctx) == nil {
		filter[daoOrg.Admin.Columns().Phone] = req.LoginName
	} else if g.Validator().Rules(`email`).Data(loginName).Run(ctx) == nil {
		filter[daoOrg.Admin.Columns().Email] = req.LoginName
	} else if g.Validator().Rules(`regex:^[\p{L}][\p{L}\p{N}_]{3,}$`).Data(loginName).Run(ctx) == nil {
		filter[daoOrg.Admin.Columns().Account] = req.LoginName
	} else {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	info, _ := daoOrg.Admin.CtxDaoModel(ctx).Filters(filter).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[daoOrg.Admin.Columns().IsStop].Uint8() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	salt, _ := cache.Salt.Get(ctx, sceneId, req.LoginName)
	if salt == `` || gmd5.MustEncrypt(info[daoOrg.Admin.Columns().Password].String()+salt) != req.Password {
		err = utils.NewErrorCode(ctx, 39990001, ``)
		return
	}

	token, err := token.NewHandler(ctx).Create(info[daoOrg.Admin.Columns().AdminId].String(), nil)
	if err != nil {
		return
	}

	res = &api.CommonTokenRes{Token: token}
	return
}

// 注册
func (controllerThis *Login) Register(ctx context.Context, req *apiCurrent.LoginRegisterReq) (res *api.CommonTokenRes, err error) {
	data := g.Map{}
	data[daoOrg.Admin.Columns().IsSuper] = 1 //只允许注册超级管理员
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	if req.Phone != `` {
		code, _ := cache.Code.Get(ctx, sceneId, req.Phone, 1) //场景：1注册(手机)
		if code == `` || code != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}

		info, _ := daoOrg.Admin.CtxDaoModel(ctx).Filter(daoOrg.Admin.Columns().Phone, req.Phone).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991000, ``)
			return
		}
		data[daoOrg.Admin.Columns().Phone] = req.Phone
		data[daoOrg.Admin.Columns().Nickname] = req.Phone[:3] + `****` + req.Phone[len(req.Phone)-4:]
	}
	if req.Email != `` {
		code, _ := cache.Code.Get(ctx, sceneId, req.Email, 11) //场景：11注册(邮箱)
		if code == `` || code != req.EmailCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}

		info, _ := daoOrg.Admin.CtxDaoModel(ctx).Filter(daoOrg.Admin.Columns().Email, req.Email).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991010, ``)
			return
		}
		data[daoOrg.Admin.Columns().Email] = req.Email
		data[daoOrg.Admin.Columns().Nickname] = gstr.Split(req.Email, `@`)[0]
	}
	if req.Account != `` {
		info, _ := daoOrg.Admin.CtxDaoModel(ctx).Filter(daoOrg.Admin.Columns().Account, req.Account).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991020, ``)
			return
		}
		data[daoOrg.Admin.Columns().Account] = req.Account
		data[daoOrg.Admin.Columns().Nickname] = req.Account
	}
	if req.Password != `` {
		data[daoOrg.Admin.Columns().Password] = req.Password
	}

	adminId, err := daoOrg.Admin.CtxDaoModel(ctx).HookInsert(data).InsertAndGetId()
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
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	filter := g.Map{}
	if req.Phone != `` {
		if g.Validator().Rules(`phone`).Data(daoOrg.Admin.GetLoginName(req.Phone)).Run(ctx) != nil {
			err = utils.NewErrorCode(ctx, 89990000, ``)
			return
		}
		code, _ := cache.Code.Get(ctx, sceneId, req.Phone, 2) //场景：2密码找回(手机)
		if code == `` || code != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		filter[daoOrg.Admin.Columns().Phone] = req.Phone
	} else if req.Email != `` {
		if g.Validator().Rules(`email`).Data(daoOrg.Admin.GetLoginName(req.Email)).Run(ctx) != nil {
			err = utils.NewErrorCode(ctx, 89990000, ``)
			return
		}
		code, _ := cache.Code.Get(ctx, sceneId, req.Email, 12) //场景：12密码找回(邮箱)
		if code == `` || code != req.EmailCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		filter[daoOrg.Admin.Columns().Email] = req.Email
	}

	daoModelOrgAdmin := daoOrg.Admin.CtxDaoModel(ctx).SetIdArr(filter)
	if len(daoModelOrgAdmin.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	_, err = daoModelOrgAdmin.HookUpdateOne(daoOrg.Admin.Columns().Password, req.Password).Update()
	if err != nil {
		return
	}
	return
}
