package controller

import (
	"api/api"
	apiLogin "api/api/platform/login"
	"api/internal/consts"
	"api/internal/dao"
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type Admin struct{}

func NewAdmin() *Admin {
	return &Admin{}
}

// 获取加密盐
func (controllerThis *Admin) Salt(ctx context.Context, req *apiLogin.AdminSaltReq) (res *api.CommonSaltRes, err error) {
	info, _ := dao.NewDaoHandler(ctx, &daoPlatform.Admin).Filter(g.Map{`loginName`: req.LoginName}).GetModel().One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[`isStop`].Int() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[`sceneCode`].String()
	saltKey := fmt.Sprintf(consts.CacheSaltFormat, sceneCode, req.LoginName)
	saltDynamic := grand.S(8)
	err = g.Redis().SetEX(ctx, saltKey, saltDynamic, 5)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{SaltStatic: info[`salt`].String(), SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *Admin) Login(ctx context.Context, req *apiLogin.AdminLoginReq) (res *api.CommonTokenRes, err error) {
	info, _ := dao.NewDaoHandler(ctx, &daoPlatform.Admin).Filter(g.Map{`loginName`: req.LoginName}).GetModel().One()
	if len(info) == 0 {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[`isStop`].Int() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[`sceneCode`].String()
	saltKey := fmt.Sprintf(consts.CacheSaltFormat, sceneCode, req.LoginName)
	saltVar, _ := g.Redis().Get(ctx, saltKey)
	salt := saltVar.String()
	if salt == `` || gmd5.MustEncrypt(info[`password`].String()+salt) != req.Password {
		err = utils.NewErrorCode(ctx, 39990001, ``)
		return
	}

	claims := utils.CustomClaims{LoginId: info[`adminId`].Uint()}
	jwt := utils.NewJWT(ctx, sceneInfo[`sceneConfig`].Map())
	token, err := jwt.CreateToken(claims)
	if err != nil {
		return
	}
	/* //缓存token（选做。限制多地登录，多设备登录等情况下可用）
	tokenKey := fmt.Sprintf(consts.CacheTokenFormat, sceneCode, claims.LoginId)
	g.Redis().SetEX(ctx, tokenKey, token, int64(jwt.ExpireTime)) */
	res = &api.CommonTokenRes{Token: token}
	return
}
