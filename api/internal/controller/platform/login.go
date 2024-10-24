package controller

import (
	"api/api"
	apiCurrent "api/api/platform"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"api/internal/utils/token"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
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
	if info[daoPlatform.Admin.Columns().IsStop].Uint() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	saltDynamic := grand.S(8)
	err = cache.NewSalt(ctx, sceneCode, req.LoginName).Set(saltDynamic, 5)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{SaltStatic: info[daoPlatform.Admin.Columns().Salt].String(), SaltDynamic: saltDynamic}
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
	if info[daoPlatform.Admin.Columns().IsStop].Uint() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	salt, _ := cache.NewSalt(ctx, sceneCode, req.LoginName).Get()
	if salt == `` || gmd5.MustEncrypt(info[daoPlatform.Admin.Columns().Password].String()+salt) != req.Password {
		err = utils.NewErrorCode(ctx, 39990001, ``)
		return
	}

	tokenInfo := token.TokenInfo{
		LoginId: info[daoPlatform.Admin.Columns().AdminId].String(),
		IP:      g.RequestFromCtx(ctx).GetClientIp(),
	}
	token, err := token.NewHandler(ctx, sceneInfo[daoAuth.Scene.Columns().SceneConfig].Map()[`token_config`].(g.Map), sceneCode).Create(tokenInfo)
	if err != nil {
		return
	}

	res = &api.CommonTokenRes{Token: token}
	return
}
