package middleware

import (
	daoAuth "api/internal/dao/auth"
	daoUsers "api/internal/dao/users"
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/gvar"
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

		sceneInfo := utils.GetCtxSceneInfo(r.GetCtx())
		jwt := utils.NewJWT(r.GetCtx(), sceneInfo[daoAuth.Scene.Columns().SceneConfig].Map())
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
		/* sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
		checkToken, _ := cache.NewToken(r.GetCtx(), sceneCode, claims.LoginId).Get()
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
		info, _ := daoUsers.Users.CtxDaoModel(r.GetCtx()).Filter(daoUsers.Users.Columns().UserId, claims.LoginId).One()
		if info.IsEmpty() {
			if isForce {
				r.SetError(utils.NewErrorCode(r.GetCtx(), 39994003, ``))
			} else {
				r.Middleware.Next()
			}
			return
		}
		if info[daoUsers.Users.Columns().IsStop].Uint() == 1 {
			if isForce {
				r.SetError(utils.NewErrorCode(r.GetCtx(), 39994004, ``))
			} else {
				r.Middleware.Next()
			}
			return
		}

		info[`login_id`] = gvar.New(claims.LoginId) //所有场景追加这个字段，方便统一调用
		utils.SetCtxLoginInfo(r, info)              //用户信息保存在协程上下文
		/**--------获取用户信息并验证 结束--------**/

		r.Middleware.Next()
	}
}
