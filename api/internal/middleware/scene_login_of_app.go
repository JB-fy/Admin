package middleware

import (
	daoAuth "api/internal/dao/auth"
	daoUsers "api/internal/dao/users"
	"api/internal/utils"
	utilsToken "api/internal/utils/token"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
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
		tokenInfo, err := utilsToken.NewHandler(r.GetCtx(), sceneInfo[daoAuth.Scene.Columns().SceneConfig].Map()[`token_config`].(g.Map), sceneInfo[daoAuth.Scene.Columns().SceneCode].String()).Parse(token, r.GetClientIp())
		if err != nil {
			if isForce {
				r.SetError(err)
			} else {
				r.Middleware.Next()
			}
			return
		}
		/**--------验证token 结束--------**/

		/**--------获取登录用户信息并验证 开始--------**/
		info, _ := daoUsers.Users.CtxDaoModel(r.GetCtx()).Filter(daoUsers.Users.Columns().UserId, tokenInfo.LoginId).One()
		if info.IsEmpty() {
			if isForce {
				r.SetError(utils.NewErrorCode(r.GetCtx(), 39994100, ``))
			} else {
				r.Middleware.Next()
			}
			return
		}
		if info[daoUsers.Users.Columns().IsStop].Uint() == 1 {
			if isForce {
				r.SetError(utils.NewErrorCode(r.GetCtx(), 39994101, ``))
			} else {
				r.Middleware.Next()
			}
			return
		}

		info[`login_id`] = gvar.New(tokenInfo.LoginId) //所有场景追加这个字段，方便统一调用
		utils.SetCtxLoginInfo(r, info)                 //用户信息保存在协程上下文
		/**--------获取用户信息并验证 结束--------**/

		r.Middleware.Next()
	}
}
