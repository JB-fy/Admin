package logic

import (
	"api/internal/consts"
	"api/internal/dao"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type sLoginPlatformAdmin struct{}

func NewLoginPlatformAdmin() *sLoginPlatformAdmin {
	return &sLoginPlatformAdmin{}
}

func init() {
	service.RegisterLoginPlatformAdmin(NewLoginPlatformAdmin())
}

// 获取加密盐
func (logicThis *sLoginPlatformAdmin) Salt(ctx context.Context, loginName string) (saltStatic string, saltDynamic string, err error) {
	sceneCode := `platform` //指定场景
	info, _ := dao.NewDaoHandler(ctx, &daoPlatform.Admin).Filter(g.Map{`loginName`: loginName, `isStop`: 0}).GetModel().One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[`isStop`].Int() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}
	saltStatic = info[`salt`].String()
	saltKey := fmt.Sprintf(consts.CacheSaltFormat, sceneCode, loginName)
	saltDynamic = grand.S(8)
	err = g.Redis().SetEX(ctx, saltKey, saltDynamic, 5)
	return
}

// 登录
func (logicThis *sLoginPlatformAdmin) Login(ctx context.Context, loginName string, password string) (token string, err error) {
	sceneCode := `platform` //指定场景
	info, _ := dao.NewDaoHandler(ctx, &daoPlatform.Admin).Filter(g.Map{`loginName`: loginName}).GetModel().One()
	if len(info) == 0 {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[`isStop`].Int() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}
	saltKey := fmt.Sprintf(consts.CacheSaltFormat, sceneCode, loginName)
	saltVar, _ := g.Redis().Get(ctx, saltKey)
	salt := saltVar.String()
	if salt == `` || gmd5.MustEncrypt(info[`password`].String()+salt) != password {
		err = utils.NewErrorCode(ctx, 39990001, ``)
		return
	}

	claims := utils.CustomClaims{
		LoginId: info[`adminId`].Uint(),
	}
	jwt := utils.NewJWT(ctx, utils.GetCtxSceneInfo(ctx)[`sceneConfig`].Map())
	token, err = jwt.CreateToken(claims)
	/* //缓存token（选做。限制多地登录，多设备登录等情况下可用）
	TokenKey := fmt.Sprintf(consts.CacheTokenFormat, sceneCode, claims.LoginId)
	g.Redis().SetEX(ctx, TokenKey, token, int64(jwt.ExpireTime)) */
	return
}
