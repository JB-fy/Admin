package middleware

import (
	dao "api/internal/dao/auth"
	"api/internal/utils"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Scene(r *ghttp.Request) {
	pathArr := strings.Split(r.URL.Path, `/`)
	sceneCode := pathArr[1]
	if sceneCode == `` {
		r.SetError(utils.NewErrorCode(r.GetCtx(), 39999999, ``))
		return
	}
	sceneInfo, _ := dao.Scene.ParseDbCtx(r.GetCtx()).Where(`sceneCode`, sceneCode).One()
	if sceneInfo.IsEmpty() {
		r.SetError(utils.NewErrorCode(r.GetCtx(), 39999999, ``))
		return
	}
	if sceneInfo[`isStop`].Int() > 0 {
		r.SetError(utils.NewErrorCode(r.GetCtx(), 39999998, ``))
		return
	}
	utils.SetCtxSceneInfo(r, sceneInfo)

	r.Middleware.Next()
}
