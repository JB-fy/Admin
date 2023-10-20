package middleware

import (
	daoUser "api/internal/dao/user"
	"api/internal/utils"

	"github.com/gogf/gf/v2/net/ghttp"
)

// isForce 是否强制验证登录
func SceneLoginOfApp(isForce bool) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		/**--------验证token 开始--------**/
		token := r.Header.Get(`AppToken`)
		if token == `` {
			if isForce {
				r.SetError(utils.NewErrorCode(r.GetCtx(), 39994000, ``))
			} else {
				r.Middleware.Next()
			}
			return
		}

		jwt := utils.NewJWT(r.GetCtx(), utils.GetCtxSceneInfo(r.GetCtx())[`sceneConfig`].Map())
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if isForce {
				r.SetError(err)
			} else {
				r.Middleware.Next()
			}
			return
		}
		/**--------验证token 结束--------**/

		/**--------限制多地登录，多设备登录等情况下用（前置条件：登录时做过token缓存） 开始--------**/
		/* checkToken, _ := cache.NewToken(r.GetCtx(), claims.LoginId).Get()
		if checkToken != token {
			if isForce {
				r.SetError(utils.NewErrorCode(r.GetCtx(), 39994002, ``))
			} else {
				r.Middleware.Next()
			}
			return
		} */
		/**--------限制多地登录，多设备登录等情况下用（前置条件：登录时做过token缓存） 结束--------**/

		/**--------获取登录用户信息并验证 开始--------**/
		info, _ := daoUser.User.ParseDbCtx(r.GetCtx()).Where(`userId`, claims.LoginId).One()
		if info.IsEmpty() {
			if isForce {
				r.SetError(utils.NewErrorCode(r.GetCtx(), 39994003, ``))
			} else {
				r.Middleware.Next()
			}
			return
		}
		if info[`isStop`].Int() == 1 {
			if isForce {
				r.SetError(utils.NewErrorCode(r.GetCtx(), 39994004, ``))
			} else {
				r.Middleware.Next()
			}
			return
		}

		utils.SetCtxLoginInfo(r, info) //用户信息保存在协程上下文
		/**--------获取用户信息并验证 结束--------**/

		r.Middleware.Next()
	}
}
