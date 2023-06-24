package middleware

import (
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"

	"github.com/gogf/gf/v2/net/ghttp"
)

func SceneLoginOfPlatform(r *ghttp.Request) {
	/**--------验证token 开始--------**/
	token := r.Header.Get("PlatformToken")
	if token == "" {
		utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 39994000, ""))
		return
	}

	jwt := utils.NewJWT(r.GetCtx(), utils.GetCtxSceneInfo(r.GetCtx())["sceneConfig"].Map())
	claims, err := jwt.ParseToken(token)
	if err != nil {
		utils.HttpFailJson(r, err)
		return
	}
	/**--------验证token 结束--------**/

	/**--------选做。限制多地登录，多设备登录等情况下可用（前提必须在登录时做过token缓存） 开始--------**/
	/* TokenKey := fmt.Sprintf(consts.CacheTokenFormat, sceneCode, claims.Account)
	checkToken, _ := g.Redis().Get(r.GetCtx(), TokenKey)
	if checkToken.String() != token {
		utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 39994002, ""))
		return
	} */
	/**--------选做。限制多地登录，多设备登录等情况下可用（前提必须在登录时做过token缓存） 结束--------**/

	/**--------获取登录用户信息并验证 开始--------**/
	info, _ := daoPlatform.Admin.ParseDbCtx(r.GetCtx()).Where("adminId", claims.LoginId).One()
	if len(info) == 0 {
		utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 39994003, ""))
		return
	}
	if info["isStop"].Int() > 0 {
		utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 39994004, ""))
		return
	}
	delete(info, "password")
	delete(info, "isStop")

	utils.SetCtxLoginInfo(r, info) //用户信息保存在协程上下文
	/**--------获取用户信息并验证 结束--------**/

	r.Middleware.Next()
}
