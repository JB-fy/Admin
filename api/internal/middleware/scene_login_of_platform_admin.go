package middleware

import (
	"api/internal/utils"

	daoPlatform "api/internal/model/dao/platform"

	"github.com/gogf/gf/v2/net/ghttp"
)

func SceneLoginOfPlatformAdmin(r *ghttp.Request) {
	/**--------验证token 开始--------**/
	token := r.Header.Get("PlatformAdminToken")
	if token == "" {
		r.Response.WriteJson(map[string]interface{}{
			"code": 39994000,
			"msg":  "token错误",
			"data": map[string]interface{}{},
		})
		return
	}

	jwt := utils.NewPlatformAdminJWT()
	claims, err := jwt.ParseToken(token)
	if err != nil {
		r.Response.WriteJson(map[string]interface{}{
			"code": 39994000,
			"msg":  err.Error(),
			"data": map[string]interface{}{},
		})
		return
	}
	/**--------验证token 结束--------**/

	/**--------选做。限制多地登录，多设备登录等情况下可用（前提必须在登录时做过token缓存） 开始--------**/
	/* TokenKey := fmt.Sprintf(consts.CacheTokenFormat, sceneCode, claims.Account)
	checkToken, _ := g.Redis().Get(r.GetCtx(), TokenKey)
	if checkToken.String() != token {
		r.Response.WriteJson(map[string]interface{}{
			"code": 39994002,
			"msg":  err.Error(),
			"data": map[string]interface{}{},
		})
		return
	} */
	/**--------选做。限制多地登录，多设备登录等情况下可用（前提必须在登录时做过token缓存） 结束--------**/

	/**--------获取登录用户信息并验证 开始--------**/

	info, _ := daoPlatform.Admin.Ctx(r.GetCtx()).Handler(daoPlatform.Admin.ParseFilter(map[string]interface{}{"adminId": claims.LoginId}, &[]string{})).One()

	if len(info) == 0 {
		r.Response.WriteJson(map[string]interface{}{
			"code": 39994003,
			"msg":  err.Error(),
			"data": map[string]interface{}{},
		})
		return
	}
	if info["isStop"].Int() > 0 {
		r.Response.WriteJson(map[string]interface{}{
			"code": 39994004,
			"msg":  err.Error(),
			"data": map[string]interface{}{},
		})
		return
	}
	delete(info, "password")
	delete(info, "isStop")

	utils.SetCtxLoginInfo(r, info) //用户信息保存在协程上下文
	/**--------获取用户信息并验证 结束--------**/
	r.Middleware.Next()
}
