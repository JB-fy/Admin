package controller

import (
	"api/api"
	apiCurrent "api/api/platform"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/grand"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

// 获取加密盐
func (controllerThis *Login) Salt(ctx context.Context, req *apiCurrent.LoginSaltReq) (res *api.CommonSaltRes, err error) {
	adminColumns := daoPlatform.Admin.Columns()
	info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filter(`loginName`, req.LoginName).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[adminColumns.IsStop].Uint() == 1 {
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
	res = &api.CommonSaltRes{SaltStatic: info[adminColumns.Salt].String(), SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *Login) Login(ctx context.Context, req *apiCurrent.LoginLoginReq) (res *api.CommonTokenRes, err error) {
	adminColumns := daoPlatform.Admin.Columns()
	info, _ := daoPlatform.Admin.CtxDaoModel(ctx).Filter(`loginName`, req.LoginName).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[adminColumns.IsStop].Uint() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	salt, _ := cache.NewSalt(ctx, sceneCode, req.LoginName).Get()
	if salt == `` || gmd5.MustEncrypt(info[adminColumns.Password].String()+salt) != req.Password {
		err = utils.NewErrorCode(ctx, 39990001, ``)
		return
	}

	claims := utils.CustomClaims{LoginId: info[daoPlatform.Admin.PrimaryKey()].Uint()}
	jwt := utils.NewJWT(ctx, sceneInfo[daoAuth.Scene.Columns().SceneConfig].Map())
	token, err := jwt.CreateToken(claims)
	if err != nil {
		return
	}
	// cache.NewToken(ctx, sceneCode, claims.LoginId).Set(token, int64(jwt.ExpireTime)) //缓存token（限制多地登录，多设备登录等情况下用）

	res = &api.CommonTokenRes{Token: token}
	return
}
