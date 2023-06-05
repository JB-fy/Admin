package logic

import (
	"api/internal/consts"
	daoPlatform "api/internal/model/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/database/gredis"
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
	encryptStr = utils.RandomStr(8)
	gredis.Instance().SetEX(ctx, encryptStrKey, encryptStr, 5)
	return
}

// 登录
func (logicThis *sLogin) Login(ctx context.Context, sceneCode string, account string, password string) (token string, err error) {
	switch sceneCode {
	case "platformAdmin":
		/**--------验证账号密码 开始--------**/
		info, _ := daoPlatform.Admin.Ctx(ctx).Handler(daoPlatform.Admin.ParseFilter(map[string]interface{}{"accountOrPhone": account}, &[]string{})).One()
		if len(info) == 0 {
			err = errors.New("39990000")
			return
		}
		if info["isStop"].Int() > 0 {
			err = errors.New("39990001")
			return
		}
		encryptStrKey := fmt.Sprintf(consts.CacheEncryptStrFormat, sceneCode, account)
		encryptStr, errTmp := gredis.Instance().Get(ctx, encryptStrKey)
		fmt.Println(errTmp)
		if errTmp != nil {
			err = errors.New("39990000")
			return
		}
		if utils.Md5(info["password"].String()+encryptStr.String()) != password {
			err = errors.New("39990000")
			return
		}
		/**--------验证账号密码 结束--------**/

		/* $payload = [
		       'id' => $info->adminId
		   ];
		   $jwt = make($sceneCode . 'Jwt');
		   $token = $jwt->createToken($payload);

		   //缓存token（选做。限制多地登录，多设备登录等情况下可用）
		   $cacheLogin = getCache(CacheLogin::class);
		   $cacheLogin->setTokenKey($payload['id'], $sceneCode);
		   $cacheLogin->setToken($token, $jwt->getConfig()['expireTime']);

		   throwSuccessJson(['token' => $token]); */
	}
	return
}
