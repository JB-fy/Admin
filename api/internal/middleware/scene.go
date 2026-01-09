package middleware

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"
	"api/internal/utils/jbctx"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Scene(sceneIdOpt ...string) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		sceneId := ``
		if len(sceneIdOpt) > 0 {
			sceneId = sceneIdOpt[0]
		}
		if sceneId == `` {
			pathArr := strings.Split(r.URL.Path, `/`)
			sceneId = pathArr[1]
		}
		if sceneId == `` {
			r.SetError(utils.NewErrorCode(r.GetCtx(), 39999998, ``))
			return
		}
		sceneInfo, _ := daoAuth.Scene.CacheGetInfo(r.GetCtx(), sceneId)
		if sceneInfo.IsEmpty() {
			r.SetError(utils.NewErrorCode(r.GetCtx(), 39999998, ``))
			return
		}
		if sceneInfo[daoAuth.Scene.Columns().IsStop].Uint8() == 1 {
			r.SetError(utils.NewErrorCode(r.GetCtx(), 39999997, ``))
			return
		}
		jbctx.SetCtxSceneInfo(r, sceneInfo)

		r.Middleware.Next()
	}
}
