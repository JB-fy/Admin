package logic

import (
	"api/internal/consts"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type sLogin struct{}

func NewLogin() *sLogin {
	return &sLogin{}
}

func init() {
	service.RegisterLogin(NewLogin())
}

// 获取登录加密字符串(前端登录操作用于加密密码后提交)
func (logicThis *sLogin) Salt(ctx context.Context, account string) (salt string, err error) {
	sceneCode := `platform` //指定场景
	info, _ := daoPlatform.Admin.ParseDbCtx(ctx).Handler(daoPlatform.Admin.ParseFilter(map[string]interface{}{`accountOrPhone`: account}, &[]string{})).One()
	if len(info) == 0 {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	saltKey := fmt.Sprintf(consts.CacheSaltFormat, sceneCode, account)
	salt = grand.S(8)
	g.Redis().SetEX(ctx, saltKey, salt, 5)
	return
}

// 登录(平台后台管理员)
func (logicThis *sLogin) PlatformAdmin(ctx context.Context, account string, password string) (token string, err error) {
	sceneCode := `platform` //指定场景
	/**--------验证账号密码 开始--------**/
	info, _ := daoPlatform.Admin.ParseDbCtx(ctx).Handler(daoPlatform.Admin.ParseFilter(map[string]interface{}{`accountOrPhone`: account}, &[]string{})).One()
	if len(info) == 0 {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[`isStop`].Int() > 0 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}
	saltKey := fmt.Sprintf(consts.CacheSaltFormat, sceneCode, account)
	salt, _ := g.Redis().Get(ctx, saltKey)
	if salt.String() == `` || gmd5.MustEncrypt(info[`password`].String()+salt.String()) != password {
		err = utils.NewErrorCode(ctx, 39990001, ``)
		return
	}
	/**--------验证账号密码 结束--------**/

	claims := utils.CustomClaims{
		LoginId:  info[`adminId`].Uint(),
		Account:  info[`account`].String(),
		Nickname: info[`nickname`].String(),
	}
	jwt := utils.NewJWT(ctx, utils.GetCtxSceneInfo(ctx)[`sceneConfig`].Map())
	token, err = jwt.CreateToken(claims)
	/* //缓存token（选做。限制多地登录，多设备登录等情况下可用）
	TokenKey := fmt.Sprintf(consts.CacheTokenFormat, sceneCode, claims.Account)
	g.Redis().SetEX(ctx, TokenKey, token, int64(jwt.ExpireTime)) */
	return
}
