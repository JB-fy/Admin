package logic

import (
	"api/internal/consts"
	daoPlatform "api/internal/model/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"errors"
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
func (logicThis *sLogin) EncryptStr(ctx context.Context, sceneCode string, account string) (encryptStr string, err error) {
	encryptStrKey := fmt.Sprintf(consts.CacheEncryptStrFormat, sceneCode, account)
	encryptStr = grand.S(8)
	g.Redis().SetEX(ctx, encryptStrKey, encryptStr, 5)
	return
}

// 登录
func (logicThis *sLogin) Login(ctx context.Context, sceneCode string, account string, password string) (token string, err error) {
	switch sceneCode {
	case "platformAdmin":
		/**--------验证账号密码 开始--------**/
		info, _ := daoPlatform.Admin.ParseDbCtx(ctx).Handler(daoPlatform.Admin.ParseFilter(map[string]interface{}{"accountOrPhone": account}, &[]string{})).One()
		if len(info) == 0 {
			err = errors.New("39990000")
			return
		}
		if info["isStop"].Int() > 0 {
			err = errors.New("39990001")
			return
		}
		encryptStrKey := fmt.Sprintf(consts.CacheEncryptStrFormat, sceneCode, account)
		encryptStr, _ := g.Redis().Get(ctx, encryptStrKey)
		if encryptStr.String() == "" || gmd5.MustEncrypt(info["password"].String()+encryptStr.String()) != password {
			err = errors.New("39990000")
			return
		}
		/**--------验证账号密码 结束--------**/

		claims := utils.CustomClaims{
			LoginId:  info["adminId"].Uint(),
			Account:  info["account"].String(),
			Nickname: info["nickname"].String(),
		}
		jwt := utils.NewJWT(ctx, utils.GetCtxSceneInfo(ctx)["sceneConfig"].Map())
		token, err = jwt.CreateToken(claims)
		/* //缓存token（选做。限制多地登录，多设备登录等情况下可用）
		TokenKey := fmt.Sprintf(consts.CacheTokenFormat, sceneCode, claims.Account)
		g.Redis().SetEX(ctx, TokenKey, token, int64(jwt.ExpireTime)) */
		return
	}
	return
}
