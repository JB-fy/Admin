package middleware

import (
	"api/internal/consts"
	daoAdmin "api/internal/dao/admin"
	daoOrg "api/internal/dao/org"
	"api/internal/utils"
	"api/internal/utils/jbctx"
	utilsToken "api/internal/utils/token"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// isForce 是否强制验证登录
func SceneLoginOfAdmin(isForce bool, tokenNameArr ...string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		/**--------验证token 开始--------**/
		token := ``
		for _, tokenName := range tokenNameArr {
			token = r.Header.Get(tokenName)
			if token != `` {
				break
			}
		}
		if token == `` {
			if isForce {
				r.SetError(utils.NewErrorCode(r.GetCtx(), 39994000, ``))
			} else {
				r.Middleware.Next()
			}
			return
		}

		tokenInfo, err := utilsToken.NewHandler(r.GetCtx()).Parse(token)
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
		info, _ := daoAdmin.Admin.CacheGetInfo(r.GetCtx(), gconv.Uint(tokenInfo.LoginId))
		if info.IsEmpty() {
			if isForce {
				r.SetError(utils.NewErrorCode(r.GetCtx(), 39994100, ``))
			} else {
				r.Middleware.Next()
			}
			return
		}
		if info[daoAdmin.Admin.Columns().IsStop].Uint8() == 1 {
			if isForce {
				r.SetError(utils.NewErrorCode(r.GetCtx(), 39994101, ``))
			} else {
				r.Middleware.Next()
			}
			return
		}
		switch info[daoAdmin.Admin.Columns().SceneId].String() {
		case consts.SCENE_ID_PLATFORM:
		case consts.SCENE_ID_ORG:
			relInfo, _ := daoOrg.Org.CacheGetInfo(r.GetCtx(), info[daoAdmin.Admin.Columns().RelId].Uint())
			if relInfo.IsEmpty() {
				if isForce {
					r.SetError(utils.NewErrorCode(r.GetCtx(), 39994200, ``))
				} else {
					r.Middleware.Next()
				}
				return
			}
			if relInfo[daoOrg.Org.Columns().IsStop].Uint8() == 1 {
				if isForce {
					r.SetError(utils.NewErrorCode(r.GetCtx(), 39994201, ``))
				} else {
					r.Middleware.Next()
				}
				return
			}
			info[`rel_info`] = gvar.New(relInfo.Map())
		default:
			err = utils.NewErrorCode(r.GetCtx(), 39999998, ``)
			return
		}

		info[consts.CTX_LOGIN_ID_NAME] = gvar.New(tokenInfo.LoginId) //所有场景追加这个字段，方便统一调用
		jbctx.SetLoginInfo(r, info)                                  //用户信息保存在协程上下文
		/**--------获取用户信息并验证 结束--------**/

		r.Middleware.Next()
	}
}
