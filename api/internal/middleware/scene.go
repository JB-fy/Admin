package middleware

import (
	dao "api/internal/dao/auth"
	"api/internal/packed"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Scene(r *ghttp.Request) {
	pathArr := strings.Split(r.URL.Path, "/")
	sceneCode := pathArr[1]
	if sceneCode == "" {
		packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 39999999, ""))
		return
	}
	sceneInfo, _ := dao.Scene.ParseDbCtx(r.GetCtx()).Where("sceneCode", sceneCode).One()
	if sceneInfo.IsEmpty() {
		packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 39999999, ""))
		return
	}
	if sceneInfo["isStop"].Int() > 0 {
		packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 39999998, ""))
		return
	}

	packed.SetCtxSceneInfo(r, sceneInfo)
	r.Middleware.Next()
}
