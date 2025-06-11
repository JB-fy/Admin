package middleware

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Scene(r *ghttp.Request) {
	pathArr := strings.Split(r.URL.Path, `/`)
	sceneId := pathArr[1]
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
	utils.SetCtxSceneInfo(r, sceneInfo)

	r.Middleware.Next()
}
